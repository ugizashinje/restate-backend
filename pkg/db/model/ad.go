package model

import (
	"gopkg.in/guregu/null.v4"
)

type Ad struct {
	BaseModel
	CompanyID      null.String `json:"companyId"`
	Company        Company     `json:"-"`
	UserID         string      `json:"userId"` // agent may be working for other company
	User           User        `json:"-"`
	Market         string      `json:"market"`
	Status         string      `json:"status"`
	Published      null.Time   `json:"published"`
	ThumbnailImage string      `json:"tumbnailImage"`
	ThumbnailText  string      `json:"tumbnailText"`
	Description    string      `json:"description"`
}

func (Ad) TableName() string {
	return "ads"
}
