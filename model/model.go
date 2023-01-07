package model

type User struct {
	ID     uint     `gorm:"primaryKey"`
	Title  string   `json:"title"`
	Amount int      `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags" gorm:"serializer:json"`
}
