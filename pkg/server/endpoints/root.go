package endpoints

import (
	v1 "warrant-api/pkg/server/endpoints/v1"

	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {

	engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	v1.Register(engine)
}
