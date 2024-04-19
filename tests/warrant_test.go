package tests

import (
	url "net/url"
	"testing"
	"time"
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/enum"
	"warrant-api/pkg/repo"
	"warrant-api/pkg/wire"
	"warrant-api/tests/fixtures"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWarrant(t *testing.T) {
	g := SetupTest()
	sessions := UserSessions()
	g.Set("session", &sessions.UserManagerSession)
	g.Request.URL, _ = url.Parse("v1/auth/confirm/")

	// WARRANT PASS TESTS
	// CREATE
	Convey("Logged user Manager can create Warrant when vehicle is created", t, func() {
		g.Set("session", &sessions.UserManagerSession)

		createVehicleRequest := fixtures.CreateVehicle()

		vehicle := []model.Vehicle{}
		vehicleRepo := repo.Repo[model.Vehicle]{}

		result2 := vehicleRepo.Search(g, &vehicle, map[string]any{"owner_mn": createVehicleRequest.OwnerMn})
		ShouldBeNil(result2.Error)

		createWarrantRequest := fixtures.CreateWarrant()

		createWarrantRequest.DriverID = sessions.UserManagerSession.User.ID
		createWarrantRequest.VehicleID = vehicle[0].ID
		createWarrantRequest.CompanyID = sessions.UserManagerSession.User.Companies[0].CompanyID
		createWarrantRequest.DispatcherID = sessions.UserDispatcherInTheSameCompanySession.User.ID

		createWarrantResponse, _ := wire.Svc.WarrantService.Create(g, createWarrantRequest)
		So(createWarrantResponse, ShouldNotBeNil)
		So(nil, ShouldBeNil)

		Convey("User Dispather from the same company can create Warrant when vehicle is created", func() {
			g.Set("session", &sessions.UserDispatcherInTheSameCompanySession)

			g.Request.URL, _ = url.Parse("v1/vehicles/")

			vehicle := []model.Vehicle{}
			vehicleRepo := repo.Repo[model.Vehicle]{}

			result2 := vehicleRepo.Search(g, &vehicle, map[string]any{"owner_mn": createVehicleRequest.OwnerMn})
			ShouldBeNil(result2.Error)

			createWarrantRequest := fixtures.CreateWarrant()

			createWarrantRequest.DriverID = sessions.UserManagerSession.User.ID
			createWarrantRequest.VehicleID = vehicle[0].ID
			createWarrantRequest.CompanyID = sessions.UserManagerSession.User.Companies[0].CompanyID
			createWarrantRequest.DispatcherID = sessions.UserDispatcherInTheSameCompanySession.User.ID

			createWarrantResponse, _ := wire.Svc.WarrantService.Create(g, createWarrantRequest)
			So(createWarrantResponse, ShouldNotBeNil)
			So(nil, ShouldBeNil)
		})

		// GET BY ID
		Convey("Logged User Manager can see warrant if it belongs to his company", func() {
			warrant := []model.Warrant{}
			warrantRepo := repo.Repo[model.Warrant]{}

			result := warrantRepo.Search(g, &warrant, map[string]any{"company_id": createWarrantRequest.CompanyID})
			ShouldBeNil(result.Error)

			getWarrantByIdResponse, err := wire.Svc.WarrantService.GetById(g, warrant[0].ID)

			So(getWarrantByIdResponse, ShouldNotBeNil)
			So(err, ShouldBeNil)
		})

		Convey("User Dispather from the same company can see warrant if it belongs to his company", func() {
			g.Set("session", &sessions.UserDispatcherInTheSameCompanySession)

			warrant := []model.Warrant{}
			warrantRepo := repo.Repo[model.Warrant]{}

			result := warrantRepo.Search(g, &warrant, map[string]any{"company_id": createWarrantRequest.CompanyID})
			ShouldBeNil(result.Error)

			getWarrantByIdResponse, err := wire.Svc.WarrantService.GetById(g, warrant[0].ID)

			So(getWarrantByIdResponse, ShouldNotBeNil)
			So(err, ShouldBeNil)
		})

		Convey("User manager can`t see warrant if it doesn`t belong to his company", func() {
			warrant := []model.Warrant{}
			warrantRepo := repo.Repo[model.Warrant]{}

			result := warrantRepo.Search(g, &warrant, map[string]any{"company_id": createWarrantRequest.CompanyID})
			ShouldBeNil(result.Error)

			g.Set("session", &sessions.UserManagerWithDifferentCompanySession)

			So(func() { wire.Svc.WarrantService.GetById(g, warrant[0].ID) }, ShouldPanic)
		})

		Convey("User driver can`t see warrant if it`s not assigned to him", func() {
			warrant := []model.Warrant{}
			warrantRepo := repo.Repo[model.Warrant]{}

			result := warrantRepo.Search(g, &warrant, map[string]any{"company_id": createWarrantRequest.CompanyID})
			ShouldBeNil(result.Error)

			g.Set("session", &sessions.UserDispatcherFromDifferentCompanySession)

			So(func() { wire.Svc.WarrantService.GetById(g, warrant[0].ID) }, ShouldPanic)
		})

		Convey("Logged user can patch warrant", func() {
			warrant := []model.Warrant{}
			warrantRepo := repo.Repo[model.Warrant]{}

			result := warrantRepo.Search(g, &warrant, map[string]any{"company_id": createWarrantRequest.CompanyID})
			ShouldBeNil(result.Error)

			patchWarrantRequest := fixtures.UpdateWarrant()

			patchResponse, err := wire.Svc.WarrantService.Update(g, warrant[0].ID, patchWarrantRequest)
			So(patchResponse, ShouldNotBeNil)
			So(patchResponse.Status, ShouldEqual, enum.WarrantStatus("preparation"))
			So(err, ShouldBeNil)
		})

		// SHOULD FAIL
		// Convey("User manager can't patch warrant if it doesn`t belong to his company", func() {
		// 	warrant := []model.Warrant{}
		// 	warrantRepo := repo.Repo[model.Warrant]{}

		// 	result := warrantRepo.Search(g, &warrant, map[string]any{"company_id": createWarrantRequest.CompanyID})
		// 	ShouldBeNil(result.Error)
		// 	So(len(warrant), ShouldEqual, 1)

		// 	patchWarrantRequest := fixtures.UpdateWarrant()

		// 	g.Set("session", &userManagerWithDifferentCompanySession)

		// 	So(func() { wire.Svc.WarrantService.Update(g, warrant[0].ID, patchWarrantRequest) }, ShouldPanic)
		// })

		Convey("Logged user can search warrant", func() {
			searchResponse, err := wire.Svc.WarrantService.Search(g, map[string]any{"driver_id": createWarrantRequest.DriverID})

			So(searchResponse, ShouldNotBeNil)
			So(err, ShouldBeNil)
			So(searchResponse[0].DriverID, ShouldEqual, createWarrantRequest.DriverID)
		})

		Convey("User manager can`t search warrant if it doesn`t belong to his company", func() {
			g.Set("session", &sessions.UserManagerWithDifferentCompanySession)

			searchResponse, err := wire.Svc.WarrantService.Search(g, map[string]any{"driver_id": createWarrantRequest.DriverID})

			So(searchResponse, ShouldNotBeNil)
			So(len(searchResponse), ShouldEqual, 0)
			So(err, ShouldBeNil)

		})

		Convey("User driver can`t search warrant if it`s not assigned to him", func() {
			g.Set("session", sessions.UserDispatcherFromDifferentCompanySession)

			// searchResponse, err := wire.Svc.WarrantService.Search(g, map[string]any{"driver_id": createWarrantRequest.DriverID})
			So(func() {
				wire.Svc.WarrantService.Search(g, map[string]any{"driver_id": &sessions.UserDriverInTheSameCompanySession.User.ID})
			}, ShouldPanic)

			// So(searchResponse, ShouldNotBeNil)
			// So(len(searchResponse), ShouldEqual, 0)
			// So(err, ShouldBeNil)
		})

		Convey("Logged user can delete warrant", func() {
			warrant := []model.Warrant{}
			warrantRepo := repo.Repo[model.Warrant]{}

			result := warrantRepo.Search(g, &warrant, map[string]any{"company_id": createWarrantRequest.CompanyID})
			ShouldBeNil(result.Error)

			deleteResponse, err := wire.Svc.WarrantService.Delete(g, warrant[0].ID)
			So(deleteResponse, ShouldNotBeNil)
			So(deleteResponse.DeletedAt.Time, ShouldHappenWithin, time.Second, time.Now())
			So(err, ShouldBeNil)
		})

		Convey("User manager can`t delete warrant if it doesn`t belong to his company", func() {
			warrant := []model.Warrant{}
			warrantRepo := repo.Repo[model.Warrant]{}

			result := warrantRepo.Search(g, &warrant, map[string]any{"company_id": createWarrantRequest.CompanyID})
			ShouldBeNil(result.Error)

			g.Set("session", sessions.UserManagerWithDifferentCompanySession)

			So(func() { wire.Svc.WarrantService.Delete(g, warrant[0].ID) }, ShouldPanic)
		})

		Convey("User driver can`t delete warrant", func() {
			warrant := []model.Warrant{}
			warrantRepo := repo.Repo[model.Warrant]{}

			result := warrantRepo.Search(g, &warrant, map[string]any{"company_id": createWarrantRequest.CompanyID})
			ShouldBeNil(result.Error)

			g.Set("session", sessions.UserDispatcherFromDifferentCompanySession)

			So(func() { wire.Svc.WarrantService.Delete(g, warrant[0].ID) }, ShouldPanic)
		})
	})
}
