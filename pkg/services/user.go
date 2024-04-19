package services

import (
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/enum"
	"warrant-api/pkg/messages"
	"warrant-api/pkg/repo"
	"warrant-api/pkg/session"
	"warrant-api/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v4"
)

type UserServiceImpl struct {
	UserRepo                  repo.User
	UserCompanyRepo           repo.Repo[model.UserCompany]
	CompanyRepo               repo.Repo[model.Company]
	TransactionalEmailService TransactionalEmailServiceImpl
}

type CreateUserRequest struct {
	Email               string      `json:"email" binding:"required"`
	Password            string      `json:"password" binding:"required"`
	FirstName           string      `json:"firstName" binding:"required"`
	LastName            string      `json:"lastName" binding:"required"`
	Mn                  string      `json:"mn" binding:"required"`
	Birthplace          string      `json:"birthplace"`
	Department          null.String `json:"department"`
	Phone               string      `json:"phone"`
	AddressID           string      `json:"addressId" binding:"required"`
	CompanyID           string      `json:"companyId" binding:"required"`
	IsDriver            bool        `json:"isDriver"`
	LicenceNumber       null.String `json:"licenceNumber"`
	LicenceSerialNumber null.String `json:"licenceSerialNumber"`
	LicenceExpiry       null.String `json:"licenceExpiry" binding:"required,date"`
	LicenceIssued       null.String `json:"licenceIssued" binding:"required,date"`
	LicenceAuthority    null.String `json:"licenceAuthority"`
	Roles               []string    `json:"roles"`
}
type PatchUserRequest struct {
	Email               null.String    `json:"email"`
	Password            null.String    `json:"password"`
	FirstName           null.String    `json:"firstName"`
	LastName            null.String    `json:"LastName"`
	Mn                  null.String    `json:"mn"`
	Birthplace          null.String    `json:"birthplace"`
	Phone               null.String    `json:"phone"`
	AddressID           null.String    `json:"addressId"`
	Roles               pq.StringArray `json:"roles" gorm:"type:text[]"`
	Department          null.String    `json:"department"`
	AddToCompany        null.String    `json:"addToCompany"`
	RemoveFromCompany   null.String    `json:"removeFromCompany"`
	IsDriver            null.Bool      `json:"isDriver"`
	LicenceNumber       null.String    `json:"licenceNumber"`
	LicenceSerialNumber null.String    `json:"licenceSerialNumber"`
	LicenceExpiry       null.String    `json:"licenceExpiry" binding:"required,date"`
	LicenceIssued       null.String    `json:"licenceIssued" binding:"required,date"`
	LicenceAuthority    null.String    `json:"licenceAuthority"`
	Status              null.String    `json:"status"`
}

func (svc *UserServiceImpl) Create(g *gin.Context, createUser CreateUserRequest) (*model.User, error) {
	companies := model.Company{}

	compResult := svc.CompanyRepo.GetById(g, &companies, createUser.CompanyID)
	if compResult.Error != nil {
		return nil, compResult.Error
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(createUser.Password), bcrypt.DefaultCost)
	utils.Handle(err)

	user := &model.User{
		Email:               createUser.Email,
		Password:            string(hash),
		FirstName:           createUser.FirstName,
		LastName:            createUser.LastName,
		Mn:                  createUser.Mn,
		Birthplace:          createUser.Birthplace,
		Phone:               createUser.Phone,
		AddressID:           createUser.AddressID,
		IsDriver:            createUser.IsDriver,
		LicenceNumber:       createUser.LicenceNumber,
		Department:          createUser.Department,
		LicenceSerialNumber: createUser.LicenceSerialNumber,
		LicenceAuthority:    createUser.LicenceAuthority,
		LicenceExpiry:       utils.ParseNullDate(createUser.LicenceExpiry.String, createUser.LicenceExpiry.Valid),
		LicenceIssued:       utils.ParseNullDate(createUser.LicenceIssued.String, createUser.LicenceIssued.Valid),
		Status:              string(enum.UserUnverified),
	}

	dbRes := svc.UserRepo.Create(g, user)
	utils.Handle(dbRes.Error)
	user.IsActive = false
	dbRes = svc.UserRepo.Save(g, user)
	utils.Handle(dbRes.Error)

	err = svc.UserRepo.ManageCompanies(g, user, null.NewString(createUser.CompanyID, true), createUser.Roles, null.String{})
	utils.Handle(err)

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
	result := svc.UserRepo.GetById(c, user, userID, repo.Preload("Address"), repo.Preload("Companies"), repo.Preload("Categories"))
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
	if patch.AddressID.Valid {
		user.AddressID = patch.AddressID.String
	}
	if patch.Department.Valid {
		user.Department = patch.Department
	}
	if patch.Email.Valid {
		user.Email = patch.Email.String
	}
	if patch.FirstName.Valid {
		user.FirstName = patch.FirstName.String
	}
	if patch.LastName.Valid {
		user.LastName = patch.LastName.String
	}
	if patch.Mn.Valid {
		user.Mn = patch.Mn.String
	}
	if patch.IsDriver.Valid {
		user.IsDriver = patch.IsDriver.Bool
	}

	if patch.IsDriver.Valid {
		user.IsDriver = patch.IsDriver.Bool
	}

	if patch.LicenceNumber.Valid {
		user.LicenceNumber = patch.LicenceNumber
	}

	if patch.LicenceSerialNumber.Valid {
		user.LicenceSerialNumber = patch.LicenceSerialNumber
	}

	if patch.LicenceAuthority.Valid {
		user.LicenceAuthority = patch.LicenceAuthority
	}

	if patch.LicenceExpiry.Valid {
		user.LicenceExpiry = utils.ParseNullDate(patch.LicenceExpiry.String, patch.LicenceExpiry.Valid)
	}

	if patch.LicenceIssued.Valid {
		user.LicenceIssued = utils.ParseNullDate(patch.LicenceIssued.String, patch.LicenceIssued.Valid)
	}
	if patch.Status.Valid {
		user.Status = patch.Status.String
	}

	if patch.AddToCompany.Valid || patch.RemoveFromCompany.Valid {
		err := svc.UserRepo.ManageCompanies(c, user, patch.AddToCompany, patch.Roles, patch.RemoveFromCompany)
		utils.Handle(err)
	}

	if patch.Password.Valid {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(patch.Password.String))
		if err != nil {
			panic(errors.Errorf("Username and password pair not found."))
		}
	}
	dbRes := svc.UserRepo.Save(c, user)
	utils.Handle(dbRes.Error)

	dbRes = svc.UserRepo.GetById(c, user, user.ID, repo.Preload("Address"), repo.Preload("Categories"), repo.Preload("Companies"))
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
	opts = append(opts, repo.GrantUser, repo.Preload("Companies"))
	result := svc.UserRepo.Search(c, &users, query, opts...)
	utils.Handle(result.Error)
	return users, nil
}
