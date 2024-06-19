package model

type Collection struct {
	BaseModel
	Name    string `json:"name"`
	OwnerID string `json:"ownerId"`
	Owner   User   `json:"owner"`
}

func (Collection) TableName() string {
	return "collections"
}
