package repo

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"warrant-api/pkg/ctx"
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/enum"
	"warrant-api/pkg/messages"
	"warrant-api/pkg/session"
	"warrant-api/pkg/utils"

	"errors"

	"github.com/GoWebProd/uuid7"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repo[T model.BaseInterface] struct {
	Model T
}

type RepoOption func(g *gin.Context, d *gorm.DB) *gorm.DB

func Preload(tableName string) RepoOption {
	return func(g *gin.Context, db *gorm.DB) *gorm.DB {
		return db.Preload(tableName)
	}
}

func GrantUser(g *gin.Context, db *gorm.DB) *gorm.DB {
	userSession := session.GetSession(g)
	if userSession == nil {
		_, ok := g.Get(ctx.Email)
		if !ok {
			utils.Handle(messages.Errorf(404, "User not found"))
		}
		return db
	}

	managerCompanies := []string{}
	dispatcherCompanies := []string{}
	driverCompanies := []string{}

	for _, userCompany := range userSession.User.Companies {
		isManager, isDispatcher, isDriver := userSession.Roles(userCompany.CompanyID)
		if isManager {
			managerCompanies = append(managerCompanies, userCompany.CompanyID)
		} else if isDispatcher {
			dispatcherCompanies = append(dispatcherCompanies, userCompany.CompanyID)
		} else if isDriver {
			driverCompanies = append(driverCompanies, userCompany.CompanyID)
		}
	}
	userCompaniesWhere := ""
	if queryRole, ok := g.Request.URL.Query()["role"]; ok && enum.RoleValid(queryRole[0]) {
		userCompaniesWhere += fmt.Sprintf(" AND ('{%s}' <@ user_companies.roles )", queryRole[0])
	}
	if queryCompanyID, ok := g.Request.URL.Query()["companyId"]; ok {
		userCompaniesWhere += fmt.Sprintf(" AND (user_companies.company_id = '%s') ", queryCompanyID[0])
	}

	db = db.Where(fmt.Sprintf(` EXISTS	
	( SELECT 1 FROM user_companies WHERE 
		users.id = user_companies.user_id %s AND
		
		(
			(user_companies.company_id in (?) )
			OR ('{driver,dispatcher}' @> user_companies.roles AND user_companies.company_id in (?) )
			OR ('{dispatcher}' <@ user_companies.roles AND user_companies.company_id in (?))
		)
   	)`, userCompaniesWhere), managerCompanies, dispatcherCompanies, driverCompanies)

	if _, ok := g.Request.URL.Query()["category"]; ok {
		db = db.Preload("Categories")
	}
	return db

}

func GrantWarrant(g *gin.Context, db *gorm.DB) *gorm.DB {

	var companies = []string{}
	var driverCompanies = []string{}
	userSession := session.GetSession(g)
	if userSession == nil {
		utils.Handle(messages.Errorf(http.StatusUnauthorized, "You are not allowed to perform this operation"))
	}
	for _, v := range userSession.User.Companies {
		isManager, isDispatcher, isDriver := userSession.Roles(v.CompanyID)
		if v.IsActive && !v.IsDeleted && (isManager || isDispatcher) {
			companies = append(companies, v.CompanyID)
		} else if isDriver {
			driverCompanies = append(driverCompanies, v.CompanyID)
		}
	}
	if len(companies) == 0 && len(driverCompanies) == 0 {
		return db.Where("1 = 0")
	} else if len(companies) > 0 && len(driverCompanies) == 0 {
		return db.Where("company_id in ? ", companies)
	} else if len(companies) == 0 && len(driverCompanies) > 0 {
		return db.Where("company_id in ? and driver_id = ?", driverCompanies, userSession.User.ID)
	} else {
		return db.Where("( (company_id in ?) OR (company_id in ? and driver_id = ?))", companies, driverCompanies, userSession.User.ID)
	}
}

func GrantRepair(g *gin.Context, db *gorm.DB) *gorm.DB {
	userSession := session.GetSession(g)
	if userSession == nil {
		utils.Handle(messages.Errorf(http.StatusUnauthorized, "You are not allowed to perform this operation"))
	}
	warrantRepo := Repo[model.Warrant]{}
	warrantId := g.Request.URL.Query()["warrantId"][0]
	warrant := &model.Warrant{}
	dbRes := warrantRepo.GetById(g, warrant, warrantId)
	utils.Handle(dbRes.Error)
	isManager, isDispatcher, isDriver := userSession.Roles(warrant.CompanyID)
	if isManager || isDispatcher ||
		(isDriver && (warrant.DriverID == userSession.User.ID)) {
		return db.Where("warrant_id", warrant.ID)
	}
	utils.Handle(messages.Errorf(http.StatusUnauthorized, "You are not allowed to perform this operation"))
	return nil
}

