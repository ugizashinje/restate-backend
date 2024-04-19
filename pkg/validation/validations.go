package validation

import (
	"time"
	"warrant-api/pkg/config"
	"warrant-api/pkg/enum"
	"warrant-api/pkg/utils"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gopkg.in/guregu/null.v4"
)

func dateFormatValidation(fl validator.FieldLevel) bool {
	v, k, _ := fl.ExtractType(fl.Field())
	if k.String() == "string" {
		if v.String() != "" {
			_, err := time.Parse(config.Format.DateFormat, v.String())
			return err == nil
		}
		return true
	}

	if date, ok := fl.Field().Interface().(null.String); ok {
		if date.Valid {
			_, err := time.Parse(config.Format.DateFormat, date.String)
			return err == nil
		}
		return true
	}
	return false
}
func timeFormatValidation(fl validator.FieldLevel) bool {
	if date, ok := fl.Field().Interface().(null.String); ok {
		if date.Valid {
			_, err := time.Parse(config.Format.TimeFormat, date.String)
			if err != nil {
				return false
			}
		}
		return true
	}
	return true
}

func validStatus(fl validator.FieldLevel) bool {
	validStatusValues := map[string]bool{
		string(enum.Template):    true,
		string(enum.Preparation): true,
		string(enum.Assigned):    true,
		string(enum.Transit):     true,
		string(enum.Transport):   true,
		string(enum.Interrupted): true,
		string(enum.Completed):   true,
	}

	nullStatus, nullOk := fl.Field().Interface().(null.String)
	if nullOk && nullStatus.Valid {
		return validStatusValues[nullStatus.String]
	}
	status := fl.Field().String()
	return validStatusValues[status]

}

func Init() {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("date", dateFormatValidation); err != nil {
			utils.Handle(err)
		}
		if err := v.RegisterValidation("time", timeFormatValidation); err != nil {
			utils.Handle(err)
		}
		if err := v.RegisterValidation("warrantStatus", validStatus); err != nil {
			utils.Handle(err)
		}

	}

}
