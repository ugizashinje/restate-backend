package services

import (
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/repo"
	"warrant-api/pkg/utils"

	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v4"
)

type DriverCategoryServiceImpl struct {
	DriverCategoryRepo repo.Repo[model.DriverCategory]
}

type CreateDriverCategoryRequest struct {
	UserID   string      `json:"userId"`
	Category string      `json:"category"`
	Issued   null.String `json:"issued" binding:"omitempty,date"`
	Expired  null.String `json:"expired" binding:"omitempty,date"`
}

type PatchDriverCategoryRequest struct {
	Category null.String `json:"category"`
	Issued   null.String `json:"issued" binding:"omitempty,date"`
	Expired  null.String `json:"expired" binding:"omitempty,date"`
}

func (svc *DriverCategoryServiceImpl) Create(g *gin.Context, createDriverCategory CreateDriverCategoryRequest) (*model.DriverCategory, error) {
	issued := utils.ParseNullDate(createDriverCategory.Issued.String, createDriverCategory.Issued.Valid)
	expired := utils.ParseNullDate(createDriverCategory.Expired.String, createDriverCategory.Expired.Valid)

	drivercategory := &model.DriverCategory{
		UserId:   createDriverCategory.UserID,
		Category: createDriverCategory.Category,
		Issued:   issued,
		Expired:  expired,
	}

	result := svc.DriverCategoryRepo.Create(g, drivercategory)
	utils.Handle(result.Error)
	return drivercategory, nil
}

func (svc *DriverCategoryServiceImpl) Delete(g *gin.Context, drivercategoryID string) (drivercategory *model.DriverCategory, err error) {
	drivercategory = &model.DriverCategory{}
	result := svc.DriverCategoryRepo.Delete(g, drivercategory, drivercategoryID)
	utils.Handle(result.Error)

	return drivercategory, nil
}

func (svc *DriverCategoryServiceImpl) GetById(c *gin.Context, drivercategoryID string) (drivercategory *model.DriverCategory, err error) {
	drivercategory = &model.DriverCategory{}
	result := svc.DriverCategoryRepo.GetById(c, drivercategory, drivercategoryID)
	utils.Handle(result.Error)
	return drivercategory, nil
}

func (svc *DriverCategoryServiceImpl) Update(c *gin.Context, drivercategoryID string, patch any) (drivercategory *model.DriverCategory, err error) {
	drivercategory = &model.DriverCategory{}
	result := svc.DriverCategoryRepo.Patch(c, drivercategory, drivercategoryID, patch)
	utils.Handle(result.Error)

	return drivercategory, nil
}
