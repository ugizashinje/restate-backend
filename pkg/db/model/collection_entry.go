package model

type CollectionEntry struct {
	BaseModel
	Description  string     `json:"description"`
	CollectionID string     `json:"ownerId"`
	Collection   Collection `json:"-"`
}

func (CollectionEntry) TableName() string {
	return "collection_entries"
}
