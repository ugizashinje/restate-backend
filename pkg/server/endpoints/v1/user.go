package v1

import (
	"warrant-api/pkg/services"
	"warrant-api/pkg/utils"
	"warrant-api/pkg/wire"

	"github.com/gin-gonic/gin"
)

func RegisterUser(g *gin.RouterGroup) {
	v1 := g.Group("users")
	v1.GET("/", SearchAlluser)
	v1.GET("/:id", getuser)
	v1.PATCH("/:id", patchuser)
	v1.POST("/", createuser)
	v1.DELETE("/:id", deleteuser)
}

func SearchAlluser(c *gin.Context) {
	useres, err := wire.Svc.UserService.Search(c, nil)
	if err != nil {
		c.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	c.PureJSON(200, useres)
}

func getuser(c *gin.Context) {
	idParam := c.Param("id")
	id := string([]byte(idParam))
	user, _ := wire.Svc.UserService.GetById(c, id)
	c.PureJSON(200, user)
}

func patchuser(c *gin.Context) {
	idParam := c.Param("id")
	id := string([]byte(idParam))
	var patch services.PatchUserRequest
	err := c.BindJSON(&patch)
	utils.Handle(err)

	user, _ := wire.Svc.UserService.Update(c, id, patch)
	c.PureJSON(200, user)
}

func createuser(c *gin.Context) {
	var createuser services.CreateUserRequest
	err := c.BindJSON(&createuser)
	utils.Handle(err)
	user, err := wire.Svc.UserService.Create(c, createuser)
	if err != nil {
		c.PureJSON(500, gin.H{"error": err.Error()})
		return
	}
	c.PureJSON(200, user)
}

func deleteuser(c *gin.Context) {
	idParam := c.Param("id")
	id := string([]byte(idParam))

	user, err := wire.Svc.UserService.Delete(c, id)
	if err != nil {
		c.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	c.PureJSON(200, user)

}
