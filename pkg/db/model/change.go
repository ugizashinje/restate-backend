package model

import (
	"warrant-api/pkg/enum"

	"gorm.io/datatypes"
)

type Change struct {
	BaseModel
	SubjectID string         `json:"subjectId"`
	UserID    string         `json:"userId"`
	Ad        string         `json:"ad"`
	Fields    datatypes.JSON `json:"fields"`
	Event     enum.EventType `json:"event"`
	Meta      datatypes.JSON `json:"meta"`
}

func (Change) TableName() string {
	return "changes"
}
