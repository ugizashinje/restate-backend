package services

import (
	"net/http"
	"time"
	"warrant-api/pkg/config"
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/messages"
	"warrant-api/pkg/repo"
	"warrant-api/pkg/session"
	"warrant-api/pkg/utils"

	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v4"
	"gorm.io/datatypes"
)

type VehicleServiceImpl struct {
	WarrantRepo repo.Repo[model.Warrant]
	VehicleRepo repo.Repo[model.Vehicle]
}

type CreateVehicleRequest struct {
	Plate                       string         `json:"plate"`
	CompanyID                   string         `json:"companyId"`
	Model                       string         `json:"model"`
	Producer                    string         `json:"producer"`
	Type                        string         `json:"type"`
	Weight                      null.Int       `json:"weight"`
	MaxWeight                   null.Int       `json:"maxWeight"`
	Capacity                    null.Int       `json:"capacity"`
	HomologationMark            string         `json:"homologationMark"`
	MotorVolume                 int            `json:"motorVolume"`
	Power                       int            `json:"power"`
	Fuel                        string         `json:"fuel"`
	Consumption                 float32        `json:"consumption"`
	ThrustWeight                null.Int       `json:"thrustWeight"`
	StandingSpace               null.Int       `json:"standingSpace"`
	SeatingSpace                null.Int       `json:"seatingSpace"`
	Color                       string         `json:"color"`
	Shaft                       int            `json:"shaft"`
	Chasie                      string         `json:"chasie"`
	Motor                       string         `json:"motor"`
	Insurance                   null.String    `json:"insurance"`
	RegistrationDate            string         `json:"registrationDate" binding:"required,date"`
	FirstRegistrationDate       null.String    `json:"firstRegistrationDate" binding:"omitempty,date"`
	OwnerMn                     int            `json:"ownerMn"`
	OwnerFirstName              string         `json:"ownerFirstName"`
	OwnerLastName               string         `json:"ownerLastName"`
	OwnerAddressID              string         `json:"ownerAddressId"`
	UserFirstName               string         `json:"userFirstName"`
	UserLastName                string         `json:"userLastName"`
	UserAddressID               string         `json:"userAddressId"`
	UserMn                      int            `json:"userMn"`
	RegistrationCertificateDate string         `json:"registrationCertificateDate" binding:"required,date"`
	RegistrationExpiry          string         `json:"registrationExpiry" binding:"required,date"`
	RegisterId                  int            `json:"registerId"`
	Status                      string         `json:"status"`
	Meta                        datatypes.JSON `json:"meta"`
}

type PatchVehicleRequest struct {
	ComapnyID                   null.String    `json:"companyId"`
	Plate                       null.String    `json:"plate"`
	Model                       null.String    `json:"model"`
	Producer                    null.String    `json:"producer"`
	Type                        null.String    `json:"type"`
	Weight                      null.Int       `json:"weight"`
	MaxWeight                   null.Int       `json:"maxWeight"`
	Capacity                    null.Int       `json:"capacity"`
	HomologationMark            null.String    `json:"homologationMark"`
	MotorVolume                 null.Int       `json:"motorVolume"`
	Power                       null.Int       `json:"power"`
	Fuel                        null.String    `json:"fuel"`
	Consumption                 null.Float     `json:"consumption"`
	ThrustWeight                null.Int       `json:"thrustWeight"`
	StandingSpace               null.Int       `json:"standingSpace"`
	SeatingSpace                null.Int       `json:"seatingSpace"`
	Color                       null.String    `json:"color"`
	Shaft                       null.Int       `json:"shaft"`
	Chasie                      null.String    `json:"chasie"`
	Motor                       null.String    `json:"motor"`
	Insurance                   null.String    `json:"insurance"`
	RegistrationDate            null.String    `json:"registrationDate" binding:"omitempty,date"`
	FirstRegistrationDate       null.String    `json:"firstRegistrationDate" binding:"omitempty,date"`
	OwnerMn                     null.Int       `json:"ownerMn"`
	OwnerLastName               null.String    `json:"ownerLastName"`
	OwnerFirstName              null.String    `json:"ownerFirstName"`
	OwnerAddressID              null.String    `json:"ownerAddressId"`
	UserFirstName               null.String    `json:"userFirstName"`
	UserLastName                null.String    `json:"userLastName"`
	UserAddressID               null.String    `json:"userAddressId"`
	UserMn                      null.Int       `json:"userMn"`
	RegistrationCertificateDate null.String    `json:"registrationCertificateDate" binding:"omitempty,date"`
	RegistrationExpiry          null.String    `json:"registrationExpiry" binding:"omitempty,date"`
	RegisterId                  null.Int       `json:"registerId"`
	Status                      null.String    `json:"status"`
	Meta                        datatypes.JSON `json:"meta"`
}

