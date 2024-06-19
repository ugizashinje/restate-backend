package model

import (
	"gopkg.in/guregu/null.v4"
)

type Address struct {
	BaseModel
	LocationID string `json:"locationId"`
	// Location Location    `json:"location"`
	Street   string      `json:"street"`
	StreetNo null.String `json:"streetNo"`
}

func (Address) TableName() string {
	return "addresses"
}
