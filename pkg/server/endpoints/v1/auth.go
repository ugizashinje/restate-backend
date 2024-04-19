package v1

import (
	"warrant-api/pkg/services"
	"warrant-api/pkg/utils"
	"warrant-api/pkg/wire"

	"github.com/gin-gonic/gin"
)

func RegisterAuth(v1 *gin.RouterGroup) {
	auth := v1.Group("/auth")
	auth.POST("/login", login)
	auth.GET("/logout", logout)
	auth.POST("/refresh", refresh)
	auth.POST("/register", registerCompanyAndManager)
	auth.GET("/confirm/:code", confirm)
	auth.GET("/superset/", getDashboardToken)

}

func login(g *gin.Context) {
	login := services.LoginRequest{}
	err := g.BindJSON(&login)
	utils.Handle(err)

	result, err := wire.Svc.AuthService.Login(g, login)
	utils.Handle(err)

	g.JSON(200, result)
}

func logout(g *gin.Context) {
	err := wire.Svc.AuthService.Logout(g)
	utils.Handle(err)
	g.JSON(200, gin.H{"success": gin.Default().TrustedPlatform})
}

func refresh(g *gin.Context) {
	refresh := services.RefreshRequest{}
	err := g.BindJSON(&refresh)
	utils.Handle(err)
	access_token, err := wire.Svc.AuthService.Refresh(g, refresh)
	utils.Handle(err)
	g.JSON(200, gin.H{"access_token": access_token})
}

func registerCompanyAndManager(g *gin.Context) {
	register := services.RegisterRequest{}
	err := g.BindJSON(&register)
	utils.Handle(err)
	err = wire.Svc.AuthService.Register(g, register)
	utils.Handle(err)
	g.JSON(200, gin.H{"register": "success"})
}

func confirm(g *gin.Context) {
	code := g.Param("code")
	confirmation, err := wire.Svc.AuthService.Confirm(g, code)
	utils.Handle(err)
	g.JSON(200, confirmation)
}

func getDashboardToken(g *gin.Context) {
	token, err := wire.Svc.AuthService.DashboardToken(g)
	utils.Handle(err)
	g.JSON(200, token)

}
