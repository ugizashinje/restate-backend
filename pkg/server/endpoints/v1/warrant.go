package v1

import (
	"warrant-api/pkg/services"
	"warrant-api/pkg/utils"
	"warrant-api/pkg/wire"

	"github.com/gin-gonic/gin"
)

func RegisterWarrant(g *gin.RouterGroup) {
	v1 := g.Group("warrants")
	v1.GET("/", SearchAllWarrant)
	v1.GET("/:id", getWarrant)
	v1.PATCH("/:id", patchWarrant)
	v1.POST("/", createWarrant)
	v1.DELETE("/:id", deleteWarrant)
}

func SearchAllWarrant(c *gin.Context) {
	warrantes, err := wire.Svc.WarrantService.Search(c, nil)
	if err != nil {
		c.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	c.PureJSON(200, warrantes)
}

func getWarrant(c *gin.Context) {
	idParam := c.Param("id")
	id := string([]byte(idParam))
	warrant, _ := wire.Svc.WarrantService.GetById(c, id)
	c.PureJSON(200, warrant)
}

func patchWarrant(c *gin.Context) {
	idParam := c.Param("id")
	id := string([]byte(idParam))
	var patch services.PatchWarrantRequest
	err := c.BindJSON(&patch)
	utils.Handle(err)

	warrant, _ := wire.Svc.WarrantService.Update(c, id, patch)
	c.PureJSON(200, warrant)

}

func createWarrant(g *gin.Context) {
	var createWarrant services.CreateWarrantRequest
	bind(g, &createWarrant)
	warrant, err := wire.Svc.WarrantService.Create(g, createWarrant)
	if err != nil {
		g.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	g.PureJSON(200, warrant)
}

func deleteWarrant(c *gin.Context) {
	idParam := c.Param("id")
	id := string([]byte(idParam))

	warrant, err := wire.Svc.WarrantService.Delete(c, id)
	if err != nil {
		c.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	c.PureJSON(200, warrant)

}
