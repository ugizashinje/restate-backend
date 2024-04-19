package v1

import (
	"warrant-api/pkg/services"
	"warrant-api/pkg/utils"
	"warrant-api/pkg/wire"

	"github.com/gin-gonic/gin"
)

func RegisterCompany(g *gin.RouterGroup) {
	v1 := g.Group("companies")
	v1.GET("/", SearchAllCompanies)
	v1.GET("/:id", getCompany)
	v1.PATCH("/:id", patchCompany)
	// v1.POST("/", createCompany)
	v1.DELETE("/:id", deleteCompany)
}

func SearchAllCompanies(c *gin.Context) {
	companyes, err := wire.Svc.CompanyService.Search(c, nil)
	if err != nil {
		c.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	c.PureJSON(200, companyes)
}

func getCompany(c *gin.Context) {
	idParam := c.Param("id")
	id := string([]byte(idParam))
	company, _ := wire.Svc.CompanyService.GetById(c, id)
	c.PureJSON(200, company)
}

func patchCompany(c *gin.Context) {
	idParam := c.Param("id")
	id := string([]byte(idParam))
	var patch services.PatchCompanyRequest
	err := c.BindJSON(&patch)
	utils.Handle(err)

	company, _ := wire.Svc.CompanyService.Update(c, id, patch)
	c.PureJSON(200, company)

}

func createCompany(c *gin.Context) {
	var createCompany services.CreateCompanyRequest
	err := c.BindJSON(&createCompany)
	utils.Handle(err)
	company, err := wire.Svc.CompanyService.Create(c, createCompany)
	if err != nil {
		c.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	c.PureJSON(200, company)
}

func deleteCompany(c *gin.Context) {
	idParam := c.Param("id")
	id := string([]byte(idParam))

	company, err := wire.Svc.CompanyService.Delete(c, id)
	if err != nil {
		c.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	c.PureJSON(200, company)

}
