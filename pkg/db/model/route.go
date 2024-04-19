package model

import (
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type Route struct {
	BaseModel
	WarrantID      string    `json:"warrantId"`
	Warrant        Warrant   `json:"-"`
	StartMileage   null.Int  `json:"startMileage"`
	StartAddressID string    `gorm:"not null" json:"startAddressId"`
	StartAddress   Address   `json:"startAddress"`
	StartTime      null.Time `json:"startTime"`
	EndMileage     null.Int  `json:"endMileage"`
	EndAddressID   string    `gorm:"not null" json:"endAddressId"`
	EndAddress     Address   `json:"endAddress"`
	EndTime        null.Time `json:"endTime"`
	Order          int       `json:"order"`
	Status         string    `json:"status" grom:"not null, default:ready"`
}

func (c Route) ResouceName() string {
	return "routes"
}
func (c Route) TableName() string {
	return "routes"
}

func (r *Route) AfterCreate(db *gorm.DB) error {
	r.StartAddress = Address{BaseModel: BaseModel{ID: r.StartAddressID}}
	r.EndAddress = Address{BaseModel: BaseModel{ID: r.EndAddressID}}

	if res := db.Find(&r.StartAddress); res.Error != nil {
		return res.Error
	}
	if res := db.Find(&r.EndAddress); res.Error != nil {
		return res.Error
	}
	return nil
}
