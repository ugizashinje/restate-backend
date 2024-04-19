package services

import (
	"warrant-api/pkg/ctx"
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/enum"
	"warrant-api/pkg/repo"
	"warrant-api/pkg/session"
	"warrant-api/pkg/utils"

	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v4"
)

type RouteServiceImpl struct {
	RouteRepo      repo.Repo[model.Route]
	WarrantService WarrantServiceImpl
	AddressService AddressServiceImpl
}

type PatchRouteRequest struct {
	StartMileage       null.Int    `json:"startMileage"`
	StartAddressID     null.String `json:"startAddressId"`
	StartTime          null.String `json:"startTime" binding:"omitempty,time"`
	EndMileage         null.Int    `json:"endMileage"`
	EndAddressID       null.String `json:"endAddressId"`
	EndTime            null.String `json:"endTime" binding:"omitempty,time"`
	Order              null.Int    `json:"order"`
	LoadingInvoiceID   null.String `json:"loadingInvoiceId"`
	UnloadingInvoiceID null.String `json:"unloadingInvoiceId"`
	Status             null.String `json:"status"`
}

type CreateRouteRequest struct {
	WarrantID          string      `json:"warrantId"`
	StartAddressID     string      `json:"startAddressId"`
	StartTime          null.String `json:"startTime" binding:"omitempty,time"`
	EndAddressID       string      `json:"endAddressId"`
	EndTime            null.String `json:"endTime" binding:"omitempty,time"`
	Order              int         `json:"order"`
	Status             string      `json:"status" binding:"required"`
	LoadingInvoiceID   null.String `json:"loadingInvoiceId"`
	UnloadingInvoiceID null.String `json:"unloadingInvoiceId"`
}

func (svc *RouteServiceImpl) Create(g *gin.Context, createRoute CreateRouteRequest) (*model.Route, error) {
	raw, _ := g.Get(ctx.Session)
	userSession, _ := raw.(session.Session)

	warrant := svc.WarrantService.CheckWarrantGrants(g, createRoute.WarrantID, true)

	startTime := utils.ParseNullTime(createRoute.StartTime.String, createRoute.StartTime.Valid)
	endTime := utils.ParseNullTime(createRoute.EndTime.String, createRoute.EndTime.Valid)

	route := &model.Route{
		StartAddressID: createRoute.StartAddressID,
		EndAddressID:   createRoute.EndAddressID,
		WarrantID:      createRoute.WarrantID,
		Order:          createRoute.Order,
		Status:         createRoute.Status,
		StartTime:      startTime,
		EndTime:        endTime,
	}

	result := svc.RouteRepo.Create(g, route, repo.Preload("StartAddress"), repo.Preload("EndAddress"))
	utils.Handle(result.Error)
	userSession.LogEvent(g, warrant, route, createRoute, enum.CreateRoute)

	return route, nil
}

func (svc *RouteServiceImpl) Delete(g *gin.Context, routeID string) (route *model.Route, err error) {

	route, err = svc.GetById(g, routeID)
	utils.Handle(err)
	svc.WarrantService.CheckWarrantGrants(g, route.WarrantID, true)

	result := svc.RouteRepo.Delete(g, route, routeID)
	utils.Handle(result.Error)

	return route, nil
}

func (svc *RouteServiceImpl) GetById(g *gin.Context, routeID string) (route *model.Route, err error) {
	route = &model.Route{}
	dbRes := svc.RouteRepo.GetById(g, route, routeID)
	utils.Handle(dbRes.Error)
	svc.WarrantService.CheckWarrantGrants(g, route.WarrantID, false)

	route = &model.Route{}
	result := svc.RouteRepo.GetById(g, route, routeID, repo.Preload("StartAddress"), repo.Preload("EndAddress"))
	utils.Handle(result.Error)
	return route, nil
}

func (svc *RouteServiceImpl) Update(g *gin.Context, routeID string, patch PatchRouteRequest) (route *model.Route, err error) {
	raw, _ := g.Get(ctx.Session)
	userSession, _ := raw.(session.Session)
	route, err = svc.GetById(g, routeID)
	utils.Handle(err)
	warrant := svc.WarrantService.CheckWarrantGrants(g, route.WarrantID, true)

	route = &model.Route{}
	result := svc.RouteRepo.Patch(g, route, routeID, patch)
	utils.Handle(result.Error)
	userSession.LogEvent(g, warrant, route, patch, enum.UpdateRoute)
	return route, nil
}

func (svc *RouteServiceImpl) Search(g *gin.Context, query map[string]any) (routes []model.Route, err error) {
	warrantID := ctx.GetMandatoryQueryParam(g, "warrantId")
	svc.WarrantService.CheckWarrantGrants(g, warrantID, false)

	routes = []model.Route{}
	result := svc.RouteRepo.Search(g, &routes, query, repo.Preload("StartAddress"), repo.Preload("EndAddress"))
	utils.Handle(result.Error)
	return routes, nil
}
