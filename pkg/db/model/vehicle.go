package model

import (
	"time"

	"gopkg.in/guregu/null.v4"
	"gorm.io/datatypes"
)

type Vehicle struct {
	BaseModel
	Plate                       string         `json:"plate"`
	IsTrailer                   bool           `json:"isTrailer"`
	Model                       string         `json:"model"`
	Producer                    string         `json:"producer"`
	Type                        string         `json:"type"`
	Weight                      null.Int       `json:"weight"`
	MaxWeight                   null.Int       `json:"maxWeight"`
	Capacity                    null.Int       `json:"capacity"`
	HomologationMark            string         `json:"homologationMark"`
	MotorVolume                 int            `json:"motorVolume"`
	Power                       int            `json:"power"`
	Fuel                        string         `json:"fuel"`
	Consumption                 float32        `json:"consumption"`
	CompanyID                   string         `json:"companyId"`
	Company                     Company        `json:"-"`
	ThrustWeight                null.Int       `json:"thrustWeight"`
	StandingSpace               null.Int       `json:"standingSpace"`
	SeatingSpace                null.Int       `json:"seatingSpace"`
	Color                       string         `json:"color"`
	Shaft                       int            `json:"shaft"`
	Chasie                      string         `json:"chasie"`
	Motor                       string         `json:"motor"`
	Insurance                   null.String    `json:"insurance"`
	RegistrationDate            time.Time      `json:"registrationDate" binding:"required,date"  format:"2006-01-02"`
	FirstRegistrationDate       null.Time      `json:"firstRegistrationDate" binding:"required,date" format:"2006-01-02"`
	OwnerMn                     int            `json:"ownerMn"`
	OwnerFirstName              string         `json:"ownerFirstName"`
	OwnerLastName               string         `json:"ownerLastName"`
	OwnerAddressID              string         `json:"ownerAddressId"`
	OwnerAddress                Address        `json:"ownerAddress"`
	UserFirstName               string         `json:"userFirstName"`
	UserLastName                string         `json:"userLastName"`
	UserAddressID               string         `json:"userAddressId"`
	UserAddress                 Address        `json:"userAddress"`
	UserMn                      int            `json:"userMn"`
	RegistrationCertificateDate time.Time      `json:"registrationCertificateDate" format:"2006-01-02"`
	RegistrationExpiry          time.Time      `json:"registrationExpiry" format:"2006-01-02"`
	RegisterId                  int            `json:"registerId"`
	Status                      string         `json:"status"`
	Meta                        datatypes.JSON `json:"meta"`
}

func (a Vehicle) ResouceName() string {
	return "vehicles"
}
func (a Vehicle) TableName() string {
	return "vehicles"
}
