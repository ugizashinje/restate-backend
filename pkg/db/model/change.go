package model

import (
	"gorm.io/datatypes"
)

type Change struct {
	BaseModel
	SubjectID string         `json:"subjectId"`
	UserID    string         `json:"userId"`
	CompanyID string         `json:"companyId"`
	Table     string         `json:"table"`
	Fields    datatypes.JSON `json:"fields"`
}

func (a Change) ResouceName() string {
	return "changes"
}

func (Change) TableName() string {
	return "changes"
}
