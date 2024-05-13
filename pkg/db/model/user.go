package model

import (
	"gopkg.in/guregu/null.v4"
)

type User struct {
	BaseModel
	Email               string        `json:"email" gorm:"unique"`
	Password            string        `json:"-"`
	FirstName           string        `json:"firstName"`
	LastName            string        `json:"lastName"`
	Mn                  string        `json:"mn"`
	Birthplace          string        `json:"birthplace"`
	Department          null.String   `json:"department"`
	Companies           []UserCompany `json:"companies,omitempty"`
	Phone               string        `json:"phone"`
	AddressID           string        `json:"addressId"`
	Address             Address       `json:"address"`
	IsDriver            bool          `json:"isDriver"`
	LicenceNumber       null.String   `json:"licenceNumber"`
	LicenceSerialNumber null.String   `json:"licenceSerialNumber"`
	LicenceExpiry       null.Time     `json:"licenceExpiry" format:"2006-01-02"`
	LicenceIssued       null.Time     `json:"licenceIssued" format:"2006-01-02"`
	LicenceAuthority    null.String   `json:"licenceAuthority"`
	Status              string        `json:"status"`
}

func (a User) TableName() string {
	return "users"
}
