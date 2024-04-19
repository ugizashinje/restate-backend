package v1

import (
	"warrant-api/pkg/services"
	"warrant-api/pkg/utils"
	"warrant-api/pkg/wire"

	"github.com/gin-gonic/gin"
)

func RegisterDriverCategory(g *gin.RouterGroup) {
	v1 := g.Group("driverCategories")
	// v1.GET("", SearchAllDriverCategorys)
	v1.GET("/:id", getDriverCategory)
	v1.PATCH("/:id", patchDriverCategory)
	v1.POST("/", createDriverCategory)
	v1.DELETE("/:id", deleteDriverCategory)
}

// func SearchAllDriverCategorys(c *gin.Context) {
// 	drivercategoryes, err := wire.Svc.DriverCategoryService.Search(c, nil)
// 	if err != nil {
// 		c.PureJSON(500, struct{ error string }{error: err.Error()})
// 	}
// 	c.PureJSON(200, drivercategoryes)
// }

func getDriverCategory(c *gin.Context) {
	idParam := c.Param("id")
	id := string([]byte(idParam))
	drivercategory, _ := wire.Svc.DriverCategoryService.GetById(c, id)
	c.PureJSON(200, drivercategory)
}

func patchDriverCategory(c *gin.Context) {
	idParam := c.Param("id")
	id := string([]byte(idParam))
	var patch services.PatchDriverCategoryRequest
	err := c.BindJSON(&patch)
	utils.Handle(err)

	drivercategory, _ := wire.Svc.DriverCategoryService.Update(c, id, patch)
	c.PureJSON(200, drivercategory)
}

func createDriverCategory(c *gin.Context) {
	var createDriverCategory services.CreateDriverCategoryRequest
	err := c.BindJSON(&createDriverCategory)
	utils.Handle(err)
	drivercategory, err := wire.Svc.DriverCategoryService.Create(c, createDriverCategory)
	if err != nil {
		c.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	c.PureJSON(200, drivercategory)
}

func deleteDriverCategory(c *gin.Context) {
	idParam := c.Param("id")
	id := string([]byte(idParam))

	drivercategory, err := wire.Svc.DriverCategoryService.Delete(c, id)
	if err != nil {
		c.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	c.PureJSON(200, drivercategory)
}
