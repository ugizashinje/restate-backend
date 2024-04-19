package model

import (
	"gorm.io/datatypes"
)

type Company struct {
	BaseModel
	Name      string         `json:"name"`
	Short     string         `json:"short"`
	PIB       string         `json:"pib" gorm:"unique"`
	AddressID string         `json:"addressId"`
	Address   Address        `json:"address"`
	Mn        string         `json:"mn" gorm:"unique"`
	Phone     string         `json:"phone"`
	Email     string         `json:"email"`
	Employees []UserCompany  `json:"companies"`
	Meta      datatypes.JSON `json:"meta"`
}

func (a Company) ResouceName() string {
	return "companies"
}

func (Company) TableName() string {
	return "companies"
}
