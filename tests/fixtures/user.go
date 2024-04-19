package fixtures

import (
	"fmt"
	"warrant-api/pkg/services"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

func RegisterRegularCompany() services.RegisterRequest {
	streetNo := null.NewString("zez", true)
	companyAddressName := null.NewString("beogradjank", true)
	userAddressStreet := null.NewString("Nehruova", true)
	userAddressStreetNo := null.NewString("80a", true)
	return services.RegisterRequest{
		CompanyAddressCity:     null.NewString("Beograd", true),
		CompanyAddressStreet:   null.NewString("Nehruova", true),
		CompanyAddressStreetNo: streetNo,
		CompanyAddressName:     companyAddressName,
		CompanyName:            "Naftna industrija srbije",
		CompanyShort:           "NIS",
		CompanyPIB:             fmt.Sprint(gofakeit.Number(10000, 1000000000)),
		CompanyAddressID:       null.String{},
		CompanyMn:              fmt.Sprint(gofakeit.Number(10000, 1000000000)),
		CompanyPhone:           gofakeit.Phone(),
		CompanyEmail:           gofakeit.Name() + "@testdomain.com",

		UserAddressCity:     null.NewString("Beograd", true),
		UserAddressStreet:   null.NewString("Nehruiova", true),
		UserAddressStreetNo: userAddressStreetNo,
		UserAddressName:     userAddressStreet,

		UserEmail:      uuid.NewString() + "@mail.com",
		UserPassword:   "asdfasdf",
		UserFirstName:  "Tester",
		UserLastName:   "Testic",
		UserMn:         fmt.Sprint(gofakeit.Number(10000, 100000)),
		UserBirthplace: "Cacak",
		UserPhone:      gofakeit.Phone(),
		UserAddressID:  null.String{},
		UserIsDriver:   false,
		UserRoles:      []string{"manager"},
	}
}

func CreateUserInSameCompany() services.CreateUserRequest {
	return services.CreateUserRequest{
		Email:      uuid.NewString() + "@mail.com",
		Password:   "asdfasdf",
		FirstName:  "InSameCompany",
		LastName:   "TestSameComp",
		Mn:         fmt.Sprint(gofakeit.Number(10000, 100000)),
		Birthplace: "Kraljevo",
		Phone:      gofakeit.Phone(),
		AddressID:  "",
		CompanyID:  "",
		IsDriver:   false,
		Roles:      []string{"dispatcher"},
	}
}

func UpdateUser() services.PatchUserRequest {
	return services.PatchUserRequest{
		Email:      null.NewString(uuid.NewString()+"@mail.com", true),
		Password:   null.NewString("asdfasdf", true),
		FirstName:  null.NewString("Updated Name", true),
		LastName:   null.NewString("Updated Last Name", true),
		Mn:         null.NewString(fmt.Sprint(gofakeit.Number(10000, 100000)), true),
		Birthplace: null.NewString("Updated Birthplace", true),
		Phone:      null.NewString("1233445566", true),
		AddressID:  null.String{},
		IsDriver:   null.NewBool(false, true),
		Roles:      []string{"dispatcher"},
	}
}

func CreateVehicle() services.CreateVehicleRequest {

	return services.CreateVehicleRequest{
		Plate:                       "New Plate",
		CompanyID:                   "",
		Model:                       "New Model",
		Producer:                    "New Producer",
		Type:                        "New Type",
		Weight:                      null.NewInt(2000, true),
		MaxWeight:                   null.NewInt(3000, true),
		Capacity:                    null.NewInt(5, true),
		HomologationMark:            "New Homologation Mark",
		MotorVolume:                 2000,
		Power:                       150,
		Fuel:                        "New Fuel",
		Consumption:                 10.5,
		ThrustWeight:                null.NewInt(15, true),
		StandingSpace:               null.NewInt(2, true),
		SeatingSpace:                null.NewInt(5, true),
		Color:                       "New Color",
		Shaft:                       2,
		Chasie:                      "New Chasie",
		Motor:                       "New Motor",
		Insurance:                   null.NewString("New Insurance", true),
		RegistrationDate:            "2023-01-01",
		FirstRegistrationDate:       null.NewString("2023-01-01", true),
		OwnerMn:                     1234567890123,
		OwnerFirstName:              "New Owner First Name",
		OwnerLastName:               "New Owner Last Name",
		OwnerAddressID:              "",
		UserFirstName:               "New User First Name",
		UserLastName:                "New User Last Name",
		UserAddressID:               "",
		UserMn:                      9876543210987,
		RegistrationCertificateDate: "2023-01-01",
		RegistrationExpiry:          "2023-01-01",
		RegisterId:                  gofakeit.Number(100, 9999),
		Status:                      "status",
	}
}

func UpdateVehicle() services.PatchVehicleRequest {

	return services.PatchVehicleRequest{
		ComapnyID:                   null.String{},
		Plate:                       null.NewString("Updated plate", true),
		Model:                       null.NewString("Updated model", true),
		Producer:                    null.NewString("Updated producer", true),
		Type:                        null.NewString("Updated type", true),
		Weight:                      null.NewInt(5000, true),
		MaxWeight:                   null.NewInt(8000, true),
		Capacity:                    null.NewInt(10, true),
		HomologationMark:            null.NewString("Updated Homologation Mark", true),
		MotorVolume:                 null.NewInt(5000, true),
		Power:                       null.NewInt(250, true),
		Fuel:                        null.NewString("Updated fuel", true),
		Consumption:                 null.NewFloat(15.7, true),
		ThrustWeight:                null.NewInt(18, true),
		StandingSpace:               null.NewInt(12, true),
		SeatingSpace:                null.NewInt(10, true),
		Color:                       null.NewString("Updated color", true),
		Shaft:                       null.NewInt(3, true),
		Chasie:                      null.NewString("54165456456", true),
		Motor:                       null.NewString("Updated motor", true),
		Insurance:                   null.NewString("Updated Insurance", true),
		RegistrationDate:            null.NewString("2023-03-03", true),
		FirstRegistrationDate:       null.NewString("2022-02-02", true),
		OwnerMn:                     null.NewInt(1554654684651654685, true),
		OwnerFirstName:              null.NewString("Updated Owner First Name", true),
		OwnerLastName:               null.NewString("Updated Owner Last Name", true),
		OwnerAddressID:              null.String{},
		UserFirstName:               null.NewString("Updated User First Name", true),
		UserLastName:                null.NewString("New User Last Name", true),
		UserAddressID:               null.String{},
		UserMn:                      null.NewInt(1111111111, true),
		RegistrationCertificateDate: null.NewString("2023-03-03", true),
		RegistrationExpiry:          null.NewString("2025-05-05", true),
		RegisterId:                  null.NewInt(7777, true),
		Status:                      null.NewString("Updated status", true),
	}
}

func CreateWarrant() services.CreateWarrantRequest {
	return services.CreateWarrantRequest{
		ExpectedStart:        "2006-01-02 15:04:05",
		DriverID:             "",
		VehicleID:            "",
		CompanyID:            "",
		DispatcherID:         "",
		TechnicalCorrectness: "ready",
		Status:               "template",
	}
}

func UpdateWarrant() services.PatchWarrantRequest {

	return services.PatchWarrantRequest{
		PrimitivePatch: services.PrimitivePatch{
			ExpectedStart:       null.Time{},
			DriverID:            null.String{},
			VehicleID:           null.String{},
			CompanyID:           null.String{},
			DispatcherID:        null.String{},
			Status:              null.NewString("preparation", true),
			TechicalCorrectness: null.NewString("updated", true),
		},
		AddPassenger: null.NewString("Updated fakeAddPassenger", true),
	}
}

func CreateRoute() services.CreateRouteRequest {
	return services.CreateRouteRequest{
		StartAddressID: "",
		StartTime:      null.NewString("2023-01-01", true),
		EndAddressID:   "",
		EndTime:        null.NewString("2023-02-02", true),
		WarrantID:      "",
	}
}

func UpdateRoute() services.PatchRouteRequest {
	return services.PatchRouteRequest{
		StartTime: null.NewString("2023-03-03", true),
		EndTime:   null.NewString("2023-04-04", true),
		Status:    null.NewString("completed", true),
	}
}

func CreateRepair() services.CreateRepairRequest {
	return services.CreateRepairRequest{
		WarrantID:  "",
		AddressID:  null.String{},
		Workshop:   "workshop",
		RepairType: "repairType",
		Start:      "2023-02-02 15:04:05",
		End:        null.NewString("2023-03-03", true),
	}
}

func UpdateRepair() services.PatchRepairRequest {
	return services.PatchRepairRequest{
		Workshop:   null.NewString("Updated workshop", true),
		RepairType: null.NewString("Updated repairType", true),
		Start:      null.NewString("2023-01-01 15:04:05", true),
		End:        null.NewString("2023-04-04", true),
	}
}

func CreateTransportCost() services.CreateTransportCostRequest {
	return services.CreateTransportCostRequest{
		WarrantID: "",
		Type:      null.NewString("type", true),
		Code:      null.NewString("code", true),
		Location:  "location",
		Amount:    null.NewFloat(25.3, true),
		FileName:  null.NewString("New name", true),
	}
}

func UpdateTransportCost() services.PatchTransportCostRequest {
	return services.PatchTransportCostRequest{
		Type:     null.NewString("Updated type", true),
		Code:     null.NewString("Updated code", true),
		Location: null.NewString("Updated location", true),
		Amount:   null.NewFloat(25.5, true),
	}
}

func CreateShippingInvoice() services.CreateShippingInvoiceRequest {
	return services.CreateShippingInvoiceRequest{
		RouteID:           "",
		Status:            "ready",
		ExternalInvoiceID: "externalInvoiceId",
		Side:              "start",
	}
}

func UpdateShippingInvoice() services.PatchShippingInvoiceRequest {
	return services.PatchShippingInvoiceRequest{
		RouteID:           null.String{},
		Status:            null.NewString("Updated status", true),
		ExternalInvoiceID: null.NewString("Updated externalInvoiceId", true),
		Side:              "end",
	}
}

func CreateDriverCategory() services.CreateDriverCategoryRequest {
	return services.CreateDriverCategoryRequest{
		UserID:   "",
		Category: "New Category",
		Issued:   null.NewString("2021-01-01", true),
		Expired:  null.NewString("2023-01-01", true),
	}
}

func UpdateDriverCategory() services.PatchDriverCategoryRequest {
	return services.PatchDriverCategoryRequest{
		Category: null.NewString("Updated Category", true),
		Issued:   null.NewString("2021-02-02", true),
		Expired:  null.NewString("2023-03-03", true),
	}
}
