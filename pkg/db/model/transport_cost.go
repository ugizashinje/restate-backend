package model

import "gopkg.in/guregu/null.v4"

type TransportCost struct {
	BaseModel
	WarrantID string      `json:"warrantID"`
	Warrant   Warrant     `json:"-"`
	Type      null.String `json:"type"` // type of cost, fuel, parking, lubricant .... ( za sad putarina i gorivo )
	Amount    null.Float  `json:"amount"`
	Location  string      `json:"location"`
	Code      null.String `json:"code"` // serial code of fiscal bill
	FileName  null.String `json:"fileName"`
	Url       string      `json:"url,omitempty" gorm:"-"`
}

func (a TransportCost) ResouceName() string {
	return "transportCosts"
}
func (a TransportCost) TableName() string {
	return "transport_costs"
}
