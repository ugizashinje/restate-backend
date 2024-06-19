package services

import (
	"net/http"
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/messages"
	"warrant-api/pkg/repo"
	"warrant-api/pkg/utils"

	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type AddressServiceImpl struct {
	AddressRepo  repo.Repo[model.Address]
	LocationRepo repo.Repo[model.Location]
}

type CreateAddressRequest struct {
	City       string      `json:"city"`
	Muni       string      `json:"muni"`
	Location   null.String `json:"location"`
	LocationID string      `json:"locationId"`
	Street     string      `json:"street"`
	StreetNo   null.String `json:"streetNo"`
	Name       null.String `json:"name"`
}

type PatchAddressRequest struct {
	LocationID null.String `json:"locationId"`
	Street     null.String `json:"street"`
	StreetNo   null.String `json:"streetNo"`
	Name       null.String `json:"name"`
}

func (svc *AddressServiceImpl) Create(g *gin.Context, createAddress CreateAddressRequest) (*model.Address, error) {
	address := &model.Address{
		LocationID: createAddress.LocationID,
		Street:     createAddress.Street,
		StreetNo:   createAddress.StreetNo,
	}

	result := svc.AddressRepo.Create(g, address)
	utils.Handle(result.Error)
	return address, nil
}

func (svc *AddressServiceImpl) FindOrCreate(g *gin.Context, locationAndStreet model.Address, addressID null.String) (*model.Address, error) {
	address := &model.Address{}
	var dbRes *gorm.DB
	var location *model.Location
	if addressID.Valid {
		dbRes = svc.AddressRepo.GetById(g, address, addressID.String)
		utils.Handle(dbRes.Error)
		if dbRes.RowsAffected == 1 {
			return address, nil
		}
		utils.Handle(messages.Errorf(http.StatusNotFound, "address not found"))
	}
	locations := []model.Location{}
	dbRes = svc.LocationRepo.Search(g, &locations, map[string]any{
		// "city":     locationAndStreet.Location.City,
		// "muni":     locationAndStreet.Location.Muni,
		// "location": locationAndStreet.Location.Location,
	})
	utils.Handle(dbRes.Error)
	if dbRes.RowsAffected == 1 {
		location = &locations[0]
	} else if dbRes.RowsAffected == 0 {
		// dbRes := svc.LocationRepo.Create(g, &locationAndStreet.Location)
		// utils.Handle(dbRes.Error)
		// location = &locationAndStreet.Location
	} else {
		utils.Handle(messages.Errorf(http.StatusBadRequest, "Molimo vas precizirajte opstinu ili komsiluk"))
	}
	address.LocationID = location.ID
	address.Street = locationAndStreet.Street
	address.StreetNo = locationAndStreet.StreetNo
	dbRes = svc.AddressRepo.Create(g, address)
	utils.Handle(dbRes.Error)
	return address, nil
}

func (svc *AddressServiceImpl) Delete(g *gin.Context, addressID string) (address *model.Address, err error) {
	address = &model.Address{}
	result := svc.AddressRepo.Delete(g, address, addressID)
	utils.Handle(result.Error)

	return address, nil
}

func (svc *AddressServiceImpl) GetById(c *gin.Context, addressID string) (address *model.Address, err error) {
	address = &model.Address{}
	result := svc.AddressRepo.GetById(c, address, addressID)
	utils.Handle(result.Error)
	return address, nil
}

func (svc *AddressServiceImpl) Update(c *gin.Context, addressID string, patch any) (address *model.Address, err error) {
	address = &model.Address{}
	result := svc.AddressRepo.Patch(c, address, addressID, patch)
	utils.Handle(result.Error)

	return address, nil
}

func (svc *AddressServiceImpl) Search(c *gin.Context, query map[string]any) (addresss []model.Address, err error) {
	addresss = []model.Address{}

	result := svc.AddressRepo.Search(c, &addresss, query)
	utils.Handle(result.Error)
	return addresss, nil
}
