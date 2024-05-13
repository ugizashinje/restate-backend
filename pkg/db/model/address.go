package model

import (
	"gopkg.in/guregu/null.v4"
)

type Address struct {
	BaseModel
	City     string      `json:"city"`
	Street   string      `json:"street"`
	StreetNo null.String `json:"streetNo"`
	Name     null.String `json:"name"`
}

func (Address) TableName() string {
	return "addresses"
}
