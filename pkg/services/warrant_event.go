package services

import (
	"warrant-api/pkg/ctx"
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/enum"
	"warrant-api/pkg/messages"
	"warrant-api/pkg/repo"
	"warrant-api/pkg/session"
	"warrant-api/pkg/utils"

	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v4"
	"gorm.io/datatypes"
)

type WarrantEventServiceImpl struct {
	WarrantRepo      repo.Repo[model.Warrant]
	WarrantEventRepo repo.Repo[model.WarrantEvent]
}

type CreateWarrantEventRequest struct {
	Time      string         `json:"time"`
	WarrantID string         `json:"warrantId"`
	RouteID   null.String    `json:"routeId"`
	Event     enum.EventType `json:"event"`
	UserID    string         `json:"userId"`
	Meta      datatypes.JSON `json:"meta"`
}

func (svc *WarrantEventServiceImpl) Create(g *gin.Context, createWarrantEvent CreateWarrantEventRequest) (*model.WarrantEvent, error) {

	warrantevent := &model.WarrantEvent{
		WarrantID: createWarrantEvent.WarrantID,
		RouteID:   createWarrantEvent.RouteID,
		Event:     createWarrantEvent.Event,
		UserID:    createWarrantEvent.UserID,
		Meta:      createWarrantEvent.Meta,
	}
	result := svc.WarrantEventRepo.Create(g, warrantevent)
	utils.Handle(result.Error)
	return warrantevent, nil
}

func (svc *WarrantEventServiceImpl) Delete(g *gin.Context, warranteventID string) (warrantevent *model.WarrantEvent, err error) {
	warrantevent = &model.WarrantEvent{}
	result := svc.WarrantEventRepo.Delete(g, warrantevent, warranteventID)
	utils.Handle(result.Error)

	return warrantevent, nil
}

func (svc *WarrantEventServiceImpl) GetById(c *gin.Context, warranteventID string) (warrantevent *model.WarrantEvent, err error) {
	warrantevent = &model.WarrantEvent{}
	result := svc.WarrantEventRepo.GetById(c, warrantevent, warranteventID)
	utils.Handle(result.Error)
	return warrantevent, nil
}

func (svc *WarrantEventServiceImpl) Update(c *gin.Context, warranteventID string, patch any) (warrantevent *model.WarrantEvent, err error) {
	warrantevent = &model.WarrantEvent{}
	result := svc.WarrantEventRepo.Patch(c, warrantevent, warranteventID, patch)
	utils.Handle(result.Error)

	return warrantevent, nil
}

func (svc *WarrantEventServiceImpl) Search(g *gin.Context, query map[string]any) (warrantevents []model.WarrantEvent, err error) {
	warrantevents = []model.WarrantEvent{}
	userSession := session.GetSession(g)
	warrantId := ctx.GetMandatoryQueryParam(g, "warrantId")
	warrant := &model.Warrant{}
	dbRes := svc.WarrantRepo.GetById(g, warrant, warrantId)
	utils.Handle(dbRes.Error)
	if isManager, isDispatcher, isDriver := userSession.Roles(warrant.CompanyID); !isManager && !isDispatcher && !(isDriver && warrant.DriverID == userSession.User.ID) {
		utils.Handle(messages.Unauthorized())
	}
	result := svc.WarrantEventRepo.Search(g, &warrantevents, map[string]any{"warrant_id": warrantId})
	utils.Handle(result.Error)
	return warrantevents, nil
}
