package model

import (
	"time"
	"warrant-api/pkg/enum"

	"github.com/lib/pq"
	"gopkg.in/guregu/null.v4"
)

type Warrant struct {
	BaseModel
	IssueDate            time.Time          `json:"issueDate" format:"2006-01-02"`
	ExpectedStart        time.Time          `json:"expectedStart"`
	ClosingDate          time.Time          `json:"closingDate" format:"2006-01-02"`
	DriverID             string             `json:"driverId"`
	Driver               User               `json:"-"`
	Passengers           pq.StringArray     `json:"passengers" gorm:"type:text[]"`
	VehicleID            string             `json:"vehicleId"`
	TrailerID            null.String        `json:"trailerId"`
	CompanyID            string             `json:"companyId"`
	Company              Company            `json:"-"`
	DispatcherID         string             `json:"dispatcherId"`
	Dispatcher           User               `json:"-"`
	TechnicalCorrectness string             `json:"technicalCorrectness"`
	Status               enum.WarrantStatus `json:"status"`
	Name                 string             `json:"name"`
	Note                 null.String        `json:"note,omitempty"`
}

func (a Warrant) TableName() string {
	return "warrants"
}
