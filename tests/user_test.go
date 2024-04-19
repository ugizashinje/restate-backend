package tests

import (
	"log"
	"net/http"
	"net/http/httptest"
	url "net/url"
	"os"
	"testing"
	"warrant-api/pkg/config"
	"warrant-api/pkg/ctx"
	"warrant-api/pkg/db"
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/enum"
	"warrant-api/pkg/repo"
	"warrant-api/pkg/services"
	"warrant-api/pkg/session"
	"warrant-api/pkg/wire"
	"warrant-api/tests/fixtures"

	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
)

func SetupTest() *gin.Context {
	wire.Svc = wire.Init("testing")
	config.Db.DbLogging = true
	gormDB, _ := db.Init(config.Db)
	w := httptest.NewRecorder()
	g, _ := gin.CreateTestContext(w)
	g.Request = &http.Request{}
	var err error
	g.Request.URL, err = url.Parse("?page=1&pageSize=10")
	if err != nil {
		os.Exit(-1)
	}

	g.Set(ctx.Transaction, gormDB)

	return g
}

type TestSessions struct {
	UserManagerSession                        session.Session
	UserManagerWithDifferentCompanySession    session.Session
	UserDispatcherFromDifferentCompanySession session.Session
	UserDriverFromDifferentCompanySession     session.Session
	UserDispatcherInTheSameCompanySession     session.Session
	UserDriverInTheSameCompanySession         session.Session
}

