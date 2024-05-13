package wire

import (
	"warrant-api/pkg/config"
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/repo"
	"warrant-api/pkg/services"
	"warrant-api/pkg/session"
	"warrant-api/pkg/storage"

	"github.com/go-resty/resty/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Wires struct {
	AuthService               services.AuthServiceImpl
	AddressService            services.AddressServiceImpl
	CompanyService            services.CompanyServiceImpl
	UserService               services.UserServiceImpl
	WarrantService            services.WarrantServiceImpl
	TransactionalEmailService services.TransactionalEmailServiceImpl
}

func Init(env string) *Wires {
	config.Init(env)
	session.Init()
	storage.Init()
	addressService := services.AddressServiceImpl{
		AddressRepo: repo.Repo[model.Address]{},
	}
	w := Wires{
		AuthService: services.AuthServiceImpl{
			Method:           &jwt.SigningMethodEd25519{},
			LoginRepo:        repo.Repo[model.Login]{},
			ConfirmationRepo: repo.Repo[model.Confirmation]{},
			AddressService:   services.AddressServiceImpl{},
			UserService:      services.UserServiceImpl{},
			CompanyService:   services.CompanyServiceImpl{},
			TransactionalEmailService: services.TransactionalEmailServiceImpl{
				ConfirmationRepo: repo.Repo[model.Confirmation]{},
			},
			RestyClient:           resty.New(),
			SupersetGuestTokenUrl: config.Superset.Url + "/security/guest_token",
		},
		AddressService: addressService,
		CompanyService: services.CompanyServiceImpl{
			AddressService: addressService,
			CompanyRepo:    repo.Repo[model.Company]{},
		},
		UserService: services.UserServiceImpl{
			UserRepo:        repo.User{},
			UserCompanyRepo: repo.Repo[model.UserCompany]{},
		},
		TransactionalEmailService: services.TransactionalEmailServiceImpl{},
	}
	return &w
}

var Svc *Wires
