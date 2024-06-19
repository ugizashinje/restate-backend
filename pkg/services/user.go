package services

import (
	"net/http"
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/enum"
	"warrant-api/pkg/messages"
	"warrant-api/pkg/repo"
	"warrant-api/pkg/session"
	"warrant-api/pkg/utils"
	"warrant-api/pkg/utils/transformer"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepo                  repo.User
	CompanyRepo               repo.Repo[model.Company]
	TransactionalEmailService TransactionalEmailServiceImpl
}

type CreateUserRequest struct {
	Email      string      `json:"email" binding:"required"`
	Avatar     null.String `json:"avatar"`
	Password   string      `json:"password" binding:"required"`
	FirstName  string      `json:"firstName" binding:"required"`
	LastName   string      `json:"lastName" binding:"required"`
	Mn         null.String `json:"mn"`
	Birthplace null.String `json:"birthplace"`
	Phone      string      `json:"phone"`
	AddressID  string      `json:"addressId" binding:"required"`
	CompanyID  null.String `json:"companyId"`
	Roles      []string    `json:"roles"`
}
type PatchUserRequest struct {
	Email             null.String    `json:"email"`
	Avatar            null.String    `json:"avatar"`
	Password          null.String    `json:"password"`
	FirstName         null.String    `json:"firstName"`
	LastName          null.String    `json:"LastName"`
	Mn                null.String    `json:"mn"`
	Birthplace        null.String    `json:"birthplace"`
	Phone             null.String    `json:"phone"`
	AddressID         null.String    `json:"addressId"`
	Roles             pq.StringArray `json:"roles" gorm:"type:text[]"`
	AddToCompany      null.String    `json:"addToCompany"`
	RemoveFromCompany null.String    `json:"removeFromCompany"`
	LicenceNumber     null.String    `json:"licenceNumber"`
	Status            null.String    `json:"status"`
}

func (svc *UserServiceImpl) Create(g *gin.Context, createUser CreateUserRequest) (*model.User, error) {
	company := model.Company{}
	var dbRes *gorm.DB
	if createUser.CompanyID.Valid {
		dbRes = svc.CompanyRepo.GetById(g, &company, createUser.CompanyID.String)
		if dbRes.RowsAffected != 1 {
			utils.Handle(messages.Errorf(http.StatusBadRequest, "Invalid company"))
		}
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(createUser.Password), bcrypt.DefaultCost)
	utils.Handle(err)

	user := &model.User{
		Email:      createUser.Email,
		Avatar:     createUser.Avatar,
		Password:   string(hash),
		FirstName:  createUser.FirstName,
		LastName:   createUser.LastName,
		Phone:      createUser.Phone,
		AddressID:  createUser.AddressID,
		Status:     string(enum.UserUnverified),
		Mn:         createUser.Mn,
		Birthplace: createUser.Birthplace,
	}

	dbRes = svc.UserRepo.Create(g, user)
	utils.Handle(dbRes.Error)
	user.IsActive = false
	dbRes = svc.UserRepo.Save(g, user)
	utils.Handle(dbRes.Error)

	// err = svc.UserRepo.ManageCompanies(g, user, createUser.CompanyID, createUser.Roles, null.String{})
	// utils.Handle(err)

	svc.TransactionalEmailService.VerifyEmail(g, user)

	return user, nil
}

func (svc *UserServiceImpl) Delete(g *gin.Context, userID string) (user *model.User, err error) {
	user = &model.User{}
	result := svc.UserRepo.Delete(g, user, userID)
	utils.Handle(result.Error)

	return user, nil
}

func (svc *UserServiceImpl) GetById(c *gin.Context, userID string) (user *model.User, err error) {
	user = &model.User{}
	result := svc.UserRepo.GetById(c, user, userID, repo.Preload("Address"), repo.Preload("Companies"))
	utils.Handle(result.Error)
	if result.RowsAffected != 1 {
		utils.Handle(messages.Errorf(404, "User not found"))
	}
	return user, nil
}

func (svc *UserServiceImpl) Update(c *gin.Context, userID string, patch PatchUserRequest) (user *model.User, err error) {
	user = &model.User{}
	result := svc.UserRepo.GetById(c, user, userID)
	utils.Handle(result.Error)

	transformer.Patch(user, patch)

	// if patch.AddToCompany.Valid || patch.RemoveFromCompany.Valid {
	// 	err := svc.UserRepo.ManageCompanies(c, user, patch.AddToCompany, patch.Roles, patch.RemoveFromCompany)
	// 	utils.Handle(err)
	// }

	if patch.Password.Valid {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(patch.Password.String))
		if err != nil {
			panic(errors.Errorf("Username and password pair not found."))
		}
	}
	dbRes := svc.UserRepo.Save(c, user)
	utils.Handle(dbRes.Error)

	dbRes = svc.UserRepo.GetById(c, user, user.ID, repo.Preload("Address"), repo.Preload("Companies"), repo.Preload("Collections"))
	userSession := session.GetSession(c)
	userSession.User = *user
	session.SaveSession(userSession)
	utils.Handle(dbRes.Error)
	return user, nil
}

func (svc *UserServiceImpl) Save(c *gin.Context, user *model.User) (res *model.User, err error) {
	dbRes := svc.UserRepo.Save(c, user)
	utils.Handle(dbRes.Error)
	return res, nil
}

func (svc *UserServiceImpl) Search(c *gin.Context, query map[string]any, opts ...repo.RepoOption) (users []model.User, err error) {
	users = []model.User{}
	// opts = append(opts, repo.GrantUser, repo.Preload("Companies"))
	result := svc.UserRepo.Search(c, &users, query, opts...)
	utils.Handle(result.Error)
	return users, nil
}
