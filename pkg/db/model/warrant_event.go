package model

import (
	"warrant-api/pkg/enum"

	"gopkg.in/guregu/null.v4"
	"gorm.io/datatypes"
)

type WarrantEvent struct {
	BaseModel
	WarrantID string         `json:"warrantId"`
	Warrant   Warrant        `json:"-"`
	RouteID   null.String    `json:"routeId,omitempty"`
	Route     *Route         `json:"-"`
	Event     enum.EventType `json:"event"`
	UserID    string         `json:"userId"`
	User      User           `json:"-"`
	Meta      datatypes.JSON `json:"meta"`
}

func (a WarrantEvent) ResouceName() string {
	return "warrantEvents"
}
func (a WarrantEvent) TableName() string {
	return "warrant_events"
}
