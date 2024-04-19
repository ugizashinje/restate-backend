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

func TestRepair(t *testing.T) {
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

		Convey("Logged user can create repair", func() {
			createRepairRequest := fixtures.CreateRepair()

			address := []model.Address{}
			addresstRepo := repo.Repo[model.Address]{}

			result := addresstRepo.Search(g, &address, map[string]any{})
			ShouldBeNil(result.Error)

			warrant := []model.Warrant{}
			warrantRepo := repo.Repo[model.Warrant]{}

			result1 := warrantRepo.Search(g, &warrant, map[string]any{"company_id": sessions.UserManagerSession.User.Companies[0].CompanyID})
			ShouldBeNil(result1.Error)

			createRepairRequest.WarrantID = warrant[0].ID

			createRepairResponse, err := wire.Svc.RepairService.Create(g, createRepairRequest)
			So(createRepairResponse, ShouldNotBeNil)
			So(err, ShouldBeNil)

			Convey("Logged user can get repair by id", func() {
				getRepairByIdResponse, err := wire.Svc.RepairService.GetById(g, createRepairResponse.ID)

				So(getRepairByIdResponse, ShouldNotBeNil)
				So(err, ShouldBeNil)
				So(getRepairByIdResponse.ID, ShouldEqual, createRepairResponse.ID)
			})

			Convey("Logged user can update repair", func() {
				patchRepairRequest := fixtures.UpdateRepair()

				patchRepairResponse, err := wire.Svc.RepairService.Update(g, createRepairResponse.ID, patchRepairRequest)

				So(patchRepairResponse, ShouldNotBeNil)
				So(err, ShouldBeNil)
				So(patchRepairResponse.Workshop, ShouldEqual, patchRepairRequest.Workshop.String)
			})

			Convey("Logged user can search repair", func() {
				g.Request.URL, _ = url.Parse("/v1/repairs/?warrantId=" + createRepairResponse.WarrantID)

				searchResponse, err := wire.Svc.RepairService.Search(g, map[string]interface{}{"start": createRepairResponse.Start})

				So(searchResponse, ShouldNotBeNil)
				So(err, ShouldBeNil)
				So(searchResponse[0].Start, ShouldEqual, createRepairResponse.Start)
			})

			Convey("Logged user can delete repair", func() {
				deleteRepairResponse, err := wire.Svc.RepairService.Delete(g, createRepairResponse.ID)

				So(deleteRepairResponse, ShouldNotBeNil)
				So(deleteRepairResponse.DeletedAt.Time.Add(time.Hour).Format("2006-01-02 15:04:00"), ShouldEqual, time.Now().Format("2006-01-02 15:04:00"))
				So(err, ShouldBeNil)
			})
		})
	})
}
