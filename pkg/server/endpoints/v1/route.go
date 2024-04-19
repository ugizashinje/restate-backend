package v1

import (
	"warrant-api/pkg/services"
	"warrant-api/pkg/utils"
	"warrant-api/pkg/wire"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(g *gin.RouterGroup) {
	v1 := g.Group("routes")
	v1.GET("/", SearchAllRoute)
	v1.GET("/:id", getRoute)
	v1.PATCH("/:id", patchRoute)
	v1.POST("/", createRoute)
	v1.DELETE("/:id", deleteRoute)
}

func SearchAllRoute(c *gin.Context) {
	routees, err := wire.Svc.RouteService.Search(c, nil)
	if err != nil {
		c.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	c.PureJSON(200, routees)
}

func getRoute(c *gin.Context) {
	idParam := c.Param("id")
	id := string([]byte(idParam))
	route, _ := wire.Svc.RouteService.GetById(c, id)
	c.PureJSON(200, route)
}

func patchRoute(c *gin.Context) {
	idParam := c.Param("id")
	id := string([]byte(idParam))
	var patch services.PatchRouteRequest
	err := c.BindJSON(&patch)
	utils.Handle(err)

	route, _ := wire.Svc.RouteService.Update(c, id, patch)
	c.PureJSON(200, route)
}

func createRoute(c *gin.Context) {
	var createRoute services.CreateRouteRequest
	err := c.BindJSON(&createRoute)
	utils.Handle(err)
	route, err := wire.Svc.RouteService.Create(c, createRoute)
	if err != nil {
		c.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	c.PureJSON(200, route)
}

func deleteRoute(c *gin.Context) {
	idParam := c.Param("id")
	id := string([]byte(idParam))
	route, err := wire.Svc.RouteService.Delete(c, id)
	if err != nil {
		c.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	c.PureJSON(200, route)
}
