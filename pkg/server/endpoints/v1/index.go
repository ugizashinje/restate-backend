package v1

import (
	"strings"
	"warrant-api/pkg/messages"
	"warrant-api/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	"gopkg.in/guregu/null.v3"
)

var validRoomsValues = map[string]bool{
	"0":   true,
	"0.5": true,
	"1":   true,
	"1.5": true,
	"2":   true,
	"2.5": true,
	"3":   true,
	"3.5": true,
	"4":   true,
	"4+":  true,
}

func roomsValidator(fl validator.FieldLevel) bool {

	nullStatus, nullOk := fl.Field().Interface().(null.String)
	if nullOk && nullStatus.Valid {
		return validRoomsValues[nullStatus.String]
	}
	status := fl.Field().String()
	return validRoomsValues[status]

}

func Register(engine *gin.Engine) {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("roomsValidator", roomsValidator); err != nil {
			utils.Handle(err)
		}

	}
	v1 := engine.Group("/v1")
	RegisterLocations(v1)
	RegisterAddress(v1)
	RegisterApartment(v1)
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
