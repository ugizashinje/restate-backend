package services

import (
	"net/http"
	"time"
	"warrant-api/pkg/config"
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/enum"
	"warrant-api/pkg/messages"
	"warrant-api/pkg/repo"
	"warrant-api/pkg/session"
	"warrant-api/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"gopkg.in/guregu/null.v4"
)

type WarrantServiceImpl struct {
	ChangeRepo  repo.Repo[model.Change]
	WarrantRepo repo.Repo[model.Warrant]
}

type CreateWarrantRequest struct {
	ExpectedStart        string             `json:"expectedStart" binding:"required,time"`
	DriverID             string             `json:"driverId"`
	VehicleID            string             `json:"vehicleId"`
	TrailerID            null.String        `json:"trailerId"`
	CompanyID            string             `json:"companyId"`
	DispatcherID         string             `json:"dispatcherId"`
	TechnicalCorrectness string             `json:"technicalCorrectness"`
	Passangers           pq.StringArray     `json:"passangers"`
	Status               enum.WarrantStatus `json:"status" binding:"omitempty,warrantStatus"`
	Name                 string             `json:"name"`
	Note                 null.String        `json:"note"`
}
type PrimitivePatch struct {
	ExpectedStart       null.Time   `json:"expectedStart" binding:"omitempty,time"`
	DriverID            null.String `json:"driverId,omitempty"`
	VehicleID           null.String `json:"vehicleId,omitempty"`
	CompanyID           null.String `json:"companyId,omitempty"`
	DispatcherID        null.String `json:"dispatcherId,omitempty"`
	Status              null.String `json:"status,omitempty" binding:"omitempty,warrantStatus"`
	TechicalCorrectness null.String `json:"techicalCorrectness,omitempty"`
	Name                null.String `json:"name,omitempty"`
	Note                null.String `json:"note,omitempty"`
}
type PatchWarrantRequest struct {
	PrimitivePatch
	AddPassenger    null.String `json:"addPassenger,omitempty"`
	RemovePassenger null.String `json:"removePassenger,omitempty"`
}

func (svc *WarrantServiceImpl) Create(g *gin.Context, createWarrant CreateWarrantRequest) (*model.Warrant, error) {

	userSession := session.GetSession(g)
	isManager, isDispatcher, _ := userSession.Roles(createWarrant.CompanyID)
	if !isManager && !isDispatcher {
		utils.Handle(messages.Errorf(401, "You do not have grants to create warrant."))
	}
	issueDate := time.Now()
	expectedStart, err := time.Parse(config.Format.TimeFormat, createWarrant.ExpectedStart)
	utils.Handle(err)

	warrant := &model.Warrant{
		IssueDate:            issueDate,
		ExpectedStart:        expectedStart,
		DriverID:             createWarrant.DriverID,
		VehicleID:            createWarrant.VehicleID,
		TrailerID:            createWarrant.TrailerID,
		CompanyID:            createWarrant.CompanyID,
		DispatcherID:         createWarrant.DispatcherID,
		Status:               createWarrant.Status,
		Passengers:           createWarrant.Passangers,
		TechnicalCorrectness: createWarrant.TechnicalCorrectness,
		Name:                 createWarrant.Name,
		Note:                 createWarrant.Note,
	}

	result := svc.WarrantRepo.Create(g, warrant)
	utils.Handle(result.Error)

	// userSession.LogEvent(g, warrant, nil, createWarrant, enum.CreateWarrant)

	return warrant, nil
}

func (svc *WarrantServiceImpl) Delete(g *gin.Context, warrantID string) (warrant *model.Warrant, err error) {
	// raw, _ := g.Get(ctx.Session)
	// userSession := raw.(*session.Session)
	svc.CheckWarrantGrants(g, warrantID, true)
	warrant = &model.Warrant{}
	result := svc.WarrantRepo.Delete(g, warrant, warrantID)
	utils.Handle(result.Error)

	// userSession.LogEvent(g, warrant, nil, nil, enum.DeleteWarrant)

	return warrant, nil
}

func (svc *WarrantServiceImpl) GetById(g *gin.Context, warrantID string) (warrant *model.Warrant, err error) {
	svc.CheckWarrantGrants(g, warrantID, false)
	warrant = &model.Warrant{}
	result := svc.WarrantRepo.GetById(g, warrant, warrantID)
	utils.Handle(result.Error)

	return warrant, nil
}

func (svc *WarrantServiceImpl) Update(g *gin.Context, warrantID string, patch PatchWarrantRequest) (warrant *model.Warrant, err error) {
	warrant = &model.Warrant{}
	result := svc.WarrantRepo.Patch(g, warrant, warrantID, patch.PrimitivePatch)
	utils.Handle(result.Error)

	if patch.AddPassenger.Valid {
		warrant.Passengers = append(warrant.Passengers, patch.AddPassenger.String)
	}
	if patch.RemovePassenger.Valid {
		passengers := pq.StringArray{}
		for _, v := range warrant.Passengers {
			if v == patch.RemovePassenger.String {
				continue
			}
			passengers = append(passengers, v)
		}
		warrant.Passengers = passengers
	}

	if patch.Status.Valid {
		svc.updateStatus(g, warrant, patch.Status.String)
	}

	result = svc.WarrantRepo.Save(g, warrant)
	utils.Handle(result.Error)

	// raw, _ := g.Get(ctx.Session)
	// userSession, _ := raw.(*session.Session)
	// userSession.LogEvent(g, warrant, nil, patch, enum.UpdateWarrant)
	return warrant, nil
}

func (svc *WarrantServiceImpl) updateStatus(g *gin.Context, warrant *model.Warrant, newStatus string) {
	if enum.WarrantStatus(warrant.Status) != enum.Assigned &&
		enum.WarrantStatus(newStatus) == enum.Assigned {
		warrant.IssueDate = time.Now()
	}
	if enum.WarrantStatus(warrant.Status) != enum.Completed &&
		enum.WarrantStatus(newStatus) == enum.Completed {
		warrant.ClosingDate = time.Now()
	}
}

func (svc *WarrantServiceImpl) Search(c *gin.Context, query map[string]any) (warrants []model.Warrant, err error) {
	warrants = []model.Warrant{}

	result := svc.WarrantRepo.Search(c, &warrants, query, repo.GrantWarrant)
	utils.Handle(result.Error)
	return warrants, nil
}

func (svc *WarrantServiceImpl) CheckWarrantGrants(g *gin.Context, warrantID string, write bool) *model.Warrant {
	if warrantID == "" {
		utils.Handle(messages.Errorf(http.StatusBadRequest, "Search query parameter warrantId is mandatory"))
	}
	warrantAccess := &model.Warrant{}
	dbRes := svc.WarrantRepo.GetById(g, warrantAccess, warrantID)
	utils.Handle(dbRes.Error)
	session := session.GetSession(g)
	isManager, isDispatcher, isDriver := session.Roles(warrantAccess.CompanyID)
	if (!isManager && !isDispatcher) && (!isDriver || warrantAccess.DriverID != session.User.ID) {
		utils.Handle(messages.Errorf(404, "Warrant not found"))
	}

	if (isManager || isDispatcher) && write {
		return warrantAccess
	}
	if isDriver && write {
		utils.Handle(messages.Errorf(401, "You are not allowed to change this warrant"))

	}
	return warrantAccess
}
