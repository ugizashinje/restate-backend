package model

import (
	"github.com/lib/pq"
	"gopkg.in/guregu/null.v4"
)

type Apartment struct {
	BaseModel
	// address
	Street   string      `json:"street"`
	StreetNo null.String `json:"streetNo"`

	// Location
	City string `json:"city"`
	Muni string `json:"muni"`
	Hood string `json:"hood"`

	// Building
	Age               int  `json:"age"`
	TotalFloors       int  `json:"totalFloors"`
	Floor             int  `json:"floors"`
	Basement          bool `json:"basement"`
	IsGroundFloor     bool `json:"isGroundFloor"`
	IsHighGroundFloor bool `json:"isHighGroundFloor"`
	IsLastFloor       bool `json:"isLastFloor"`
	MonthlyFees       int  `json:"monthlyFees"`

	// Structure
	IsSallon    bool    `json:"isSallon"`
	IsDuplex    bool    `json:"isDuplex"`
	IsPenthouse bool    `json:"isPenthouse"`
	IsEnergyEff bool    `json:"isEnergyEff"`
	Condition   string  `json:"condition"`
	Heating     string  `json:"heating"`
	Rooms       float32 `json:"rooms"`

	// Legal
	IsRegistered    bool   `json:"isRegistered"`
	IsOccupiable    bool   `json:"isOccupiable"`
	IsWithinHouse   bool   `json:"isWithinHouse"`
	IsTaxRefundable bool   `json:"isTaxRefundable"`
	IsSwappable     bool   `json:"isSwappable"`
	IsMorgage       bool   `json:"isMorgage"`
	IsUrgent        bool   `json:"isUrgent"`
	Seller          string `json:"seller"` // agency , owner, investor
	// MISC
	Tags pq.StringArray `json:"tags" gorm:"type:text[]"`
}

func (Apartment) TableName() string {
	return "apartements"
}
