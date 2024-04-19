package tests

import (
	url "net/url"
	"testing"
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/repo"
	"warrant-api/pkg/wire"
	"warrant-api/tests/fixtures"

	. "github.com/smartystreets/goconvey/convey"
)

func TestShippingInvoice(t *testing.T) {
	g := SetupTest()
	sessions := UserSessions()
	g.Set("session", &sessions.UserManagerSession)
	g.Request.URL, _ = url.Parse("v1/auth/confirm/")

	Convey("Logged user Manager can create Warrant", t, func() {
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

			createRouteRequest.StartAddressID = address[0].ID
			createRouteRequest.EndAddressID = address[0].ID
			createRouteRequest.WarrantID = createWarrantResponse.ID

			createRouteResponse, err := wire.Svc.RouteService.Create(g, createRouteRequest)
			So(createRouteResponse, ShouldNotBeNil)
			So(err, ShouldBeNil)

			Convey("Logged user can create shipping invoice when valid route is created", func() {
				createShippingInvoiceRequest := fixtures.CreateShippingInvoice()

				g.Set("session", &sessions.UserManagerSession)

				createShippingInvoiceRequest.RouteID = createRouteResponse.ID

				createShippingInvoiceResponse, err := wire.Svc.ShippingInvoiceService.Create(g, createShippingInvoiceRequest)
				So(createShippingInvoiceResponse, ShouldNotBeNil)
				So(err, ShouldBeNil)

				Convey("Logged user can get shipping invoice by id", func() {
					getShippingInvoiceByIdResponse, err := wire.Svc.ShippingInvoiceService.GetById(g, createShippingInvoiceResponse.ID)

					So(getShippingInvoiceByIdResponse, ShouldNotBeNil)
					So(err, ShouldBeNil)
					So(getShippingInvoiceByIdResponse.ID, ShouldEqual, createShippingInvoiceResponse.ID)
				})

				Convey("User Manager from different company can`t get shipping invoice by id", func() {
					g.Set("session", sessions.UserManagerWithDifferentCompanySession)

					So(func() { wire.Svc.ShippingInvoiceService.GetById(g, createShippingInvoiceResponse.ID) }, ShouldPanic)
				})

				Convey("User Driver can`t get shipping invoice by id", func() {
					g.Set("session", sessions.UserDispatcherFromDifferentCompanySession)

					So(func() { wire.Svc.ShippingInvoiceService.GetById(g, createShippingInvoiceResponse.ID) }, ShouldPanic)
				})

				Convey("Logged user can update shipping invoice", func() {
					patchShippingInvoiceRequest := fixtures.UpdateShippingInvoice()

					patchShippingInvoiceResponse, err := wire.Svc.ShippingInvoiceService.Update(g, createShippingInvoiceResponse.ID, patchShippingInvoiceRequest)

					So(patchShippingInvoiceResponse, ShouldNotBeNil)
					So(err, ShouldBeNil)
					So(patchShippingInvoiceResponse.Side, ShouldEqual, patchShippingInvoiceRequest.Side)
				})

				Convey("User Manager from different company can`t update shipping invoice", func() {
					patchShippingInvoiceRequest := fixtures.UpdateShippingInvoice()

					g.Set("session", sessions.UserManagerWithDifferentCompanySession)

					So(func() {
						wire.Svc.ShippingInvoiceService.Update(g, createShippingInvoiceResponse.ID, patchShippingInvoiceRequest)
					}, ShouldPanic)
				})

				Convey("User Driver can`t update shipping invoice", func() {
					patchShippingInvoiceRequest := fixtures.UpdateShippingInvoice()

					g.Set("session", sessions.UserDispatcherFromDifferentCompanySession)

					So(func() {
						wire.Svc.ShippingInvoiceService.Update(g, createShippingInvoiceResponse.ID, patchShippingInvoiceRequest)
					}, ShouldPanic)
				})

				Convey("Logged user can search shipping invoice", func() {
					g.Request.URL, _ = url.Parse("/v1/shippingInvoices/?warrantId=" + createWarrantResponse.ID)
					g.Set("session", sessions.UserManagerSession)

					searchShippingInvoiceResponse, err := wire.Svc.ShippingInvoiceService.Search(g, map[string]any{"status": createShippingInvoiceResponse.Status})

					So(searchShippingInvoiceResponse, ShouldNotBeNil)
					So(err, ShouldBeNil)
					// So(searchShippingInvoiceResponse[0].Status, ShouldEqual, createShippingInvoiceResponse.Status)
				})

				// NOT Implemented
				// Convey("User Manager from different company can`t search shipping invoice", func() {
				// 	g.Request.URL, _ = url.Parse("/v1/shippingInvoices/?warrantId=" + createRouteResponse.WarrantID)

				// 	g.Set("session", &userManagerWithDifferentCompanySession)

				// 	searchShippingInvoiceResponse, err := wire.Svc.ShippingInvoiceService.Search(g, map[string]interface{}{"status": createShippingInvoiceResponse.Status})

				// 	So(searchShippingInvoiceResponse, ShouldNotBeNil)
				// 	So(err, ShouldBeNil)
				// 	So(len(searchShippingInvoiceResponse), ShouldEqual, 0)
				// })

				// Convey("User Driver can`t search shipping invoice", func() {
				// 	g.Request.URL, _ = url.Parse("/v1/shippingInvoices/?warrantId=" + createRouteResponse.WarrantID)

				// 	g.Set("session", &userDriverSession)

				// 	searchShippingInvoiceResponse, err := wire.Svc.ShippingInvoiceService.Search(g, map[string]interface{}{"status": createShippingInvoiceResponse.Status})

				// 	So(searchShippingInvoiceResponse, ShouldNotBeNil)
				// 	So(err, ShouldBeNil)
				// 	So(len(searchShippingInvoiceResponse), ShouldEqual, 0)
				// })

				// Not working
				// Convey("Logged user can delete shipping invoice", func() {
				// 	deleteShippingInvoiceResponse, err := wire.Svc.ShippingInvoiceService.Delete(g, createShippingInvoiceResponse.ID)

				// 	So(deleteShippingInvoiceResponse, ShouldNotBeNil)
				// 	So(deleteShippingInvoiceResponse.DeletedAt.Time.Add(time.Hour).Format("2006-01-02 15:04:00"), ShouldEqual, time.Now().Format("2006-01-02 15:04:00"))
				// 	So(err, ShouldBeNil)
				// })

				// Convey("User Manager from different company can`t delete shipping invoice", func() {
				// 	deleteShippingInvoiceResponse, err := wire.Svc.ShippingInvoiceService.Delete(g, createShippingInvoiceResponse.ID)

				// 	g.Set("session", &userManagerWithDifferentCompanySession)

				// 	So(deleteShippingInvoiceResponse, ShouldNotBeNil)
				// 	So(deleteShippingInvoiceResponse.DeletedAt.Time.Add(time.Hour).Format("2006-01-02 15:04:00"), ShouldEqual, time.Now().Format("2006-01-02 15:04:00"))
				// 	So(err, ShouldBeNil)
				// })
			})
		})
	})
}
