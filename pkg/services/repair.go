package services

import (
	"time"
	"warrant-api/pkg/config"
	"warrant-api/pkg/ctx"
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/repo"
	"warrant-api/pkg/storage"
	"warrant-api/pkg/utils"

	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v4"
)

type RepairServiceImpl struct {
	WarrantService WarrantServiceImpl
	RepairRepo     repo.Repo[model.Repair]
}

/*
repari se kreira lokacijom kordinatama
gde, kad i slika - ostatak se popunjava kroz web ( dispatcher verovatno)
*/

type CreateRepairRequest struct {
	WarrantID  string      `json:"warrantId"`
	AddressID  null.String `json:"addressId"`
	Location   null.String `json:"location"`
	Workshop   string      `json:"workshop"`
	RepairType string      `json:"repairType"`
	Start      string      `json:"start" binding:"required,time"`
	End        null.String `json:"end" binding:"omitempty,time"`
	FileName   null.String `json:"fileName"`
}

/*
patch postoji ali mobilni dirver ograniceno ce menjati repair
*/
type PatchRepairRequest struct {
	AddressID  null.String `json:"addressId"`
	Workshop   null.String `json:"workshop"`
	RepairType null.String `json:"repairType"`
	Start      null.String `json:"start" binding:"omitempty,time"`
	End        null.String `json:"end" binding:"omitempty,time"`
	Location   null.String `json:"location"`
	FileName   null.String `json:"fileName"`
}

func (svc *RepairServiceImpl) Create(g *gin.Context, createRepair CreateRepairRequest) (*model.Repair, error) {

	svc.WarrantService.CheckWarrantGrants(g, createRepair.WarrantID, true)

	start, err := time.Parse(config.Format.TimeFormat, createRepair.Start)
	utils.Handle(err)
	end := utils.ParseNullTime(createRepair.End.String, createRepair.End.Valid)

	repair := &model.Repair{
		WarrantID:  createRepair.WarrantID,
		AddressID:  createRepair.AddressID,
		Workshop:   createRepair.Workshop,
		RepairType: createRepair.RepairType,
		FileName:   createRepair.FileName,
		Start:      start,
		End:        end,
	}
	result := svc.RepairRepo.Create(g, repair)
	utils.Handle(result.Error)

	if createRepair.FileName.Valid {
		storage.PutRepair(g, repair)
	}
	return repair, nil
}

func (svc *RepairServiceImpl) Delete(g *gin.Context, repairID string) (repair *model.Repair, err error) {
	repair = &model.Repair{}
	result := svc.RepairRepo.GetById(g, repair, repairID)
	utils.Handle(result.Error)

	svc.WarrantService.CheckWarrantGrants(g, repair.WarrantID, false)

	result = svc.RepairRepo.Delete(g, repair, repairID)
	utils.Handle(result.Error)

	return repair, nil
}

func (svc *RepairServiceImpl) GetById(g *gin.Context, repairID string) (repair *model.Repair, err error) {
	repair = &model.Repair{}
	result := svc.RepairRepo.GetById(g, repair, repairID)
	utils.Handle(result.Error)

	svc.WarrantService.CheckWarrantGrants(g, repair.WarrantID, false)

	err = storage.GetRepair(g, repair)
	utils.Handle(err)

	return repair, nil
}

func (svc *RepairServiceImpl) Update(g *gin.Context, repairID string, patch PatchRepairRequest) (repair *model.Repair, err error) {

	repair = &model.Repair{}
	dbRes := svc.RepairRepo.GetById(g, repair, repairID)
	utils.Handle(dbRes.Error)
	oldName := ""
	if repair.FileName.Valid && patch.FileName.Valid {
		oldName = repair.FileName.String
	}

	svc.WarrantService.CheckWarrantGrants(g, repair.WarrantID, false)

	repair = &model.Repair{}
	result := svc.RepairRepo.Patch(g, repair, repairID, patch)
	utils.Handle(result.Error)

	if repair.FileName.Valid {
		if oldName == "" {
			err = storage.PutRepair(g, repair)
		} else {
			err = storage.UpdateRepair(g, repair, oldName)
		}
	} else {
		err = storage.GetRepair(g, repair)
	}
	utils.Handle(err)

	return repair, nil
}

func (svc *RepairServiceImpl) Search(c *gin.Context, query map[string]any) (repairs []model.Repair, err error) {
	repairs = []model.Repair{}
	ctx.GetMandatoryQueryParam(c, "warrantId")
	result := svc.RepairRepo.Search(c, &repairs, query, repo.GrantRepair, repo.Preload("Address"))
	utils.Handle(result.Error)
	return repairs, nil
}
