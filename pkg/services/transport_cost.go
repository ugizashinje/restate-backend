package services

import (
	"warrant-api/pkg/ctx"
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/repo"
	"warrant-api/pkg/storage"
	"warrant-api/pkg/utils"

	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v4"
)

type TransportCostServiceImpl struct {
	TransportCostRepo repo.Repo[model.TransportCost]
	WarrantService    WarrantServiceImpl
}

type CreateTransportCostRequest struct {
	WarrantID string      `json:"warrantID"`
	Type      null.String `json:"type"`
	Code      null.String `json:"code"`
	Location  string      `json:"location"`
	Amount    null.Float  `json:"amount"`
	FileName  null.String `json:"fileName"`
}

type PatchTransportCostRequest struct {
	Type     null.String `json:"type"`
	Code     null.String `json:"code"`
	Location null.String `json:"location"`
	Amount   null.Float  `json:"amount"`
	FileName null.String `json:"fileName"`
}

func (svc *TransportCostServiceImpl) Create(g *gin.Context, createTransportCost CreateTransportCostRequest) (*model.TransportCost, error) {
	svc.WarrantService.CheckWarrantGrants(g, createTransportCost.WarrantID, true)

	transportcost := &model.TransportCost{
		WarrantID: createTransportCost.WarrantID,
		Type:      createTransportCost.Type,
		Code:      createTransportCost.Code,
		Location:  createTransportCost.Location,
		FileName:  createTransportCost.FileName,
	}
	result := svc.TransportCostRepo.Create(g, transportcost)
	utils.Handle(result.Error)

	err := storage.PutTransportCost(g, transportcost)
	utils.Handle(err)

	return transportcost, nil
}

func (svc *TransportCostServiceImpl) Delete(g *gin.Context, transportcostID string) (transportcost *model.TransportCost, err error) {
	transportcost = &model.TransportCost{}
	dbRes := svc.TransportCostRepo.GetById(g, transportcost, transportcostID)
	utils.Handle(dbRes.Error)

	svc.WarrantService.CheckWarrantGrants(g, transportcost.WarrantID, true)

	transportcost = &model.TransportCost{}
	result := svc.TransportCostRepo.Delete(g, transportcost, transportcostID)
	utils.Handle(result.Error)
	return transportcost, nil
}

func (svc *TransportCostServiceImpl) GetById(g *gin.Context, transportcostID string) (transportcost *model.TransportCost, err error) {
	transportcost = &model.TransportCost{}
	result := svc.TransportCostRepo.GetById(g, transportcost, transportcostID)
	utils.Handle(result.Error)
	svc.WarrantService.CheckWarrantGrants(g, transportcost.WarrantID, false)

	err = storage.GetTransportCost(g, transportcost)
	utils.Handle(err)

	return transportcost, nil
}

func (svc *TransportCostServiceImpl) Update(g *gin.Context, transportcostID string, patch PatchTransportCostRequest) (transportcost *model.TransportCost, err error) {
	transportcost = &model.TransportCost{}

	dbRes := svc.TransportCostRepo.GetById(g, transportcost, transportcostID)
	utils.Handle(dbRes.Error)
	oldName := ""
	if transportcost.FileName.Valid && patch.FileName.Valid {
		oldName = transportcost.FileName.String
	}

	svc.WarrantService.CheckWarrantGrants(g, transportcost.WarrantID, false)

	transportcost = &model.TransportCost{}
	result := svc.TransportCostRepo.Patch(g, transportcost, transportcostID, patch)
	utils.Handle(result.Error)

	if transportcost.FileName.Valid {
		if oldName == "" {
			err = storage.PutTransportCost(g, transportcost)
		} else {
			err = storage.UpdateTransportConst(g, transportcost, oldName)
		}
	} else {
		err = storage.GetTransportCost(g, transportcost)
	}
	utils.Handle(err)

	return transportcost, nil
}

func (svc *TransportCostServiceImpl) Search(g *gin.Context, query map[string]any) (transportcosts []model.TransportCost, err error) {
	warrantID := ctx.GetMandatoryQueryParam(g, "warrantId")
	svc.WarrantService.CheckWarrantGrants(g, warrantID, false)
	query["warrant_id"] = warrantID
	transportcosts = []model.TransportCost{}

	result := svc.TransportCostRepo.Search(g, &transportcosts, query)
	utils.Handle(result.Error)
	return transportcosts, nil
}
