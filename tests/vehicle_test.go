package tests

import (
	url "net/url"
	"testing"
	"time"
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/repo"
	"warrant-api/pkg/wire"
	"warrant-api/tests/fixtures"

	. "github.com/smartystreets/goconvey/convey"
)

func TestVehicle(t *testing.T) {
	g := SetupTest()
	sessions := UserSessions()
	g.Set("session", &sessions.UserManagerSession)
	g.Request.URL, _ = url.Parse("v1/auth/confirm/")

	// VEHICLE PASS TESTS
	// CREATE
	Convey("Logged user Manager can create Vehicle", t, func() {
		registerRequest := fixtures.RegisterRegularCompany()

		g.Set("session", &sessions.UserManagerSession)
		g.Request.URL, _ = url.Parse("/v1/addresses")

		address := []model.Address{}

		addressRepo := repo.Repo[model.Address]{}

		result := addressRepo.Search(g, &address, map[string]any{"street": registerRequest.UserAddressStreet})
		ShouldBeNil(result.Error)

		createVehicleRequest := fixtures.CreateVehicle()

		createVehicleRequest.OwnerAddressID = address[0].ID
		createVehicleRequest.UserAddressID = address[0].ID
		createVehicleRequest.CompanyID = sessions.UserManagerSession.User.Companies[0].CompanyID

		createVehicleResponse, err := wire.Svc.VehicleService.Create(g, createVehicleRequest)
		So(createVehicleResponse, ShouldNotEqual, nil)
		So(err, ShouldBeNil)

		Convey("Logged user Dispatcher from the same company can create Vehicle", func() {
			address := []model.Address{}
			addressRepo := repo.Repo[model.Address]{}

			result := addressRepo.Search(g, &address, map[string]any{"street": registerRequest.UserAddressStreet})
			ShouldBeNil(result.Error)

			createVehicleRequest := fixtures.CreateVehicle()

			createVehicleRequest.OwnerAddressID = address[0].ID
			createVehicleRequest.UserAddressID = address[0].ID
			createVehicleRequest.CompanyID = sessions.UserManagerSession.User.Companies[0].CompanyID

			g.Set("session", &sessions.UserDispatcherInTheSameCompanySession)

			createVehicleResponse, _ := wire.Svc.VehicleService.Create(g, createVehicleRequest)
			So(createVehicleResponse, ShouldNotEqual, nil)
			So(err, ShouldBeNil)
		})

		// GET BY ID
		Convey("Logged user Manager can get vehicle by id", func() {
			getVehicleByIdResponse, err := wire.Svc.VehicleService.GetById(g, createVehicleResponse.ID)

			So(getVehicleByIdResponse, ShouldNotBeNil)
			So(err, ShouldBeNil)
			So(getVehicleByIdResponse.ID, ShouldEqual, createVehicleResponse.ID)
		})

		Convey("Logged user Dispather from the same company can get vehicle by id", func() {
			g.Set("session", &sessions.UserDispatcherInTheSameCompanySession)

			getVehicleByIdResponse, err := wire.Svc.VehicleService.GetById(g, createVehicleResponse.ID)

			So(getVehicleByIdResponse, ShouldNotBeNil)
			So(err, ShouldBeNil)
			So(getVehicleByIdResponse.ID, ShouldEqual, createVehicleResponse.ID)
		})

		// UPDATE
		Convey("Logged user Manager can update vehicle", func() {
			patchVehicleRequest := fixtures.UpdateVehicle()

			patchVehicleResponse, err := wire.Svc.VehicleService.Update(g, createVehicleResponse.ID, patchVehicleRequest)

			So(patchVehicleResponse, ShouldNotBeNil)
			So(err, ShouldBeNil)
			So(patchVehicleResponse.Status, ShouldEqual, patchVehicleRequest.Status.String)
		})

		Convey("Logged user Dispatcher from the same company can update vehicle", func() {
			patchVehicleRequest := fixtures.UpdateVehicle()

			g.Set("session", &sessions.UserDispatcherInTheSameCompanySession)

			patchVehicleResponse, err := wire.Svc.VehicleService.Update(g, createVehicleResponse.ID, patchVehicleRequest)

			So(patchVehicleResponse, ShouldNotBeNil)
			So(err, ShouldBeNil)
			So(patchVehicleResponse.Status, ShouldEqual, patchVehicleRequest.Status.String)
		})

		// SEARCH
		Convey("Logged user Manager can search vehicle", func() {
			searchVehicleResponse, err := wire.Svc.VehicleService.Search(g, map[string]interface{}{"owner_mn": createVehicleRequest.OwnerMn})

			So(searchVehicleResponse, ShouldNotBeNil)
			So(err, ShouldBeNil)
			So(searchVehicleResponse[0].OwnerMn, ShouldEqual, createVehicleRequest.OwnerMn)
		})

		Convey("Logged user Dispatcher from the same company can search vehicle", func() {
			g.Set("session", &sessions.UserDispatcherInTheSameCompanySession)

			searchVehicleResponse, err := wire.Svc.VehicleService.Search(g, map[string]interface{}{"owner_mn": createVehicleRequest.OwnerMn})

			So(searchVehicleResponse, ShouldNotBeNil)
			So(err, ShouldBeNil)
			So(searchVehicleResponse[0].OwnerMn, ShouldEqual, createVehicleRequest.OwnerMn)
		})

		// DELETE
		Convey("Logged user Manager can delete vehicle", func() {
			deleteVehicleResponse, err := wire.Svc.VehicleService.Delete(g, createVehicleResponse.ID)

			So(deleteVehicleResponse, ShouldNotBeNil)
			So(deleteVehicleResponse.DeletedAt.Time.Add(time.Hour).Format("2006-01-02 15:04:00"), ShouldEqual, time.Now().Format("2006-01-02 15:04:00"))
			So(err, ShouldBeNil)
		})

		Convey("Logged user Dispatcher from the same company can delete vehicle", func() {
			g.Set("session", &sessions.UserDispatcherInTheSameCompanySession)

			deleteVehicleResponse, err := wire.Svc.VehicleService.Delete(g, createVehicleResponse.ID)

			So(deleteVehicleResponse, ShouldNotBeNil)
			So(deleteVehicleResponse.DeletedAt.Time.Add(time.Hour).Format("2006-01-02 15:04:00"), ShouldEqual, time.Now().Format("2006-01-02 15:04:00"))
			So(err, ShouldBeNil)
		})

		// VEHICLE FAIL TESTS
		// CREATE
		Convey("User Dispatcher can`t create Vehicle if he is in different company", func() {
			address := []model.Address{}
			addressRepo := repo.Repo[model.Address]{}

			result := addressRepo.Search(g, &address, map[string]any{"street": registerRequest.UserAddressStreet})
			ShouldBeNil(result.Error)

			createVehicleRequest := fixtures.CreateVehicle()

			createVehicleRequest.OwnerAddressID = address[0].ID
			createVehicleRequest.UserAddressID = address[0].ID
			createVehicleRequest.CompanyID = sessions.UserManagerSession.User.Companies[0].CompanyID

			g.Set("session", &sessions.UserDispatcherFromDifferentCompanySession)

			So(func() { wire.Svc.VehicleService.Create(g, createVehicleRequest) }, ShouldPanic)
		})

		Convey("User Driver from the same company can`t create Vehicle", func() {
			address := []model.Address{}
			addressRepo := repo.Repo[model.Address]{}

			result := addressRepo.Search(g, &address, map[string]any{"street": registerRequest.UserAddressStreet})
			ShouldBeNil(result.Error)

			company := []model.Company{}
			companyRepo := repo.Repo[model.Company]{}

			result1 := companyRepo.Search(g, &company, map[string]any{"pib": registerRequest.CompanyPIB})
			ShouldBeNil(result1.Error)

			createVehicleRequest := fixtures.CreateVehicle()

			createVehicleRequest.OwnerAddressID = address[0].ID
			createVehicleRequest.UserAddressID = address[0].ID
			createVehicleRequest.CompanyID = sessions.UserManagerSession.User.Companies[0].CompanyID

			g.Set("session", &sessions.UserDriverInTheSameCompanySession)

			So(func() { wire.Svc.VehicleService.Create(g, createVehicleRequest) }, ShouldPanic)
		})

		Convey("User Driver from different company can`t create Vehicle", func() {
			address := []model.Address{}
			addressRepo := repo.Repo[model.Address]{}

			result := addressRepo.Search(g, &address, map[string]any{"street": registerRequest.UserAddressStreet})
			ShouldBeNil(result.Error)

			createVehicleRequest := fixtures.CreateVehicle()

			createVehicleRequest.OwnerAddressID = address[0].ID
			createVehicleRequest.UserAddressID = address[0].ID
			createVehicleRequest.CompanyID = sessions.UserManagerSession.User.Companies[0].CompanyID

			g.Set("session", &sessions.UserDriverFromDifferentCompanySession)

			So(func() { wire.Svc.VehicleService.Create(g, createVehicleRequest) }, ShouldPanic)
		})

		// GET BY ID
		Convey("User Manager from differrent company can`t get vehicle by id", func() {
			g.Set("session", &sessions.UserManagerWithDifferentCompanySession)

			So(func() { wire.Svc.VehicleService.GetById(g, createVehicleResponse.ID) }, ShouldPanic)
		})

		Convey("User Dispatcher from the different company can`t get vehicle by id", func() {
			g.Set("session", &sessions.UserDispatcherFromDifferentCompanySession)

			So(func() { wire.Svc.VehicleService.GetById(g, createVehicleResponse.ID) }, ShouldPanic)
		})

		Convey("User Driver from the same company can`t get vehicle by id", func() {
			g.Set("session", &sessions.UserDriverInTheSameCompanySession)

			So(func() { wire.Svc.VehicleService.GetById(g, createVehicleResponse.ID) }, ShouldPanic)
		})

		Convey("User Driver from the different company can`t get vehicle by id", func() {
			g.Set("session", &sessions.UserDriverFromDifferentCompanySession)

			So(func() { wire.Svc.VehicleService.GetById(g, createVehicleResponse.ID) }, ShouldPanic)
		})

		// UPDATE
		Convey("User Manager from differrent company can`t update vehicle", func() {
			patchVehicleRequest := fixtures.UpdateVehicle()

			g.Set("session", &sessions.UserManagerWithDifferentCompanySession)

			So(func() { wire.Svc.VehicleService.Update(g, createVehicleResponse.ID, patchVehicleRequest) }, ShouldPanic)
		})

		Convey("User Dispather from different company can`t update vehicle", func() {
			patchVehicleRequest := fixtures.UpdateVehicle()

			g.Set("session", &sessions.UserDispatcherFromDifferentCompanySession)

			So(func() { wire.Svc.VehicleService.Update(g, createVehicleResponse.ID, patchVehicleRequest) }, ShouldPanic)
		})

		Convey("User Driver from the same company can`t update vehicle", func() {
			patchVehicleRequest := fixtures.UpdateVehicle()

			g.Set("session", &sessions.UserDriverInTheSameCompanySession)

			So(func() { wire.Svc.VehicleService.Update(g, createVehicleResponse.ID, patchVehicleRequest) }, ShouldPanic)
		})

		Convey("User Driver from different company can`t update vehicle", func() {
			patchVehicleRequest := fixtures.UpdateVehicle()

			g.Set("session", &sessions.UserDriverFromDifferentCompanySession)

			So(func() { wire.Svc.VehicleService.Update(g, createVehicleResponse.ID, patchVehicleRequest) }, ShouldPanic)
		})

		// SEARCH
		Convey("User Manager from different company can`t search vehicle", func() {
			g.Set("session", &sessions.UserManagerWithDifferentCompanySession)

			searchVehicleResponse, err := wire.Svc.VehicleService.Search(g, map[string]interface{}{"owner_mn": createVehicleRequest.OwnerMn})

			So(searchVehicleResponse, ShouldNotBeNil)
			So(len(searchVehicleResponse), ShouldEqual, 0)
			So(err, ShouldBeNil)
		})

		Convey("User Dispatcher from different company can`t search vehicle", func() {
			g.Set("session", &sessions.UserDispatcherFromDifferentCompanySession)

			searchVehicleResponse, err := wire.Svc.VehicleService.Search(g, map[string]interface{}{"owner_mn": createVehicleRequest.OwnerMn})

			So(searchVehicleResponse, ShouldNotBeNil)
			So(len(searchVehicleResponse), ShouldEqual, 0)
			So(err, ShouldBeNil)
		})

		Convey("User Driver from the same company can`t search vehicle", func() {
			g.Set("session", &sessions.UserDriverInTheSameCompanySession)

			searchVehicleResponse, err := wire.Svc.VehicleService.Search(g, map[string]interface{}{"owner_mn": createVehicleRequest.OwnerMn})

			So(searchVehicleResponse, ShouldNotBeNil)
			So(len(searchVehicleResponse), ShouldEqual, 0)
			So(err, ShouldBeNil)
		})

		Convey("User Driver from different company can`t search vehicle", func() {
			g.Set("session", &sessions.UserDriverFromDifferentCompanySession)

			searchVehicleResponse, err := wire.Svc.VehicleService.Search(g, map[string]interface{}{"owner_mn": createVehicleRequest.OwnerMn})

			So(searchVehicleResponse, ShouldNotBeNil)
			So(len(searchVehicleResponse), ShouldEqual, 0)
			So(err, ShouldBeNil)
		})

		// DELETE
		Convey("User Manager from different company can`t delete vehicle", func() {
			g.Set("session", &sessions.UserManagerWithDifferentCompanySession)

			So(func() { wire.Svc.VehicleService.Delete(g, createVehicleResponse.ID) }, ShouldPanic)
		})

		Convey("User Dispather from different company can`t delete vehicle", func() {
			g.Set("session", &sessions.UserDispatcherFromDifferentCompanySession)

			So(func() { wire.Svc.VehicleService.Delete(g, createVehicleResponse.ID) }, ShouldPanic)
		})

		Convey("User Driver from the same company can`t delete vehicle", func() {
			g.Set("session", &sessions.UserDriverInTheSameCompanySession)

			So(func() { wire.Svc.VehicleService.Delete(g, createVehicleResponse.ID) }, ShouldPanic)
		})

		Convey("User Driver from different company can`t delete vehicle", func() {
			g.Set("session", &sessions.UserDriverFromDifferentCompanySession)

			So(func() { wire.Svc.VehicleService.Delete(g, createVehicleResponse.ID) }, ShouldPanic)
		})
	})
}
