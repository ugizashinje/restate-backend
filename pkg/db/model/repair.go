package model

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type Repair struct {
	BaseModel
	WarrantID  string      `json:"warrantId"`
	Warrant    Warrant     `json:"-"`
	Workshop   string      `json:"workshop"`
	RepairType string      `json:"repairType"`
	Location   null.String `json:"location"`
	AddressID  null.String `json:"addressId"`
	Address    Address     `json:"address"`
	Start      time.Time   `json:"start"`
	End        null.Time   `json:"end"`
	FileName   null.String `json:"fileName"`
	Url        string      `json:"url,omitempty" gorm:"-"`
}

func (c Repair) ResouceName() string {
	return "repairs"
}

func (Repair) TableName() string {
	return "repairs"
}
