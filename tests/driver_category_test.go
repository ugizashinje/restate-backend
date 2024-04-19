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

func TestDriverCategory(t *testing.T) {
	g := SetupTest()
	sessions := UserSessions()
	g.Set("session", &sessions.UserManagerSession)
	g.Request.URL, _ = url.Parse("v1/auth/confirm/")

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

		// DRIVER CATEGORY PASS TESTS
		// CREATE
		Convey("Logged user Manager can create driver category", func() {
			createDriverCategoryRequest := fixtures.CreateDriverCategory()

			g.Set("session", &sessions.UserManagerSession)

			createDriverCategoryRequest.UserID = sessions.UserManagerSession.User.ID

			createDriverCategoryResponse, err := wire.Svc.DriverCategoryService.Create(g, createDriverCategoryRequest)
			So(createDriverCategoryResponse, ShouldNotBeNil)
			So(err, ShouldBeNil)

			Convey("Logged user Dispather from the same company can create driver category", func() {
				g.Set("session", &sessions.UserDispatcherInTheSameCompanySession)

				createDriverCategoryRequest := fixtures.CreateDriverCategory()

				createDriverCategoryRequest.UserID = sessions.UserDispatcherInTheSameCompanySession.User.ID

				createDriverCategoryResponse, err := wire.Svc.DriverCategoryService.Create(g, createDriverCategoryRequest)
				So(createDriverCategoryResponse, ShouldNotBeNil)
				So(err, ShouldBeNil)
			})

			// SHOULD FAIL!!! CHECK
			// Convey("Logged user Driver from the same company can create driver category", func() {
			// 	g.Set("session", &userDriverInTheSameCompanySession)

			// 	createDriverCategoryRequest := fixtures.CreateDriverCategory()

			// 	createDriverCategoryRequest.UserID = users[0].ID

			// 	createDriverCategoryResponse, err := wire.Svc.DriverCategoryService.Create(g, createDriverCategoryRequest)
			// 	So(createDriverCategoryResponse, ShouldNotBeNil)
			// 	So(err, ShouldBeNil)
			// })

			// Convey("Logged user Driver from different company can create driver category", func() {
			// 	g.Set("session", &userDriverFromDifferentCompanySession)

			// 	createDriverCategoryRequest := fixtures.CreateDriverCategory()

			// 	createDriverCategoryRequest.UserID = users[0].ID

			// 	createDriverCategoryResponse, err := wire.Svc.DriverCategoryService.Create(g, createDriverCategoryRequest)
			// 	So(createDriverCategoryResponse, ShouldNotBeNil)
			// 	So(err, ShouldBeNil)
			// })

			// GET BY ID
			Convey("Logged user Manager can get driver category by id", func() {
				getDriverCategoryByIdResponse, err := wire.Svc.DriverCategoryService.GetById(g, createDriverCategoryResponse.ID)

				So(getDriverCategoryByIdResponse, ShouldNotBeNil)
				So(err, ShouldBeNil)
				So(getDriverCategoryByIdResponse.ID, ShouldEqual, createDriverCategoryResponse.ID)
			})

			Convey("Logged user Dispather from the same company can get driver category by id", func() {
				g.Set("session", &sessions.UserDispatcherInTheSameCompanySession)

				getDriverCategoryByIdResponse, err := wire.Svc.DriverCategoryService.GetById(g, createDriverCategoryResponse.ID)

				So(getDriverCategoryByIdResponse, ShouldNotBeNil)
				So(err, ShouldBeNil)
				So(getDriverCategoryByIdResponse.ID, ShouldEqual, createDriverCategoryResponse.ID)
			})

			// UPDATE
			Convey("Logged user Manager can update driver category", func() {
				patchDriverCategoryRequest := fixtures.UpdateDriverCategory()

				patchDriverCategoryResponse, err := wire.Svc.DriverCategoryService.Update(g, createDriverCategoryResponse.ID, patchDriverCategoryRequest)

				So(patchDriverCategoryResponse, ShouldNotBeNil)
				So(err, ShouldBeNil)
				So(patchDriverCategoryResponse.Category, ShouldEqual, patchDriverCategoryRequest.Category.String)
			})
			// ??
			Convey("User Dispather from the same company can update driver category", func() {
				g.Set("session", &sessions.UserDispatcherInTheSameCompanySession)

				patchDriverCategoryRequest := fixtures.UpdateDriverCategory()

				patchDriverCategoryResponse, err := wire.Svc.DriverCategoryService.Update(g, createDriverCategoryResponse.ID, patchDriverCategoryRequest)

				So(patchDriverCategoryResponse, ShouldNotBeNil)
				So(err, ShouldBeNil)
				So(patchDriverCategoryResponse.Category, ShouldEqual, patchDriverCategoryRequest.Category.String)
			})

			// Currently don`t exist
			// Convey("Logged user can search driver category", func() {
			// 	g.Request.URL, _ = url.Parse("/v1/routes/?warrantId=" + createDriverCategoryResponse.WarrantID)

			// 	searchDriverCategoryResponse, err := wire.Svc.DriverCategoryService.Search(g, map[string]interface{}{"issued": createDriverCategoryResponse.Issued})

			// 	So(searchDriverCategoryResponse, ShouldNotBeNil)
			// 	So(len(searchDriverCategoryResponse), ShouldEqual, 1)
			// 	So(err, ShouldBeNil)
			// 	So(searchDriverCategoryResponse[0].Issued, ShouldEqual, createDriverCategoryResponse.Issued)
			// })

			// DELETE
			Convey("Logged user Manager can delete driver category", func() {
				deleteDriverCategoryResponse, err := wire.Svc.DriverCategoryService.Delete(g, createDriverCategoryResponse.ID)

				So(deleteDriverCategoryResponse, ShouldNotBeNil)
				So(deleteDriverCategoryResponse.DeletedAt.Time.Add(time.Hour).Format("2006-01-02 15:04:00"), ShouldEqual, time.Now().Format("2006-01-02 15:04:00"))
				So(err, ShouldBeNil)
			})
			// ??
			Convey("user Dispather from the same company can delete driver category", func() {
				g.Set("session", &sessions.UserDispatcherInTheSameCompanySession)

				deleteDriverCategoryResponse, err := wire.Svc.DriverCategoryService.Delete(g, createDriverCategoryResponse.ID)

				So(deleteDriverCategoryResponse, ShouldNotBeNil)
				So(deleteDriverCategoryResponse.DeletedAt.Time.Add(time.Hour).Format("2006-01-02 15:04:00"), ShouldEqual, time.Now().Format("2006-01-02 15:04:00"))
				So(err, ShouldBeNil)
			})
		})
	})
}
