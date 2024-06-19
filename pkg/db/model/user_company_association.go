package model

// import (
// 	"time"

// 	"github.com/lib/pq"
// 	"gopkg.in/guregu/null.v4"
// )

// type UserCompany struct {
// 	IsDeleted bool           `json:"isDeleted" gorm:"default:false"`
// 	IsActive  bool           `json:"isActive" gorm:"default:true"`
// 	CreatedAt time.Time      `json:"-" gorm:"default:now()"`
// 	UpdatedAt null.Time      `json:"-" gorm:"default:null"`
// 	DeletedAt null.Time      `json:"-" gorm:"default:null"`
// 	Roles     pq.StringArray `json:"roles" gorm:"type:text[]"`
// 	UserID    string         `json:"-" gorm:"primaryKey"`
// 	CompanyID string         `json:"companyId" gorm:"primaryKey"`
// }

// func (uc UserCompany) TableName() string {
// 	return "user_companies"
// }
