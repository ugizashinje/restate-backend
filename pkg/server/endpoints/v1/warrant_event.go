package v1

import (
	"warrant-api/pkg/wire"

	"github.com/gin-gonic/gin"
)

func RegisterWarrantEvent(g *gin.RouterGroup) {
	v1 := g.Group("warrantevents")
	v1.GET("/", SearchAllWarrantEvent)
}

func SearchAllWarrantEvent(c *gin.Context) {
	warranteventes, err := wire.Svc.WarrantEventService.Search(c, nil)
	if err != nil {
		c.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	c.PureJSON(200, warranteventes)
}
