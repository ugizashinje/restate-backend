package v1

import (
	"warrant-api/pkg/services"
	"warrant-api/pkg/utils"
	"warrant-api/pkg/wire"

	"github.com/gin-gonic/gin"
)

func RegisterRepair(g *gin.RouterGroup) {
	v1 := g.Group("repairs")
	v1.GET("/", SearchAllRepair)
	v1.GET("/:id", getRepair)
	v1.PATCH("/:id", patchRepair)
	v1.POST("/", createRepair)
	v1.DELETE("/:id", deleteRepair)
}

func SearchAllRepair(c *gin.Context) {
	repaires, err := wire.Svc.RepairService.Search(c, nil)
	if err != nil {
		c.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	c.PureJSON(200, repaires)
}

func getRepair(c *gin.Context) {
	idParam := c.Param("id")
	id := string([]byte(idParam))
	repair, _ := wire.Svc.RepairService.GetById(c, id)
	c.PureJSON(200, repair)
}

func patchRepair(c *gin.Context) {
	idParam := c.Param("id")
	id := string([]byte(idParam))
	var patch services.PatchRepairRequest
	err := c.BindJSON(&patch)
	utils.Handle(err)

	repair, _ := wire.Svc.RepairService.Update(c, id, patch)
	c.PureJSON(200, repair)

}

func createRepair(c *gin.Context) {
	var createRepair services.CreateRepairRequest
	err := c.BindJSON(&createRepair)
	utils.Handle(err)
	repair, err := wire.Svc.RepairService.Create(c, createRepair)
	if err != nil {
		c.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	c.PureJSON(200, repair)
}

func deleteRepair(c *gin.Context) {
	idParam := c.Param("id")
	id := string([]byte(idParam))

	repair, err := wire.Svc.RepairService.Delete(c, id)
	if err != nil {
		c.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	c.PureJSON(200, repair)

}
