package model

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type UserCollection struct {
	IsDeleted    bool      `json:"isDeleted" gorm:"default:false"`
	IsActive     bool      `json:"isActive" gorm:"default:true"`
	CreatedAt    time.Time `json:"-" gorm:"default:now()"`
	UpdatedAt    null.Time `json:"-" gorm:"default:null"`
	DeletedAt    null.Time `json:"-" gorm:"default:null"`
	Role         string    `json:"role"`
	UserID       string    `json:"-" gorm:"primaryKey"`
	CollectionID string    `json:"collectionId" gorm:"primaryKey"`
}

func (uc UserCollection) TableName() string {
	return "user_collections"
}
