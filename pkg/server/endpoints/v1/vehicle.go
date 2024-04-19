package v1

import (
	"warrant-api/pkg/services"
	"warrant-api/pkg/utils"
	"warrant-api/pkg/wire"

	"github.com/gin-gonic/gin"
)

func RegisterVehicle(g *gin.RouterGroup) {
	v1 := g.Group("vehicles")
	v1.GET("/", SearchAllVehicles)
	v1.GET("/:id", getVehicle)
	v1.PATCH("/:id", patchVehicle)
	v1.POST("/", createVehicle)
	v1.DELETE("/:id", deleteVehicle)
}

func SearchAllVehicles(c *gin.Context) {
	vehiclees, err := wire.Svc.VehicleService.Search(c, nil)
	if err != nil {
		c.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	c.PureJSON(200, vehiclees)
}

func getVehicle(c *gin.Context) {
	idParam := c.Param("id")
	id := string([]byte(idParam))
	vehicle, _ := wire.Svc.VehicleService.GetById(c, id)
	c.PureJSON(200, vehicle)
}

func patchVehicle(c *gin.Context) {
	idParam := c.Param("id")
	id := string([]byte(idParam))
	var patch services.PatchVehicleRequest
	err := c.BindJSON(&patch)
	utils.Handle(err)

	vehicle, _ := wire.Svc.VehicleService.Update(c, id, patch)
	c.PureJSON(200, vehicle)
}

func createVehicle(g *gin.Context) {
	var createVehicle services.CreateVehicleRequest
	bind(g, &createVehicle)
	vehicle, err := wire.Svc.VehicleService.Create(g, createVehicle)
	if err != nil {
		g.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	g.PureJSON(200, vehicle)
}

func deleteVehicle(c *gin.Context) {
	idParam := c.Param("id")
	id := string([]byte(idParam))

	vehicle, err := wire.Svc.VehicleService.Delete(c, id)
	if err != nil {
		c.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	c.PureJSON(200, vehicle)
}