func UserSessions() *TestSessions {
	g := SetupTest()
	registerRequest := fixtures.RegisterRegularCompany()
	err5 := wire.Svc.AuthService.Register(g, registerRequest)
	if err5 != nil {
		os.Exit(-1)
	}

	registerManagerWithDifferentCompanyRequest := fixtures.RegisterRegularCompany()
	err0 := wire.Svc.AuthService.Register(g, registerManagerWithDifferentCompanyRequest)
	if err0 != nil {
		os.Exit(-1)
	}

	registerDispatcherRequest := fixtures.RegisterRegularCompany()
	registerDispatcherRequest.UserRoles = []string{"dispatcher"}
	registerDispatcherRequest.UserIsDriver = true
	err1 := wire.Svc.AuthService.Register(g, registerDispatcherRequest)
	if err1 != nil {
		os.Exit(-1)
	}

	registerDriverRequest := fixtures.RegisterRegularCompany()
	registerDriverRequest.UserRoles = []string{"driver"}
	registerDriverRequest.UserIsDriver = true
	err2 := wire.Svc.AuthService.Register(g, registerDriverRequest)
	if err2 != nil {
		os.Exit(-1)
	}

	userRepo := repo.User{}
	confirmationRepo := repo.Repo[model.Confirmation]{}
	confirmations := []model.Confirmation{}
	confirmationsManagerWithDifferentCompany := []model.Confirmation{}
	confirmationsDispatcher := []model.Confirmation{}
	confirmationsDriver := []model.Confirmation{}
	users := []model.User{}
	usersManagerWithDifferentCompany := []model.User{}
	userDispatcher := []model.User{}
	userDriver := []model.User{}

	g.Request.URL, _ = url.Parse("v1/auth/confirm/")

	result := userRepo.Search(g, &users, map[string]any{"email": registerRequest.UserEmail}, repo.Preload("Companies"))
	if result.Error != nil {
		os.Exit(-1)
	}

	result0 := userRepo.Search(g, &usersManagerWithDifferentCompany, map[string]any{"email": registerManagerWithDifferentCompanyRequest.UserEmail}, repo.Preload("Companies"))
	if result0.Error != nil {
		os.Exit(-1)
	}

	result1 := userRepo.Search(g, &userDispatcher, map[string]any{"email": registerDispatcherRequest.UserEmail}, repo.Preload("Companies"))
	if result1.Error != nil {
		os.Exit(-1)
	}

	result2 := userRepo.Search(g, &userDriver, map[string]any{"email": registerDriverRequest.UserEmail}, repo.Preload("Companies"))
	if result2.Error != nil {
		os.Exit(-1)
	}

	result3 := confirmationRepo.Search(g, &confirmations, map[string]any{"user_id": users[0].ID})
	if result3.Error != nil {
		os.Exit(-1)
	}

	result4 := confirmationRepo.Search(g, &confirmationsManagerWithDifferentCompany, map[string]any{"user_id": usersManagerWithDifferentCompany[0].ID})
	if result4.Error != nil {
		os.Exit(-1)
	}

	result5 := confirmationRepo.Search(g, &confirmationsDispatcher, map[string]any{"user_id": userDispatcher[0].ID})
	if result5.Error != nil {
		os.Exit(-1)
	}

	result6 := confirmationRepo.Search(g, &confirmationsDriver, map[string]any{"user_id": userDriver[0].ID})
	if result6.Error != nil {
		os.Exit(-1)
	}

	userManagerSession := session.Session{
		User:         users[0],
		ThrottleHits: [6]int{},
	}

	userManagerWithDifferentCompanySession := session.Session{
		User:         usersManagerWithDifferentCompany[0],
		ThrottleHits: [6]int{},
	}

	userDispatcherFromDifferentCompanySession := session.Session{
		User:         userDispatcher[0],
		ThrottleHits: [6]int{},
	}

	userDriverFromDifferentCompanySession := session.Session{
		User:         userDriver[0],
		ThrottleHits: [6]int{},
	}

	g.Set("session", &userManagerSession)

	loginRequest := services.LoginRequest{
		Email:    registerRequest.UserEmail,
		Password: registerRequest.UserPassword,
	}

	confirmation, err := wire.Svc.AuthService.Confirm(g, confirmations[0].Code)

	if err != nil || confirmation.Status != enum.Confirmed {
		log.Fatal("Failed to confirm userManagerSession or confirmation status is not Confirmed")
	}

	confirmationDispatcher, err := wire.Svc.AuthService.Confirm(g, confirmationsDispatcher[0].Code)
	if err != nil || confirmationDispatcher.Status != enum.Confirmed {
		log.Fatal("Failed to confirm confirmationDispatcher or confirmation status is not Confirmed")
	}

	confirmationDriver, err := wire.Svc.AuthService.Confirm(g, confirmationsDriver[0].Code)
	if err != nil || confirmationDriver.Status != enum.Confirmed {
		log.Fatal("Failed to confirm confirmationsDriver or confirmation status is not Confirmed")
	}

	authResponse, err := wire.Svc.AuthService.Login(g, loginRequest)

	if err != nil || authResponse.AccessToken == "" {
		os.Exit(-1)
	}

	createDispatherUserInTheSameCompanyRequest := fixtures.CreateUserInSameCompany()

	address := []model.Address{}
	addressRepo := repo.Repo[model.Address]{}

	result7 := addressRepo.Search(g, &address, map[string]any{"street": registerRequest.UserAddressStreet})
	if result7.Error != nil {
		os.Exit(-1)
	}
	company := []model.Company{}
	companyRepo := repo.Repo[model.Company]{}

	result8 := companyRepo.Search(g, &company, map[string]any{"pib": registerRequest.CompanyPIB})
	if result8.Error != nil {
		os.Exit(-1)
	}

	createDispatherUserInTheSameCompanyRequest.AddressID = address[0].ID
	createDispatherUserInTheSameCompanyRequest.CompanyID = company[0].ID

	createDispatherUserInTheSameCompanyResponse, err := wire.Svc.UserService.Create(g, createDispatherUserInTheSameCompanyRequest)

	if err != nil || createDispatherUserInTheSameCompanyResponse.Email != createDispatherUserInTheSameCompanyRequest.Email {
		log.Fatal("Failed to createDispatherUserInTheSameCompanyResponse")
	}

	dispatherUserInTheSameCompany := []model.User{}
	dispatherUserInTheSameCompanyConfurmation := []model.Confirmation{}

	result10 := userRepo.Search(g, &dispatherUserInTheSameCompany, map[string]any{"email": createDispatherUserInTheSameCompanyRequest.Email}, repo.Preload("Companies"))
	if result10.Error != nil {
		os.Exit(-1)
	}

	result11 := confirmationRepo.Search(g, &dispatherUserInTheSameCompanyConfurmation, map[string]any{"user_id": dispatherUserInTheSameCompany[0].ID})
	if result11.Error != nil {
		os.Exit(-1)
	}

	userDispatcherInTheSameCompanySession := session.Session{
		User:         dispatherUserInTheSameCompany[0],
		ThrottleHits: [6]int{},
	}

	confirmationDispatherUserInTheSameCompany, err := wire.Svc.AuthService.Confirm(g, dispatherUserInTheSameCompanyConfurmation[0].Code)
	if err != nil || confirmationDispatherUserInTheSameCompany.Status != enum.Confirmed {
		log.Fatal("Failed to confirm confirmationDispatherUserInTheSameCompany or confirmation status is not Confirmed")
	}

	createDriverUserInTheSameCompanyRequest := fixtures.CreateUserInSameCompany()
	createDriverUserInTheSameCompanyRequest.Roles = []string{"driver"}
	createDriverUserInTheSameCompanyRequest.IsDriver = true

	address1 := []model.Address{}
	addressRepo1 := repo.Repo[model.Address]{}

	result20 := addressRepo1.Search(g, &address1, map[string]any{"street": registerRequest.UserAddressStreet})
	if result20.Error != nil {
		os.Exit(-1)
	}

	company1 := []model.Company{}
	companyRepo1 := repo.Repo[model.Company]{}

	result15 := companyRepo1.Search(g, &company1, map[string]any{"pib": registerRequest.CompanyPIB})
	if result15.Error != nil {
		os.Exit(-1)
	}

	createDriverUserInTheSameCompanyRequest.AddressID = address1[0].ID
	createDriverUserInTheSameCompanyRequest.CompanyID = company1[0].ID

	createDriverUserInTheSameCompanyResponse, err := wire.Svc.UserService.Create(g, createDriverUserInTheSameCompanyRequest)

	if err != nil || createDriverUserInTheSameCompanyResponse.Email != createDriverUserInTheSameCompanyRequest.Email {
		log.Fatal("Failed to createDriverUserInTheSameCompanyResponse")
	}

	driverUserInTheSameCompany := []model.User{}
	driverUserInTheSameCompanyConfirmation := []model.Confirmation{}

	result44 := userRepo.Search(g, &driverUserInTheSameCompany, map[string]any{"email": createDriverUserInTheSameCompanyRequest.Email}, repo.Preload("Companies"))
	if result44.Error != nil {
		os.Exit(-1)
	}

	result21 := confirmationRepo.Search(g, &driverUserInTheSameCompanyConfirmation, map[string]any{"user_id": driverUserInTheSameCompany[0].ID})
	if result21.Error != nil {
		os.Exit(-1)
	}

	userDriverInTheSameCompanySession := session.Session{
		User:         driverUserInTheSameCompany[0],
		ThrottleHits: [6]int{},
	}

	confirmationDispatherUserInSameCompany, err := wire.Svc.AuthService.Confirm(g, driverUserInTheSameCompanyConfirmation[0].Code)
	if err != nil || confirmationDispatherUserInSameCompany.Status != enum.Confirmed {
		log.Fatal("Failed to confirm confirmationDispatherUserInSameCompany or confirmation status is not Confirmed")
	}

	return &TestSessions{
		UserManagerSession:                        userManagerSession,
		UserManagerWithDifferentCompanySession:    userManagerWithDifferentCompanySession,
		UserDispatcherFromDifferentCompanySession: userDispatcherFromDifferentCompanySession,
		UserDriverFromDifferentCompanySession:     userDriverFromDifferentCompanySession,
		UserDispatcherInTheSameCompanySession:     userDispatcherInTheSameCompanySession,
		UserDriverInTheSameCompanySession:         userDriverInTheSameCompanySession,
	}
}

