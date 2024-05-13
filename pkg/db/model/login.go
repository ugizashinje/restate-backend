package model

import (
	"warrant-api/pkg/enum"
)

type Login struct {
	BaseModel
	Email        string           `json:"email"`
	UserAgent    string           `json:"device"`
	AccessToken  string           `json:"access_token"`
	RefreshToken string           `json:"refresh_token"`
	Result       enum.LoginResult `json:"result"`
}

func (a Login) TableName() string {
	return "login"
}
