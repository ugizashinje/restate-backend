package model

type Like struct {
	BaseModel
	Name   string `json:"name"`
	UserID string `json:"userId"`
	User   User   `json:"user"`
	AdID   string `json:"adId"`
	Ad     Ad     `json:"-"`
}

func (Like) TableName() string {
	return "likes"
}
