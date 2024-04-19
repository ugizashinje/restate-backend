package server

import (
	docs "warrant-api/docs"
	"warrant-api/pkg/server/endpoints"
	"warrant-api/pkg/validation"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func init() {
	log.SetLevel(log.InfoLevel) // NEW
}

func Start(gormDB *gorm.DB) error {

	r := gin.Default()

	// ("date", dateValidations)
	gin.EnableJsonDecoderDisallowUnknownFields()
	validation.Init()
	r.Use(recovery())
	r.Use(dbMiddleware(gormDB))
	r.Use(corsMiddleware())
	r.Use(authMiddleware())
	endpoints.Register(r)
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Host = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":9876")
	return nil
}
