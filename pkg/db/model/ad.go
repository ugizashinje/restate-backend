package model

import (
	"gopkg.in/guregu/null.v4"
)

type Ad struct {
	BaseModel
	City     string      `json:"city"`
	Street   string      `json:"street"`
	StreetNo null.String `json:"streetNo"`
	Name     null.String `json:"name"`
}

func (Ad) TableName() string {
	return "ad"
}
