package server

import (
	"net/http"
	"reflect"
	"runtime/debug"
	"strings"
	"time"
	"warrant-api/pkg/config"
	"warrant-api/pkg/ctx"
	"warrant-api/pkg/session"

	// "warrant-api/pkg/session"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func dbMiddleware(gormDB *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(ctx.Transaction, gormDB)
		c.Next()
	}
}

func recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logLine := string(debug.Stack())
				log.Errorf("stacktrace from panic: \n" + logLine)
				t := reflect.TypeOf(err)
				i := reflect.TypeOf((*error)(nil)).Elem()
				if t.Implements(i) {
					e := err.(error)
					c.JSON(500, gin.H{
						"error": e.Error()})
				} else {
					c.JSON(500, gin.H{
						"error": err})
				}
			}
		}()

		c.Next()
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	}
}

func authMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.URL.Path, "/v1/auth/superset") && (strings.HasPrefix(c.Request.URL.Path, "/v1/auth") ||
			strings.HasPrefix(c.Request.URL.Path, "/swagger")) {
			c.Next()
			return
		}

		claims := &jwt.MapClaims{}

		authorization := c.Request.Header.Get("Authorization")

		if strings.HasPrefix(authorization, "Bearer") {
			splited := strings.Split(authorization, " ")
			strToken := splited[1]
			jwtToken, err := jwt.ParseWithClaims(strToken, claims, func(t *jwt.Token) (interface{}, error) {
				return config.JwtPrivateKey.Public(), nil
			})
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"errors": []string{"Access denied."}})
				c.Abort()
				return
			}
			if expiry, err := jwtToken.Claims.GetExpirationTime(); expiry.Time.Before(time.Now()) || err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"errors": []string{"Token expired."}})
				c.Abort()
				return
			}
			if err = session.UpdateThrottle(c, jwtToken); err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"errors": []string{"Too many requests."}})
				c.Abort()
				return
			}
		}
		if subject, err := claims.GetSubject(); subject == "" || err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"errors": []string{"Access denied"}})
			c.Abort()
		} else {
			c.Next()
		}

	}
}
