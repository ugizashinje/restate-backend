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
	RouteService              services.RouteServiceImpl
	CompanyService            services.CompanyServiceImpl
	UserService               services.UserServiceImpl
	VehicleService            services.VehicleServiceImpl
	WarrantEventService       services.WarrantEventServiceImpl
	WarrantService            services.WarrantServiceImpl
	DriverCategoryService     services.DriverCategoryServiceImpl
	RepairService             services.RepairServiceImpl
	ShippingInvoiceService    services.ShippingInvoiceServiceImpl
	TransportCostService      services.TransportCostServiceImpl
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
		RouteService: services.RouteServiceImpl{
			RouteRepo:      repo.Repo[model.Route]{},
			WarrantService: services.WarrantServiceImpl{},
			AddressService: addressService,
		},
		CompanyService: services.CompanyServiceImpl{
			AddressService: addressService,
			CompanyRepo:    repo.Repo[model.Company]{},
		},
		VehicleService: services.VehicleServiceImpl{
			WarrantRepo: repo.Repo[model.Warrant]{},
			VehicleRepo: repo.Repo[model.Vehicle]{},
		},
		WarrantEventService: services.WarrantEventServiceImpl{
			WarrantRepo:      repo.Repo[model.Warrant]{},
			WarrantEventRepo: repo.Repo[model.WarrantEvent]{},
		},
		WarrantService: services.WarrantServiceImpl{
			ChangeRepo:  repo.Repo[model.Change]{},
			WarrantRepo: repo.Repo[model.Warrant]{},
		},
		UserService: services.UserServiceImpl{
			UserRepo:        repo.User{},
			UserCompanyRepo: repo.Repo[model.UserCompany]{},
		},
		DriverCategoryService: services.DriverCategoryServiceImpl{
			DriverCategoryRepo: repo.Repo[model.DriverCategory]{},
		},
		RepairService: services.RepairServiceImpl{
			WarrantService: services.WarrantServiceImpl{},
			RepairRepo:     repo.Repo[model.Repair]{},
		},
		TransportCostService: services.TransportCostServiceImpl{
			WarrantService:    services.WarrantServiceImpl{},
			TransportCostRepo: repo.Repo[model.TransportCost]{},
		},
		ShippingInvoiceService: services.ShippingInvoiceServiceImpl{
			RouteRepo:           repo.Repo[model.Route]{},
			WarrantRepo:         repo.Repo[model.Warrant]{},
			ShippingInvoiceRepo: repo.Repo[model.ShippingInvoice]{},
		},
		TransactionalEmailService: services.TransactionalEmailServiceImpl{},
	}
	return &w
}

var Svc *Wires
