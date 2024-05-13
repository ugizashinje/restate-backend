package v1

import (
	"strings"
	"warrant-api/pkg/messages"
	"warrant-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {

	v1 := engine.Group("/v1")
	RegisterAddress(v1)
	RegisterCompany(v1)
	RegisterUser(v1)
	RegisterAuth(v1)
}

func bind(g *gin.Context, req interface{}) {
	err := g.BindJSON(req)
	if err != nil && strings.HasPrefix(err.Error(), "json") {
		utils.Handle(messages.Errorf(404, err.Error()))
	}
	utils.Handle(err)
}