func (svc *VehicleServiceImpl) Create(g *gin.Context, createVehicle CreateVehicleRequest) (*model.Vehicle, error) {
	registrationDate, err := time.Parse(config.Format.DateFormat, createVehicle.RegistrationDate)
	utils.Handle(err)
	firstRegistrationDate := utils.ParseNullDate(createVehicle.FirstRegistrationDate.String, createVehicle.FirstRegistrationDate.Valid)
	registrationCertificateDate, err := time.Parse(config.Format.DateFormat, createVehicle.RegistrationCertificateDate)
	utils.Handle(err)
	registrationExpiry, err := time.Parse(config.Format.DateFormat, createVehicle.RegistrationExpiry)
	utils.Handle(err)

	userSession := session.GetSession(g)
	if userSession == nil {
		utils.Handle(messages.Errorf(http.StatusUnauthorized, "You are not allowed to perform this operation"))
	}
	if isManager, isDispatcher, _ := userSession.Roles(createVehicle.CompanyID); !(isManager || isDispatcher) {
		utils.Handle(messages.Errorf(http.StatusUnauthorized, "You are not allowed to perform this operation"))
	}
	vehicle := &model.Vehicle{
		CompanyID:                   createVehicle.CompanyID,
		Plate:                       createVehicle.Plate,
		Model:                       createVehicle.Model,
		Producer:                    createVehicle.Producer,
		Type:                        createVehicle.Type,
		Weight:                      createVehicle.Weight,
		MaxWeight:                   createVehicle.MaxWeight,
		Capacity:                    createVehicle.Capacity,
		HomologationMark:            createVehicle.HomologationMark,
		MotorVolume:                 createVehicle.MotorVolume,
		Power:                       createVehicle.Power,
		Fuel:                        createVehicle.Fuel,
		Consumption:                 createVehicle.Consumption,
		ThrustWeight:                createVehicle.ThrustWeight,
		StandingSpace:               createVehicle.StandingSpace,
		SeatingSpace:                createVehicle.SeatingSpace,
		Color:                       createVehicle.Color,
		Shaft:                       createVehicle.Shaft,
		Chasie:                      createVehicle.Chasie,
		Motor:                       createVehicle.Motor,
		Insurance:                   createVehicle.Insurance,
		RegistrationDate:            registrationDate,
		FirstRegistrationDate:       firstRegistrationDate,
		OwnerMn:                     createVehicle.OwnerMn,
		OwnerFirstName:              createVehicle.OwnerFirstName,
		OwnerLastName:               createVehicle.OwnerLastName,
		OwnerAddressID:              createVehicle.OwnerAddressID,
		UserFirstName:               createVehicle.UserFirstName,
		UserLastName:                createVehicle.UserLastName,
		UserAddressID:               createVehicle.UserAddressID,
		UserMn:                      createVehicle.UserMn,
		RegistrationCertificateDate: registrationCertificateDate,
		RegistrationExpiry:          registrationExpiry,
		RegisterId:                  createVehicle.RegisterId,
		Meta:                        createVehicle.Meta,
	}

	result := svc.VehicleRepo.Create(g, vehicle)
	utils.Handle(result.Error)
	return vehicle, nil
}

func (svc *VehicleServiceImpl) Delete(g *gin.Context, vehicleID string) (vehicle *model.Vehicle, err error) {
	vehicle = &model.Vehicle{}
	dbRes := svc.VehicleRepo.GetById(g, vehicle, vehicleID)
	utils.Handle(dbRes.Error)
	userSession := session.GetSession(g)
	if isManager, isDispatcher, _ := userSession.Roles(vehicle.CompanyID); !(isManager || isDispatcher) {
		utils.Handle(messages.Errorf(http.StatusUnauthorized, "You are not allowed to perform this operation"))
	}
	vehicle = &model.Vehicle{}
	result := svc.VehicleRepo.Delete(g, vehicle, vehicleID)
	utils.Handle(result.Error)

	return vehicle, nil
}

func (svc *VehicleServiceImpl) GetById(g *gin.Context, vehicleID string) (vehicle *model.Vehicle, err error) {
	vehicle = &model.Vehicle{}
	dbRes := svc.VehicleRepo.GetById(g, vehicle, vehicleID)
	utils.Handle(dbRes.Error)
	userSession := session.GetSession(g)
	isManager, isDispatcher, isDriver := userSession.Roles(vehicle.CompanyID)
	if !(isManager || isDispatcher || isDriver) {
		utils.Handle(messages.Errorf(http.StatusUnauthorized, "You are not allowed to perform this operation"))
	} else if isDriver {
		vehicleWarrants := []model.Warrant{}
		svc.WarrantRepo.Search(g, &vehicleWarrants, map[string]any{"company_id": vehicle.CompanyID, "driver_id": userSession.User.ID}, repo.JustFirst)
		if len(vehicleWarrants) == 0 {
			utils.Handle(messages.Errorf(http.StatusUnauthorized, "You are not allowed to perform this operation"))
		}
	}

	vehicle = &model.Vehicle{}
	result := svc.VehicleRepo.GetById(g, vehicle, vehicleID, repo.Preload("OwnerAddress"), repo.Preload("UserAddress"))
	utils.Handle(result.Error)
	return vehicle, nil
}

func (svc *VehicleServiceImpl) Update(g *gin.Context, vehicleID string, patch any) (vehicle *model.Vehicle, err error) {
	vehicle = &model.Vehicle{}
	dbRes := svc.VehicleRepo.GetById(g, vehicle, vehicleID)
	utils.Handle(dbRes.Error)
	userSession := session.GetSession(g)
	if isManager, isDispatcher, _ := userSession.Roles(vehicle.CompanyID); !(isManager || isDispatcher) {
		utils.Handle(messages.Errorf(http.StatusUnauthorized, "You are not allowed to perform this operation"))
	}
	vehicle = &model.Vehicle{}

	// TODO can driver change status or meta on vehicle
	result := svc.VehicleRepo.Patch(g, vehicle, vehicleID, patch)
	utils.Handle(result.Error)

	return vehicle, nil
}

func (svc *VehicleServiceImpl) Search(c *gin.Context, query map[string]any) (vehicles []model.Vehicle, err error) {
	vehicles = []model.Vehicle{}
	result := svc.VehicleRepo.Search(c, &vehicles, query, repo.GrantVehicle, repo.Preload("OwnerAddress"), repo.Preload("UserAddress"))
	utils.Handle(result.Error)
	return vehicles, nil
}
