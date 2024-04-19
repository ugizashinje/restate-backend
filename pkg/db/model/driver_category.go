package model

import (
	"gopkg.in/guregu/null.v4"
)

type DriverCategory struct {
	BaseModel
	UserId   string    `json:"userId"`
	Category string    `json:"category"`
	Issued   null.Time `json:"issued"  format:"2006-01-02"`
	Expired  null.Time `json:"expired"  format:"2006-01-02"`
}

func (a DriverCategory) ResouceName() string {
	return "driverCategories"
}

func (DriverCategory) TableName() string {
	return "driver_categories"
}
