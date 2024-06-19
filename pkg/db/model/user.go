package model

import "gopkg.in/guregu/null.v4"

type User struct {
	BaseModel
	Email       string           `json:"email" gorm:"unique"`
	Avatar      null.String      `json:"avatar"`
	Password    string           `json:"-"`
	FirstName   string           `json:"firstName"`
	LastName    string           `json:"lastName"`
	Mn          null.String      `json:"mn"`
	Birthplace  null.String      `json:"birthplace"`
	Phone       string           `json:"phone"`
	AddressID   string           `json:"addressId"`
	Address     Address          `json:"address"`
	CompanyID   null.String      `json:"companyId,omitempty"`
	Company     Company          `json:"-,omitempty"`
	Role        string           `json:"role"`
	Collections []UserCollection `json:"-"`
	Likes       []Like           `json:"-"`
	Status      string           `json:"status"`
}

func (a User) TableName() string {
	return "users"
}