func GrantVehicle(g *gin.Context, db *gorm.DB) *gorm.DB {
	userSession := session.GetSession(g)
	if userSession == nil {
		_, ok := g.Get(ctx.Email)
		if !ok {
			utils.Handle(messages.Errorf(404, "User not found"))
		}
		return db
	}
	companies := []string{}
	for _, company := range userSession.User.Companies {
		isManager, isDispatcher, _ := userSession.Roles(company.CompanyID)
		if isManager || isDispatcher {
			companies = append(companies, company.CompanyID)
		}
	}
	return db.Where("company_id in (?)", companies)
}

func JustFirst(g *gin.Context, db *gorm.DB) *gorm.DB {
	return db.Limit(1)
}

func (r Repo[T]) Create(g *gin.Context, t *T, opts ...RepoOption) *gorm.DB {
	database, ok := g.Get(ctx.Transaction)
	if !ok {
		utils.Handle(errors.New("transaction not initalized"))
	}

	db := database.(*gorm.DB)
	for _, option := range opts {
		db = option(g, db)
	}
	return db.Clauses(clause.Returning{}).Create(&t)
}

func (r Repo[T]) Save(g *gin.Context, t *T) *gorm.DB {
	database, ok := g.Get(ctx.Transaction)
	if !ok {
		utils.Handle(errors.New("transaction not initalized"))
	}
	db := database.(*gorm.DB)
	return db.Omit(clause.Associations).Clauses(clause.Returning{}).Save(&t)
}
func (r Repo[T]) Patch(g *gin.Context, t *T, id string, patch any) *gorm.DB {
	database, ok := g.Get(ctx.Transaction)
	if !ok {
		utils.Handle(errors.New("transaction not initalized"))
	}
	db := database.(*gorm.DB)
	db.Model(t).Clauses(clause.Returning{}).Where("id = ?", id).Updates(patch)
	return db
}

func (r Repo[T]) Delete(g *gin.Context, t *T, id string) *gorm.DB {
	database, ok := g.Get(ctx.Transaction)
	if !ok {
		utils.Handle(errors.New("transaction not initalized"))
	}
	db := database.(*gorm.DB)

	patch := map[string]interface{}{
		"is_deleted": true,
		"is_active":  false,
		"deleted_at": time.Now(),
	}
	db.Model(t).Clauses(clause.Returning{}).Where("id = ?", id).Updates(patch)
	return db
}

func (r Repo[T]) GetById(g *gin.Context, t *T, id string, opts ...RepoOption) *gorm.DB {
	database, ok := g.Get(ctx.Transaction)
	if !ok {
		utils.Handle(errors.New("transaction not initalized"))
	}
	db := database.(*gorm.DB)
	for _, option := range opts {
		db = option(g, db)
	}
	return db.First(&t, "id = ?", id)
}

func (r Repo[T]) Search(g *gin.Context, t *[]T, filter map[string]any, opts ...RepoOption) *gorm.DB {
	database, ok := g.Get(ctx.Transaction)
	if !ok {
		utils.Handle(errors.New("transaction not initalized"))
	}
	db := database.(*gorm.DB)

	db = db.Model(t).Where(filter)

	if queryFilter := utils.QueryFilter(g, db, t); len(queryFilter) > 0 {
		for k, v := range queryFilter {
			db = db.Where(k, v)
		}
	}
	page, _ := strconv.Atoi(g.Query("page"))
	size, _ := strconv.Atoi(g.Query("pageSize"))
	if (size == 0) || (page == 0) {
		page = 1
		size = 20
	}
	offset := (page - 1) * size
	db = db.Offset(offset).Limit(size)

	for _, option := range opts {
		db = option(g, db)
	}

	db.
		Where(`"`+r.Model.TableName()+`"."is_deleted" = ? `, false).
		Find(&t)
	return db
}

var GenUUID = uuid7.New()
