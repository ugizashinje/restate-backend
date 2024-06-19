package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/lib/pq"
	"gopkg.in/guregu/null.v4"
)

type Apartment struct {
	BaseModel
	Ad
	// address
	Name     string      `json:"name"`
	Street   null.String `json:"street"`
	StreetNo null.String `json:"streetNo"`

	// Location
	City     string      `json:"city"`
	Muni     null.String `json:"muni"`
	Location null.String `json:"location"`

	// Building
	Age             null.Int `json:"age"`
	TotalFloors     int      `json:"totalFloors"`
	Floor           int      `json:"floors"`
	Basement        bool     `json:"basement"`
	GroundFloor     bool     `json:"groundFloor"`
	HighGroundFloor bool     `json:"highGroundFloor"`
	LastFloor       bool     `json:"lastFloor"`
	MonthlyFees     int      `json:"monthlyFees"`

	// Structure
	Sallon    bool   `json:"sallon"`
	Duplex    bool   `json:"duplex"`
	Penthouse bool   `json:"penthouse"`
	EnergyEff bool   `json:"energyEff"`
	Condition string `json:"condition"`
	Heating   string `json:"heating"`
	Rooms     string `json:"rooms"`

	// Legal
	Registered    bool   `json:"registered"`
	Occupiable    bool   `json:"occupiable"`
	WithinHouse   bool   `json:"withinHouse"`
	TaxRefundable bool   `json:"taxRefundable"`
	Swappable     bool   `json:"swappable"`
	Morgage       bool   `json:"morgage"`
	Urgent        bool   `json:"urgent"`
	Seller        string `json:"seller"` // agency , owner, investor
	// MISC
	Tags         pq.StringArray `json:"tags" gorm:"type:text[]"`
	Thumbnail    string         `json:"thumbnail"`
	FloorPlan    FloorPlan      `json:"floorPlan"`
	AllRoundTour AllRoundTour   `json:"allRoundTour"`
	VirtualTour  null.String    `json:"virutalTour"`
}

type FloorPlan []Floor

type Floor struct {
	Plan string `json:"plan"`
}

func (a *FloorPlan) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to unmarshal FloorPlan value")
	}

	return json.Unmarshal(bytes, a)
}

// Implement the valuer interface
func (a FloorPlan) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *AllRoundTour) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to unmarshal FloorPlan value")
	}

	return json.Unmarshal(bytes, a)
}

// Implement the valuer interface
func (a AllRoundTour) Value() (driver.Value, error) {
	return json.Marshal(a)
}

type AllRoundTour []AllRoundFloor
type AllRoundFloor struct {
	Plan string `json:"plan"`
}

func (Apartment) TableName() string {
	return "apartments"
}
