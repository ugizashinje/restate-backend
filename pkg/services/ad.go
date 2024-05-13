package services

import (
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/repo"
	"warrant-api/pkg/utils"

	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v4"
)

type AdImpl struct {
	AddressRepo repo.Repo[model.Address]
}

type CreateAdRequest struct {
	City     string      `json:"city"`
	Street   string      `json:"street"`
	StreetNo null.String `json:"streetNo"`
	Name     null.String `json:"name"`
}

type PatchAdRequest struct {
	City     null.String `json:"city"`
	Street   null.String `json:"street"`
	StreetNo null.String `json:"streetNo"`
	Name     null.String `json:"name"`
}

func (svc *AdImpl) Create(g *gin.Context, createAddress CreateAddressRequest) (*model.Address, error) {
	address := &model.Address{
		City:   createAddress.City,
		Street: createAddress.Street}
	if createAddress.Name.Valid {
		address.Name = createAddress.Name
	}
	if createAddress.StreetNo.Valid {
		address.StreetNo = createAddress.StreetNo
	}

	result := svc.AddressRepo.Create(g, address)
	utils.Handle(result.Error)
	return address, nil
}

func (svc *AdImpl) Delete(g *gin.Context, addressID string) (address *model.Address, err error) {
	address = &model.Address{}
	result := svc.AddressRepo.Delete(g, address, addressID)
	utils.Handle(result.Error)

	return address, nil
}

func (svc *AdImpl) GetById(c *gin.Context, addressID string) (address *model.Address, err error) {
	address = &model.Address{}
	result := svc.AddressRepo.GetById(c, address, addressID)
	utils.Handle(result.Error)
	return address, nil
}

func (svc *AdImpl) Update(c *gin.Context, addressID string, patch any) (address *model.Address, err error) {
	address = &model.Address{}
	result := svc.AddressRepo.Patch(c, address, addressID, patch)
	utils.Handle(result.Error)

	return address, nil
}

func (svc *AdImpl) Search(c *gin.Context, query map[string]any) (addresss []model.Address, err error) {
	addresss = []model.Address{}

	result := svc.AddressRepo.Search(c, &addresss, query)
	utils.Handle(result.Error)
	return addresss, nil
}
