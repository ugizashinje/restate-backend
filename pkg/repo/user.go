package repo

import (
	"errors"
	"warrant-api/pkg/ctx"
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/utils"

	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type User struct {
	Repo[model.User]
}

func (r User) ManageCompanies(g *gin.Context, user *model.User, addToCompany null.String, roles []string, removeFromCompany null.String) error {
	database, ok := g.Get(ctx.Transaction)
	if !ok {
		utils.Handle(errors.New("transaction not initalized"))
	}
	db := database.(*gorm.DB)

	if addToCompany.Valid {
		company := &model.Company{BaseModel: model.BaseModel{ID: addToCompany.String, IsDeleted: false}}
		dbRes := db.First(company)
		utils.Handle(dbRes.Error)
		if dbRes.RowsAffected == 0 {
			return errors.New("trying to add non existant company")
		}
		userCompany := &model.UserCompany{UserID: user.ID, CompanyID: company.ID}
		dbRes = db.First(&userCompany)
		if dbRes.RowsAffected == 0 {
			userCompany.IsActive = true
			userCompany.Roles = roles
		}

		if dbRes := db.Save(userCompany); dbRes.Error != nil {
			return dbRes.Error
		}

	}
	if removeFromCompany.Valid {
		company := &model.Company{BaseModel: model.BaseModel{ID: removeFromCompany.String}}
		dbRes := db.First(company)
		utils.Handle(dbRes.Error)
		if dbRes.RowsAffected == 0 {
			return errors.New("trying to add non existant company")
		}
		userCompany := &model.UserCompany{UserID: user.ID, CompanyID: company.ID, IsActive: false}
		if dbRes := db.Save(userCompany); dbRes.Error != nil {
			return dbRes.Error
		}
	}
	return nil
}
