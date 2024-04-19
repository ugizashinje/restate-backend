package model

import (
	"gopkg.in/guregu/null.v4"
)

type Alias struct {
	BaseModel
	AddressID string      `json:"addressId"`
	Address   Address     `json:"address"`
	CompanyID string      `json:"companyId"`
	Company   Company     `json:"company"`
	Name      null.String `json:"name"`
}

func (a Alias) ResouceName() string {
	return "aliases"
}

func (a Alias) TableName() string {
	return "aliases"
}
