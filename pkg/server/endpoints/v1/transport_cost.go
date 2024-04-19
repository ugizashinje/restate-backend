package v1

import (
	"warrant-api/pkg/services"
	"warrant-api/pkg/utils"
	"warrant-api/pkg/wire"

	"github.com/gin-gonic/gin"
)

func RegisterTransportCost(g *gin.RouterGroup) {
	v1 := g.Group("transportCosts")
	v1.GET("/", SearchAllTransportCost)
	v1.GET("/:id", getTransportCost)
	v1.PATCH("/:id", patchTransportCost)
	v1.POST("/", createTransportCost)
	v1.DELETE("/:id", deleteTransportCost)
}

func SearchAllTransportCost(c *gin.Context) {
	transportcostes, err := wire.Svc.TransportCostService.Search(c, make(map[string]any))
	if err != nil {
		c.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	c.PureJSON(200, transportcostes)
}

func getTransportCost(c *gin.Context) {
	idParam := c.Param("id")
	id := string([]byte(idParam))
	transportcost, _ := wire.Svc.TransportCostService.GetById(c, id)
	c.PureJSON(200, transportcost)
}

func patchTransportCost(c *gin.Context) {
	idParam := c.Param("id")
	id := string([]byte(idParam))
	var patch services.PatchTransportCostRequest
	err := c.BindJSON(&patch)
	utils.Handle(err)

	transportcost, _ := wire.Svc.TransportCostService.Update(c, id, patch)
	c.PureJSON(200, transportcost)

}

func createTransportCost(c *gin.Context) {
	var createTransportCost services.CreateTransportCostRequest
	err := c.BindJSON(&createTransportCost)
	utils.Handle(err)
	transportcost, err := wire.Svc.TransportCostService.Create(c, createTransportCost)
	if err != nil {
		c.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	c.PureJSON(200, transportcost)
}

func deleteTransportCost(c *gin.Context) {
	idParam := c.Param("id")
	id := string([]byte(idParam))

	transportcost, err := wire.Svc.TransportCostService.Delete(c, id)
	if err != nil {
		c.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	c.PureJSON(200, transportcost)

}