func TestSpec(t *testing.T) {
	g := SetupTest()
	sessions := UserSessions()
	g.Set("session", &sessions.UserManagerSession)
	g.Request.URL, _ = url.Parse("v1/auth/confirm/")

	Convey("Register company and user", t, func() {
		registerRequest := fixtures.RegisterRegularCompany()
		err := wire.Svc.AuthService.Register(g, registerRequest)
		So(err, ShouldEqual, nil)

		Convey("Confirm email of registred", func() {
			userRepo := repo.User{}
			confirmationRepo := repo.Repo[model.Confirmation]{}
			confirmations := []model.Confirmation{}
			users := []model.User{}

			g.Request.URL, _ = url.Parse("v1/auth/confirm/")

			result := userRepo.Search(g, &users, map[string]any{"email": registerRequest.UserEmail}, repo.Preload("Companies"))
			ShouldBeNil(result.Error)
			So(len(users), ShouldEqual, 1)

			result = confirmationRepo.Search(g, &confirmations, map[string]any{"user_id": users[0].ID})
			So(result.Error, ShouldBeNil)
			So(len(confirmations), ShouldEqual, 1)

			userManagerSession := session.Session{
				User:         users[0],
				ThrottleHits: [6]int{},
			}

			g.Set("session", &userManagerSession)

			loginRequest := services.LoginRequest{
				Email:    registerRequest.UserEmail,
				Password: registerRequest.UserPassword,
			}

			So(func() { wire.Svc.AuthService.Login(g, loginRequest) }, ShouldPanic)

			confirmation, err := wire.Svc.AuthService.Confirm(g, confirmations[0].Code)
			So(confirmation.Status, ShouldEqual, enum.Confirmed)

			Convey("New user should be active", func() {
				authResponse, _ := wire.Svc.AuthService.Login(g, loginRequest)

				So(authResponse, ShouldNotBeNil)
				So(authResponse.AccessToken, ShouldNotBeZeroValue)
				So(1, ShouldNotEqual, err)

				Convey("New Manager user can create user Dispatcher", func() {
					createDispatherUserInTheSameCompanyRequest := fixtures.CreateUserInSameCompany()

					address := []model.Address{}
					addressRepo := repo.Repo[model.Address]{}

					result := addressRepo.Search(g, &address, map[string]any{"street": registerRequest.UserAddressStreet})
					ShouldBeNil(result.Error)

					company := []model.Company{}
					companyRepo := repo.Repo[model.Company]{}

					result1 := companyRepo.Search(g, &company, map[string]any{"pib": registerRequest.CompanyPIB})
					ShouldBeNil(result1.Error)

					createDispatherUserInTheSameCompanyRequest.AddressID = address[0].ID
					createDispatherUserInTheSameCompanyRequest.CompanyID = company[0].ID

					createDispatherUserInTheSameCompanyResponse, _ := wire.Svc.UserService.Create(g, createDispatherUserInTheSameCompanyRequest)

					So(createDispatherUserInTheSameCompanyResponse, ShouldNotBeNil)
					So(err, ShouldBeNil)
					So(createDispatherUserInTheSameCompanyResponse.Email, ShouldEqual, createDispatherUserInTheSameCompanyRequest.Email)

					dispatherUserInTheSameCompany := []model.User{}
					dispatherUserInTheSameCompanyConfurmation := []model.Confirmation{}

					result10 := userRepo.Search(g, &dispatherUserInTheSameCompany, map[string]any{"email": createDispatherUserInTheSameCompanyRequest.Email}, repo.Preload("Companies"))
					ShouldBeNil(result10.Error)
					So(len(dispatherUserInTheSameCompany), ShouldEqual, 1)

					result11 := confirmationRepo.Search(g, &dispatherUserInTheSameCompanyConfurmation, map[string]any{"user_id": dispatherUserInTheSameCompany[0].ID})
					So(result11.Error, ShouldBeNil)
					So(len(dispatherUserInTheSameCompanyConfurmation), ShouldEqual, 1)

					// userDispatcherInTheSameCompanySession := session.Session{
					// 	User:         dispatherUserInTheSameCompany[0],
					// 	ThrottleHits: [6]int{},
					// }

					confirmationDispatherUserInTheSameCompany, err := wire.Svc.AuthService.Confirm(g, dispatherUserInTheSameCompanyConfurmation[0].Code)
					So(confirmationDispatherUserInTheSameCompany.Status, ShouldEqual, enum.Confirmed)
					So(err, ShouldBeNil)

					Convey("New Manager user can create user Driver", func() {
						createDriverUserInTheSameCompanyRequest := fixtures.CreateUserInSameCompany()
						createDriverUserInTheSameCompanyRequest.Roles = []string{"driver"}
						createDriverUserInTheSameCompanyRequest.IsDriver = true

						address := []model.Address{}
						addressRepo := repo.Repo[model.Address]{}

						result := addressRepo.Search(g, &address, map[string]any{"street": registerRequest.UserAddressStreet})
						ShouldBeNil(result.Error)

						company := []model.Company{}
						companyRepo := repo.Repo[model.Company]{}

						result1 := companyRepo.Search(g, &company, map[string]any{"pib": registerRequest.CompanyPIB})
						ShouldBeNil(result1.Error)

						createDriverUserInTheSameCompanyRequest.AddressID = address[0].ID
						createDriverUserInTheSameCompanyRequest.CompanyID = company[0].ID

						createDriverUserInTheSameCompanyResponse, _ := wire.Svc.UserService.Create(g, createDriverUserInTheSameCompanyRequest)

						So(createDriverUserInTheSameCompanyResponse, ShouldNotBeNil)
						So(err, ShouldBeNil)
						So(createDriverUserInTheSameCompanyResponse.Email, ShouldEqual, createDriverUserInTheSameCompanyRequest.Email)

						driverUserInTheSameCompany := []model.User{}
						driverUserInTheSameCompanyConfirmation := []model.Confirmation{}

						result10 := userRepo.Search(g, &driverUserInTheSameCompany, map[string]any{"email": createDriverUserInTheSameCompanyRequest.Email}, repo.Preload("Companies"))
						ShouldBeNil(result10.Error)
						So(len(driverUserInTheSameCompany), ShouldEqual, 1)

						result11 := confirmationRepo.Search(g, &driverUserInTheSameCompanyConfirmation, map[string]any{"user_id": driverUserInTheSameCompany[0].ID})
						So(result11.Error, ShouldBeNil)
						So(len(driverUserInTheSameCompanyConfirmation), ShouldEqual, 1)

						userDriverInTheSameCompanySession := session.Session{
							User:         driverUserInTheSameCompany[0],
							ThrottleHits: [6]int{},
						}

						confirmationDispatherUserInSameCompany, err := wire.Svc.AuthService.Confirm(g, driverUserInTheSameCompanyConfirmation[0].Code)
						So(confirmationDispatherUserInSameCompany.Status, ShouldEqual, enum.Confirmed)
						So(err, ShouldBeNil)

						// GET BY ID
						Convey("New Manager user can get Users by id", func() {
							getUserByIdResponse, err := wire.Svc.UserService.GetById(g, createDispatherUserInTheSameCompanyResponse.ID)

							So(getUserByIdResponse, ShouldNotBeNil)
							So(err, ShouldBeNil)
							So(getUserByIdResponse.ID, ShouldEqual, createDispatherUserInTheSameCompanyResponse.ID)
						})

						// UPDATE ???
						Convey("User Manager can patch user Dispatcher", func() {
							patchUserRequest := fixtures.UpdateUser()

							patchUserResponse, err := wire.Svc.UserService.Update(g, createDispatherUserInTheSameCompanyResponse.ID, patchUserRequest)
							So(patchUserResponse, ShouldNotBeNil)
							So(patchUserResponse.FirstName, ShouldEqual, patchUserRequest.FirstName.String)
							So(err, ShouldBeNil)
						})

						Convey("User Manager can patch user Driver", func() {
							patchUserRequest := fixtures.UpdateUser()

							patchUserResponse, err := wire.Svc.UserService.Update(g, createDriverUserInTheSameCompanyResponse.ID, patchUserRequest)
							So(patchUserResponse, ShouldNotBeNil)
							So(patchUserResponse.FirstName, ShouldEqual, patchUserRequest.FirstName.String)
							So(err, ShouldBeNil)
						})

						Convey("User Driver can patch user Driver", func() {
							g.Set("session", &userDriverInTheSameCompanySession)
							patchUserRequest := fixtures.UpdateUser()

							patchUserResponse, err := wire.Svc.UserService.Update(g, createDriverUserInTheSameCompanyResponse.ID, patchUserRequest)
							So(patchUserResponse, ShouldNotBeNil)
							So(patchUserResponse.FirstName, ShouldEqual, patchUserRequest.FirstName.String)
							So(err, ShouldBeNil)
						})
					})
				})
			})
		})
	})
}
