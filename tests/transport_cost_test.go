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

func TestTransportCost(t *testing.T) {
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

		Convey("Logged user can create transport cost", func() {
			createTransportCostRequest := fixtures.CreateTransportCost()

			warrant := []model.Warrant{}
			warrantRepo := repo.Repo[model.Warrant]{}

			result1 := warrantRepo.Search(g, &warrant, map[string]any{"company_id": sessions.UserManagerSession.User.Companies[0].CompanyID})
			ShouldBeNil(result1.Error)

			createTransportCostRequest.WarrantID = warrant[0].ID

			g.Set("session", &sessions.UserManagerSession)

			createTransportCostResponse, err := wire.Svc.TransportCostService.Create(g, createTransportCostRequest)
			So(createTransportCostResponse, ShouldNotBeNil)
			So(err, ShouldBeNil)

			Convey("Logged user can get transport cost by id", func() {
				getTransportCostByIdResponse, err := wire.Svc.TransportCostService.GetById(g, createTransportCostResponse.ID)

				So(getTransportCostByIdResponse, ShouldNotBeNil)
				So(err, ShouldBeNil)
				So(getTransportCostByIdResponse.ID, ShouldEqual, createTransportCostResponse.ID)
			})

			Convey("Logged user can update transport cost", func() {
				patchTransportCostRequest := fixtures.UpdateTransportCost()

				patchTransportCostResponse, err := wire.Svc.TransportCostService.Update(g, createTransportCostResponse.ID, patchTransportCostRequest)

				So(patchTransportCostResponse, ShouldNotBeNil)
				So(err, ShouldBeNil)
				So(patchTransportCostResponse.Code.String, ShouldEqual, patchTransportCostRequest.Code.String)
			})

			Convey("Logged user can search transport cost", func() {
				g.Request.URL, _ = url.Parse("/v1/transportCosts/?warrantId=" + createTransportCostResponse.WarrantID)

				searchTransportCostResponse, err := wire.Svc.TransportCostService.Search(g, map[string]interface{}{"code": createTransportCostResponse.Code})

				So(searchTransportCostResponse, ShouldNotBeNil)
				So(err, ShouldBeNil)
				So(searchTransportCostResponse[0].Code, ShouldEqual, createTransportCostResponse.Code)
			})

			Convey("Logged user can delete transport cost", func() {
				deleteTransportCostResponse, err := wire.Svc.TransportCostService.Delete(g, createTransportCostResponse.ID)

				So(deleteTransportCostResponse, ShouldNotBeNil)
				So(deleteTransportCostResponse.DeletedAt.Time.Add(time.Hour).Format("2006-01-02 15:04:00"), ShouldEqual, time.Now().Format("2006-01-02 15:04:00"))
				So(err, ShouldBeNil)
			})
		})
	})
}
