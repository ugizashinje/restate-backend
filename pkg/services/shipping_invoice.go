package services

import (
	"warrant-api/pkg/ctx"
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/messages"
	"warrant-api/pkg/repo"
	"warrant-api/pkg/session"
	"warrant-api/pkg/utils"

	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type ShippingInvoiceServiceImpl struct {
	RouteRepo           repo.Repo[model.Route]
	WarrantRepo         repo.Repo[model.Warrant]
	ShippingInvoiceRepo repo.Repo[model.ShippingInvoice]
}

type CreateShippingInvoiceRequest struct {
	RouteID           string `json:"routeId" binding:"required"`
	Status            string `json:"status" binding:"required"`
	ExternalInvoiceID string `json:"externalInvoiceId"`
	Side              string `json:"side" binding:"required,oneof=start end"`
}

type PatchShippingInvoiceRequest struct {
	RouteID           null.String `json:"routeId"`
	Side              string      `json:"side" binding:"omitempty,oneof=start end"`
	Status            null.String `json:"status"`
	ExternalInvoiceID null.String `json:"externalInvoiceId"`
}

func (svc *ShippingInvoiceServiceImpl) Create(g *gin.Context, createShippingInvoice CreateShippingInvoiceRequest) (*model.ShippingInvoice, error) {
	shippinginvoice := &model.ShippingInvoice{
		RouteID:           createShippingInvoice.RouteID,
		Status:            createShippingInvoice.Status,
		ExternalInvoiceID: createShippingInvoice.ExternalInvoiceID,
		Side:              createShippingInvoice.Side,
	}
	route := &model.Route{}
	result := svc.RouteRepo.GetById(g, route, shippinginvoice.RouteID)
	utils.Handle(result.Error)

	warrant := &model.Warrant{}
	result = svc.WarrantRepo.GetById(g, warrant, route.WarrantID)
	utils.Handle(result.Error)

	userSession := session.GetSession(g)
	if userSession == nil {
		utils.Handle(messages.Unauthorized())

	}
	if isManager, isDispatcher, isDriver := userSession.Roles(warrant.CompanyID); !isManager && !isDispatcher && !(isDriver && warrant.DriverID == userSession.User.ID) {
		utils.Handle(messages.Unauthorized())
	}
	result = svc.ShippingInvoiceRepo.Create(g, shippinginvoice)
	utils.Handle(result.Error)
	return shippinginvoice, nil
}

func (svc *ShippingInvoiceServiceImpl) Delete(g *gin.Context, shippinginvoiceID string) (shippinginvoice *model.ShippingInvoice, err error) {
	svc.checkAccess(g, shippinginvoiceID, "")

	result := svc.ShippingInvoiceRepo.Delete(g, shippinginvoice, shippinginvoiceID)
	utils.Handle(result.Error)

	return shippinginvoice, nil
}

func (svc *ShippingInvoiceServiceImpl) GetById(g *gin.Context, shippinginvoiceID string) (shippinginvoice *model.ShippingInvoice, err error) {
	svc.checkAccess(g, shippinginvoiceID, "")
	shippinginvoice = &model.ShippingInvoice{}

	result := svc.ShippingInvoiceRepo.GetById(g, shippinginvoice, shippinginvoiceID)
	utils.Handle(result.Error)
	return shippinginvoice, nil
}

func (svc *ShippingInvoiceServiceImpl) Update(g *gin.Context, shippinginvoiceID string, patch any) (shippinginvoice *model.ShippingInvoice, err error) {
	svc.checkAccess(g, shippinginvoiceID, "")
	shippinginvoice = &model.ShippingInvoice{}
	result := svc.ShippingInvoiceRepo.Patch(g, shippinginvoice, shippinginvoiceID, patch)
	utils.Handle(result.Error)

	return shippinginvoice, nil
}

func (svc *ShippingInvoiceServiceImpl) Search(g *gin.Context, query map[string]any) (shippinginvoices []model.ShippingInvoice, err error) {

	warrantID := ctx.GetMandatoryQueryParam(g, "warrantId")
	warrant := &model.Warrant{}
	dbRes := svc.WarrantRepo.GetById(g, warrant, warrantID)
	utils.Handle(dbRes.Error)
	routes := []model.Route{}
	dbRes = svc.RouteRepo.Search(g, &routes, nil)
	utils.Handle(dbRes.Error)
	routeIds := []string{}
	for _, route := range routes {
		routeIds = append(routeIds, route.ID)
	}
	shippinginvoices = []model.ShippingInvoice{}

	result := svc.ShippingInvoiceRepo.Search(g, &shippinginvoices, query, func(g *gin.Context, db *gorm.DB) *gorm.DB {
		return db.Where("route_id in (?)", routeIds)
	})
	utils.Handle(result.Error)
	return shippinginvoices, nil
}

func (svc *ShippingInvoiceServiceImpl) checkAccess(g *gin.Context, shippinginvoiceID string, routeID string) {
	shippinginvoice := &model.ShippingInvoice{}
	if routeID == "" {
		result := svc.ShippingInvoiceRepo.GetById(g, shippinginvoice, shippinginvoiceID)
		utils.Handle(result.Error)
	}

	route := &model.Route{}
	result := svc.RouteRepo.GetById(g, route, shippinginvoice.RouteID)
	utils.Handle(result.Error)

	warrant := &model.Warrant{}
	result = svc.WarrantRepo.GetById(g, warrant, route.WarrantID)
	utils.Handle(result.Error)

	userSession := session.GetSession(g)
	if userSession == nil {
		utils.Handle(messages.Unauthorized())

	}
	if isManager, isDispatcher, isDriver := userSession.Roles(warrant.CompanyID); !isManager && !isDispatcher && !(isDriver && warrant.DriverID == userSession.User.ID) {
		utils.Handle(messages.Unauthorized())
	}

}
