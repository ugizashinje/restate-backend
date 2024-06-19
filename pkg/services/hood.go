package services

import (
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/repo"
	"warrant-api/pkg/utils"

	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v4"
)

type LocationServiceImpl struct {
	LocationRepo repo.Repo[model.Location]
}

type CreateLocationRequest struct {
	City     string      `json:"city"`
	Muni     string      `json:"muni"`
	Location null.String `json:"location"`
}

func (svc *LocationServiceImpl) Create(g *gin.Context, createLocation CreateLocationRequest) (*model.Location, error) {
	location := &model.Location{
		City:     createLocation.City,
		Muni:     createLocation.Muni,
		Location: createLocation.Location,
	}

	result := svc.LocationRepo.Create(g, location)
	utils.Handle(result.Error)
	return location, nil
}

func (svc *LocationServiceImpl) Delete(g *gin.Context, locationID string) (location *model.Location, err error) {
	location = &model.Location{}
	result := svc.LocationRepo.Delete(g, location, locationID)
	utils.Handle(result.Error)

	return location, nil
}

func (svc *LocationServiceImpl) GetById(c *gin.Context, locationID string) (location *model.Location, err error) {
	location = &model.Location{}
	result := svc.LocationRepo.GetById(c, location, locationID)
	utils.Handle(result.Error)
	return location, nil
}

func (svc *LocationServiceImpl) Update(c *gin.Context, locationID string, patch any) (location *model.Location, err error) {
	location = &model.Location{}
	result := svc.LocationRepo.Patch(c, location, locationID, patch)
	utils.Handle(result.Error)

	return location, nil
}

func (svc *LocationServiceImpl) Search(c *gin.Context, query map[string]any) (locations []model.Location, err error) {
	locations = []model.Location{}

	result := svc.LocationRepo.Search(c, &locations, query)
	utils.Handle(result.Error)
	return locations, nil
}
