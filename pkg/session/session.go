package session

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
	"warrant-api/pkg/config"
	"warrant-api/pkg/ctx"
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/enum"
	"warrant-api/pkg/messages"
	"warrant-api/pkg/utils"
	"warrant-api/pkg/utils/transformer"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

// {
// 	"resources": [
// 	  {
// 		"id": "string",
// 		"type": "dashboard"
// 	  }
// 	],
// 	"rls": [
// 	  {
// 		"clause": "string",
// 		"dataset": 0
// 	  }
// 	],
// 	"user": {
// 	  "first_name": "string",
// 	  "last_name": "string",
// 	  "username": "string"
// 	}
//   }

type SupersetResources struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type SupersetRLS struct {
	Clause  string `json:"clause"`
	Dataset string `json:"dataset"`
}

type SupersetUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
}
type SupersetGuestRequest struct {
	Resources []SupersetResources `json:"resources"`
	RLS       []SupersetRLS       `json:"rls"`
	User      SupersetUser        `json:"user"`
}

type Session struct {
	User         model.User `json:"user"`
	ThrottleHits [6]int     `json:"Hits"`
}

func (s *Session) Roles(companyId string) (isManager, isDispatcher, isDriver bool) {
	for _, userCompany := range s.User.Companies {
		if companyId == userCompany.CompanyID {
			for _, role := range userCompany.Roles {
				if role == string(enum.Manager) {
					isManager = true
				}
				if role == string(enum.Dispatcher) {
					isDispatcher = true
				}
				if role == string(enum.Driver) {
					isDriver = true
				}

			}

		}
	}
	return isManager, isDispatcher, isDriver
}
func (s *Session) LogEvent(g *gin.Context, warrant *model.Warrant, object any, eventType enum.EventType) error {
	if warrant == nil {
		return messages.Errorf(http.StatusBadRequest, "Warrant not provided")
	}
	found := false
	for _, c := range s.User.Companies {
		if c.CompanyID == warrant.CompanyID {
			found = true
		}
	}
	if !found {
		return messages.Unauthorized()
	}
	event := model.Change{
		Ad:     warrant.ID,
		UserID: s.User.ID,
		Event:  eventType,
	}
	pruned := transformer.Map(object)
	buff, err := json.Marshal(pruned)
	if err != nil {
		return err
	}
	event.Meta = buff

	dbRaw, ok := g.Get(ctx.Transaction)
	if !ok {
		return messages.Errorf(http.StatusInternalServerError, "Ivalid session")
	}
	db, ok := dbRaw.(*gorm.DB)
	if !ok {
		return messages.Errorf(http.StatusInternalServerError, "Ivalid session")
	}
	db = db.Create(&event)
	utils.Handle(db.Error)
	return nil
}

var redisPool *redis.Pool

func Init() {
	redisPool = &redis.Pool{
		MaxIdle:   80,
		MaxActive: 1200,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", config.Redis.Host+":"+config.Redis.Port, redis.DialPassword(config.Redis.Password))
			if err != nil {
				log.Printf("Error:  failed to init redis", err.Error())
				os.Exit(2)
			}
			return conn, nil
		},
	}
}

func UpdateThrottle(g *gin.Context, jwtToken *jwt.Token) error {
	conn := redisPool.Get()
	defer conn.Close()
	userId, err := jwtToken.Claims.GetSubject()
	if err != nil {
		return nil
	}
	session := Session{}
	data, err := conn.Do("GET", userId)
	if err != nil {
		return messages.Errorf(http.StatusTooManyRequests, "Rate limit")
	}

	err = json.Unmarshal(data.([]byte), &session)
	if err != nil {
		return messages.Errorf(http.StatusTooManyRequests, "Rate limit")
	}
	slot := (time.Now().Unix() / (5 * 60)) % 6
	session.ThrottleHits[slot] += 1

	for _, v := range session.ThrottleHits {
		if v > config.Auth.RateLimitPerSlot {
			return messages.Errorf(http.StatusTooManyRequests, "Rate limit")
		}
	}
	data, err = json.Marshal(session)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", userId, data)
	if err != nil {
		return messages.Errorf(http.StatusTooManyRequests, "Rate limit")
	}
	g.Set(ctx.Session, &session)
	return nil
}

func SaveSession(userSession *Session) error {
	conn := redisPool.Get()
	defer conn.Close()
	buff, err := json.Marshal(userSession)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", userSession.User.ID, buff)
	return err

}

func GetSession(g *gin.Context) *Session {
	val, ok := g.Get(ctx.Session)
	if !ok {
		return nil
	}
	userSession, ok := val.(*Session)
	if !ok {
		return nil
	}
	return userSession

}
