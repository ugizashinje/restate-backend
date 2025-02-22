package services

import (
	"fmt"
	"time"
	"warrant-api/pkg/config"
	"warrant-api/pkg/ctx"
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/enum"
	"warrant-api/pkg/messages"
	"warrant-api/pkg/repo"
	"warrant-api/pkg/session"
	"warrant-api/pkg/superset"
	"warrant-api/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v4"
	"gorm.io/datatypes"
)

type AuthServiceImpl struct {
	Method                    jwt.SigningMethod
	UserService               UserServiceImpl
	CompanyService            CompanyServiceImpl
	LoginRepo                 repo.Repo[model.Login]
	LocationRepo              repo.Repo[model.Location]
	ConfirmationRepo          repo.Repo[model.Confirmation]
	AddressService            AddressServiceImpl
	TransactionalEmailService TransactionalEmailServiceImpl
	RestyClient               *resty.Client
	SupersetGuestTokenUrl     string
}

type RegisterCompanyManagerRequest struct {
	CompanyAddress   model.Address  `json:"companyAddress"`
	CompanyAddressID null.String    `json:"comapnyAddressID"`
	CompanyName      string         `json:"companyName"`
	CompanyShort     string         `json:"companyShort"`
	CompanyPIB       string         `json:"pib"`
	CompanyMn        string         `json:"companyMn"`
	CompanyPhone     string         `json:"companyPhone"`
	CompanyEmail     string         `json:"companyEmail" binding:"required,email"`
	CompanyMeta      datatypes.JSON `json:"meta"`
	UserAddress      model.Address  `json:"userAddress"`
	UserAddressID    null.String    `json:"userAddressId"`
	UserEmail        string         `json:"userEmail" binding:"required"`
	UserPassword     string         `json:"userPassword" binding:"required"`
	UserFirstName    string         `json:"userFirstName" binding:"required"`
	UserLastName     string         `json:"userLastName" binding:"required"`
	UserMn           null.String    `json:"userMn" binding:"required"`
	UserBirthplace   null.String    `json:"userBirthplace"`
	UserPhone        string         `json:"userPhone"`
}
type RegisterUserRequest struct {
	UserAddress    model.Address `json:"userAddress"`
	UserAddressID  null.String   `json:"userAddressId"`
	CompanyID      null.String   `json:"companyId"`
	UserEmail      string        `json:"userEmail" binding:"required"`
	UserPassword   string        `json:"userPassword" binding:"required"`
	UserFirstName  string        `json:"userFirstName" binding:"required"`
	UserLastName   string        `json:"userLastName" binding:"required"`
	UserMn         null.String   `json:"userMn" binding:"required"`
	UserBirthplace null.String   `json:"userBirthplace"`
	UserPhone      string        `json:"userPhone"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (svc *AuthServiceImpl) Login(g *gin.Context, loginRequest LoginRequest) (authToken *AuthResponse, err error) {
	login := &model.Login{}

	g.Set(ctx.Email, login.Email)
	users, err := svc.UserService.Search(g, map[string]interface{}{"users.email": loginRequest.Email, "users.is_active": true}, repo.Preload("Companies"))
	if err != nil {
		utils.Handle(errors.Errorf("Username password pair not found."))
	}

	if len(users) != 1 {
		login.Result = enum.UserNotFound
		svc.LoginRepo.Create(g, login)
		utils.Handle(errors.Errorf("Username password pair not found."))
	}
	user := users[0]
	login.Email = loginRequest.Email
	login.UserAgent = g.GetHeader("User-Agent")

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		login.Result = enum.WrongPassword
		svc.LoginRepo.Create(g, login)
		utils.Handle(errors.Errorf("Username password pair not found."))
	}

	//   key = /* Load key from somewhere, for example a file */
	accessClaims := jwt.NewWithClaims(svc.Method,
		jwt.MapClaims{
			"sub": user.BaseModel.ID,
			"exp": time.Now().Add(time.Minute * 30).Unix(),
		})
	accessToken, err := accessClaims.SignedString(config.JwtPrivateKey)
	if err != nil {
		login.Result = enum.InvalidToken
		svc.LoginRepo.Create(g, login)
		utils.Handle(err)
	}
	login.AccessToken = accessToken

	refreshClaims := jwt.NewWithClaims(svc.Method,
		jwt.MapClaims{
			"sub": user.BaseModel.ID,
			"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		})

	refreshToken, err := refreshClaims.SignedString(config.JwtPrivateKey)

	if err != nil {
		login.Result = enum.InvalidToken
		svc.LoginRepo.Create(g, login)
		utils.Handle(err)
	}
	login.RefreshToken = refreshToken
	login.Result = enum.Success

	svc.LoginRepo.Create(g, login)
	authToken = &AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	userSession := session.Session{
		User:         user,
		ThrottleHits: [6]int{},
	}
	if err := session.SaveSession(&userSession); err != nil {
		utils.Handle(err)
	}
	return authToken, nil
}

func (svc *AuthServiceImpl) Logout(g *gin.Context) (err error) {
	return nil
}

func (svc *AuthServiceImpl) Refresh(g *gin.Context, refresh RefreshRequest) (*AuthResponse, error) {
	claims := &jwt.MapClaims{}
	refreshToken, err := jwt.ParseWithClaims(refresh.RefreshToken, claims, func(t *jwt.Token) (interface{}, error) {
		public := config.JwtPrivateKey.Public()
		return public, nil
	})
	utils.Handle(err)

	userId, err := refreshToken.Claims.GetSubject()
	utils.Handle(err)

	_, err = svc.UserService.GetById(g, userId)
	utils.Handle(err)

	accessToken := jwt.NewWithClaims(svc.Method,
		jwt.MapClaims{
			"sub": userId,
			"exp": time.Now().Add(time.Minute * 30).Unix(),
		})
	signed, err := accessToken.SignedString(config.JwtPrivateKey)
	utils.Handle(err)
	authToken := &AuthResponse{}
	authToken.AccessToken = signed
	authToken.RefreshToken = refresh.RefreshToken
	return authToken, nil
}

// TODO Should be transactional
func (svc *AuthServiceImpl) RegisterCompanyManager(g *gin.Context, req RegisterCompanyManagerRequest) (err error) {
	companyAddress, _ := svc.AddressService.FindOrCreate(g, req.CompanyAddress, req.CompanyAddressID)
	userAddress, _ := svc.AddressService.FindOrCreate(g, req.CompanyAddress, req.CompanyAddressID)

	createCompany := CreateCompanyRequest{
		Name:      req.CompanyName,
		Short:     req.CompanyShort,
		PIB:       req.CompanyPIB,
		AddressID: companyAddress.ID,
		Mn:        req.CompanyMn,
		Phone:     req.CompanyPhone,
		Email:     req.CompanyEmail,
		Meta:      req.CompanyMeta,
	}
	company, err := svc.CompanyService.Create(g, createCompany)
	utils.Handle(err)

	createUser := CreateUserRequest{
		Email:      req.UserEmail,
		Password:   req.UserPassword,
		FirstName:  req.UserFirstName,
		LastName:   req.UserLastName,
		Mn:         req.UserMn,
		Birthplace: req.UserBirthplace,
		Phone:      req.UserPhone,
		AddressID:  userAddress.ID,
		CompanyID:  null.StringFrom(company.ID),
		Roles:      []string{"manager"},
	}
	_, err = svc.UserService.Create(g, createUser)
	utils.Handle(err)

	return nil
}
func (svc *AuthServiceImpl) RegisterUser(g *gin.Context, req RegisterUserRequest) (err error) {
	userAddress, _ := svc.AddressService.FindOrCreate(g, req.UserAddress, req.UserAddressID)

	createUser := CreateUserRequest{
		Email:      req.UserEmail,
		Password:   req.UserPassword,
		FirstName:  req.UserFirstName,
		LastName:   req.UserLastName,
		Mn:         req.UserMn,
		Birthplace: req.UserBirthplace,
		Phone:      req.UserPhone,
		AddressID:  userAddress.ID,
		CompanyID:  req.CompanyID,
		Roles:      []string{"manager"},
	}
	_, err = svc.UserService.Create(g, createUser)
	utils.Handle(err)

	return nil
}

func (svc *AuthServiceImpl) Confirm(g *gin.Context, code string) (*model.Confirmation, error) {
	confirmations := []model.Confirmation{}
	dbRes := svc.ConfirmationRepo.Search(g, &confirmations, map[string]any{"is_active": true, "code": code})
	utils.Handle(dbRes.Error)
	if dbRes.RowsAffected != 1 {
		utils.Handle(messages.Errorf(404, "Invalid code"))
	}
	confirmation := confirmations[0]
	confirmation.IsActive = false
	confirmation.Status = enum.Confirmed
	dbRes = svc.ConfirmationRepo.Save(g, &confirmation)
	utils.Handle(dbRes.Error)

	user, err := svc.UserService.GetById(g, confirmation.UserID)
	utils.Handle(err)
	user.IsActive = true
	svc.UserService.Save(g, user)
	return &confirmation, nil
}

func (svc *AuthServiceImpl) DashboardToken(g *gin.Context) (*superset.AuthSuccess, error) {

	guestRequest := session.SupersetGuestRequest{
		User: session.SupersetUser{
			Username: "embed_role",
		},
		RLS: []session.SupersetRLS{},
		Resources: []session.SupersetResources{{
			Type: "dashboard",
			Id:   "37b3c100-5a14-4070-8ce6-44d319b68d7f",
		}},
	}
	authSuccess := &superset.AuthSuccess{}
	authError := &superset.AuthError{}

	type TokenResponse struct {
		Result string `json:"result"`
	}
	tkn := TokenResponse{}
	type TokenError struct {
		Msg string `json:"msg"`
	}
	tknError := TokenError{}
	response, err := svc.RestyClient.R().
		SetResult(&tkn).
		SetHeader("Authorization", "Bearer "+superset.Token).
		SetError(&tknError).
		Get(config.Superset.Url + "/security/csrf_token/")

	fmt.Println("CSRF", response, "token", superset.Token)

	response, err = svc.RestyClient.R().
		SetBody(&guestRequest).
		SetHeader("Authorization", "Bearer "+superset.Token).
		SetHeader("X-CSRFToken", tkn.Result).
		SetResult(authSuccess).
		SetError(authError).
		Post(svc.SupersetGuestTokenUrl)
	//.Body(&guestRequest).Post(svc.SupersetGuestTokenUrl)
	cookie := response.Header().Get("Set-Cookie")
	authSuccess.AccessToken = cookie
	fmt.Println("response", response)
	if err != nil {
		utils.Handle(messages.Unauthorized())
	}
	return authSuccess, nil
}
