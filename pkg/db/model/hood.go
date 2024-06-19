package model

import "gopkg.in/guregu/null.v4"

type Location struct {
	BaseModel
	City     string      `json:"city"`
	Muni     string      `json:"muni"`
	Location null.String `json:"location"`
}

func (Location) TableName() string {
	return "locations"
}
