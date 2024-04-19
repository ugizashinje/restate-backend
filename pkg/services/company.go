package services

import (
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/repo"
	"warrant-api/pkg/utils"

	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v4"
	"gorm.io/datatypes"
)

type CompanyServiceImpl struct {
	CompanyRepo    repo.Repo[model.Company]
	AddressService AddressServiceImpl
}

type PatchCompanyRequest struct {

	// Under serbian law name of copmany is immutable
	// PIB MN also
	// Name      null.String `json:"name"`
	// Pib       null.String `json:"pib"`
	// MN        null.String `json:"mn"`
	Short     null.String `json:"short"`
	AddressID null.String `json:"addressId"`
	Phone     null.String `json:"phone"`
	Email     null.String `json:"email"`
}

type CreateCompanyRequest struct {
	Name      string         `json:"name"`
	Short     string         `json:"short"`
	PIB       string         `json:"pib"`
	AddressID string         `json:"addressId"`
	Mn        string         `json:"mn"`
	Phone     string         `json:"phone"`
	Email     string         `json:"email" binding:"required,email"`
	Meta      datatypes.JSON `json:"meta"`
}

func (svc *CompanyServiceImpl) Create(g *gin.Context, createCompany CreateCompanyRequest) (*model.Company, error) {
	company := &model.Company{
		AddressID: createCompany.AddressID,
		Short:     createCompany.Short,
		Name:      createCompany.Name,
		PIB:       createCompany.PIB,
		Mn:        createCompany.Mn,
		Phone:     createCompany.Phone,
		Email:     createCompany.Email,
		Meta:      createCompany.Meta}

	result := svc.CompanyRepo.Create(g, company)
	utils.Handle(result.Error)
	return company, nil
}

func (svc *CompanyServiceImpl) Delete(g *gin.Context, companyID string) (company *model.Company, err error) {
	company = &model.Company{}
	result := svc.CompanyRepo.Delete(g, company, companyID)
	utils.Handle(result.Error)

	return company, nil
}

func (svc *CompanyServiceImpl) GetById(c *gin.Context, companyID string) (company *model.Company, err error) {
	company = &model.Company{}
	result := svc.CompanyRepo.GetById(c, company, companyID)
	utils.Handle(result.Error)
	return company, nil
}

func (svc *CompanyServiceImpl) Update(c *gin.Context, companyID string, patch any) (company *model.Company, err error) {
	company = &model.Company{}
	result := svc.CompanyRepo.Patch(c, company, companyID, patch)
	utils.Handle(result.Error)

	return company, nil
}

func (svc *CompanyServiceImpl) Search(c *gin.Context, query map[string]any) (companys []model.Company, err error) {
	companys = []model.Company{}
	result := svc.CompanyRepo.Search(c, &companys, query)
	utils.Handle(result.Error)
	return companys, nil
}
