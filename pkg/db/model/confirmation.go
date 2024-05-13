package model

import (
	"warrant-api/pkg/enum"
)

type Confirmation struct {
	BaseModel
	Url    string                  `json:"-"`
	Code   string                  `json:"code"`
	UserID string                  `json:"-"`
	User   User                    `json:"-"`
	Status enum.ConfirmationStatus `json:"status"`
}

func (Confirmation) TableName() string {
	return "confirmations"
}
