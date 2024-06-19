package fixtures

import (
	"fmt"
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/services"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

func RegisterRegularCompany() services.RegisterCompanyManagerRequest {
	return services.RegisterCompanyManagerRequest{
		CompanyName:  gofakeit.Company(),
		CompanyShort: "ASDF",
		CompanyPIB:   gofakeit.StreetNumber(),
		CompanyAddress: model.Address{
			Location: model.Location{
				City:     gofakeit.City(),
				Muni:     gofakeit.Address().State,
				Location: null.NewString("Gardos", true),
			},
			Street: gofakeit.Address().Street,
		},
		CompanyMn:    fmt.Sprint(gofakeit.Number(10000, 1000000000)),
		CompanyPhone: gofakeit.Phone(),
		CompanyEmail: gofakeit.Name() + "@testdomain.com",

		UserEmail:      uuid.NewString() + "@mail.com",
		UserPassword:   "asdfasdf",
		UserFirstName:  "Tester",
		UserLastName:   "Testic",
		UserMn:         null.NewString(fmt.Sprintf("%d", gofakeit.Number(10000, 100000)), true),
		UserBirthplace: null.NewString("Cacak", true),
		UserPhone:      gofakeit.Phone(),
		UserAddress: model.Address{
			Location: model.Location{
				City:     gofakeit.City(),
				Muni:     gofakeit.Address().State,
				Location: null.NewString("Gardos", true),
			},
			Street: gofakeit.Address().Street,
		},
	}
}

func CreateUserInSameCompany() services.CreateUserRequest {
	return services.CreateUserRequest{
		Email:      uuid.NewString() + "@mail.com",
		Password:   "asdfasdf",
		FirstName:  "InSameCompany",
		LastName:   "TestSameComp",
		Mn:         null.NewString(fmt.Sprintf("%d", gofakeit.Number(10000, 100000)), true),
		Birthplace: null.NewString("Kraljevo", false),
		Phone:      gofakeit.Phone(),
		AddressID:  "",
		CompanyID:  null.NewString("", false),
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
		Roles:      []string{"dispatcher"},
	}
}
