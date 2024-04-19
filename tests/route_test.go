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

func TestRoute(t *testing.T) {
	g := SetupTest()
	sessions := UserSessions()
	g.Set("session", &sessions.UserManagerSession)
	g.Request.URL, _ = url.Parse("v1/auth/confirm/")

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

		Convey("Logged user can create route", func() {
			createRouteRequest := fixtures.CreateRoute()

			address := []model.Address{}
			addresstRepo := repo.Repo[model.Address]{}

			result := addresstRepo.Search(g, &address, map[string]any{})
			ShouldBeNil(result.Error)

			warrant := []model.Warrant{}
			warrantRepo := repo.Repo[model.Warrant]{}

			g.Request.URL, _ = url.Parse("v1/warrants/")

			result1 := warrantRepo.Search(g, &warrant, map[string]interface{}{"company_id": createWarrantResponse.CompanyID})
			ShouldBeNil(result1.Error)

			createRouteRequest.StartAddressID = address[0].ID
			createRouteRequest.EndAddressID = address[0].ID
			createRouteRequest.WarrantID = warrant[0].ID

			createRouteResponse, err := wire.Svc.RouteService.Create(g, createRouteRequest)
			So(createRouteResponse, ShouldNotBeNil)
			So(err, ShouldBeNil)

			Convey("Logged user can get route by id", func() {
				getRouteByIdResponse, err := wire.Svc.RouteService.GetById(g, createRouteResponse.ID)

				So(getRouteByIdResponse, ShouldNotBeNil)
				So(err, ShouldBeNil)
				So(getRouteByIdResponse.ID, ShouldEqual, createRouteResponse.ID)
			})

			Convey("User manager from different company can`t get route by id", func() {
				g.Set("session", sessions.UserManagerWithDifferentCompanySession)

				So(func() { wire.Svc.RouteService.GetById(g, createRouteResponse.ID) }, ShouldPanic)
			})

			Convey("User driver can`t get route by id", func() {
				g.Set("session", sessions.UserDispatcherFromDifferentCompanySession)

				So(func() { wire.Svc.RouteService.GetById(g, createRouteResponse.ID) }, ShouldPanic)
			})

			Convey("Logged user can update route", func() {
				patchRouteRequest := fixtures.UpdateRoute()

				patchRouteResponse, err := wire.Svc.RouteService.Update(g, createRouteResponse.ID, patchRouteRequest)

				So(patchRouteResponse, ShouldNotBeNil)
				So(err, ShouldBeNil)
				//wait for route statuses
				// So(patchRouteResponse.Status, ShouldEqual, enum.Confirmed)
			})

			Convey("User Manger from different company can`t update route", func() {
				patchRouteRequest := fixtures.UpdateRoute()

				g.Set("session", sessions.UserManagerWithDifferentCompanySession)

				So(func() { wire.Svc.RouteService.Update(g, createRouteResponse.ID, patchRouteRequest) }, ShouldPanic)
			})

			Convey("User Driver can`t update route", func() {
				patchRouteRequest := fixtures.UpdateRoute()

				g.Set("session", sessions.UserDispatcherFromDifferentCompanySession)

				So(func() { wire.Svc.RouteService.Update(g, createRouteResponse.ID, patchRouteRequest) }, ShouldPanic)
			})

			Convey("Logged user can search route", func() {
				g.Request.URL, _ = url.Parse("/v1/routes/?warrantId=" + createRouteResponse.WarrantID)

				searchResponse, err := wire.Svc.RouteService.Search(g, map[string]interface{}{"start_time": createRouteResponse.StartTime})

				So(searchResponse, ShouldNotBeNil)
				So(err, ShouldBeNil)
				So(searchResponse[0].StartTime, ShouldEqual, createRouteResponse.StartTime)
			})

			Convey("User Manger from different company can`t search route", func() {
				g.Request.URL, _ = url.Parse("/v1/routes/?warrantId=" + createRouteResponse.WarrantID)

				g.Set("session", sessions.UserManagerWithDifferentCompanySession)

				So(func() {
					wire.Svc.RouteService.Search(g, map[string]interface{}{"start_time": createRouteResponse.StartTime})
				}, ShouldPanic)
			})

			Convey("User Driver from can`t search route", func() {
				g.Request.URL, _ = url.Parse("/v1/routes/?warrantId=" + createRouteResponse.WarrantID)

				g.Set("session", sessions.UserDispatcherFromDifferentCompanySession)

				So(func() {
					wire.Svc.RouteService.Search(g, map[string]interface{}{"start_time": createRouteResponse.StartTime})
				}, ShouldPanic)
			})

			Convey("Logged user can delete route", func() {
				deleteRouteResponse, err := wire.Svc.RouteService.Delete(g, createRouteResponse.ID)

				So(deleteRouteResponse, ShouldNotBeNil)
				So(deleteRouteResponse.DeletedAt.Time.Add(time.Hour).Format("2006-01-02 15:04:00"), ShouldEqual, time.Now().Format("2006-01-02 15:04:00"))
				So(err, ShouldBeNil)
			})

			Convey("User Manger from different company can`t delete route", func() {
				g.Set("session", sessions.UserManagerWithDifferentCompanySession)

				So(func() { wire.Svc.RouteService.Delete(g, createRouteResponse.ID) }, ShouldPanic)
			})

			Convey("User Driver can`t delete route", func() {
				g.Set("session", sessions.UserDispatcherFromDifferentCompanySession)

				So(func() { wire.Svc.RouteService.Delete(g, createRouteResponse.ID) }, ShouldPanic)
			})
		})
	})
}
