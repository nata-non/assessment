package model

import "github.com/lib/pq"

type Expenses struct {
	ID     uint           `gorm:"primary key;autoIncrement" json:"id"`
	Title  string         `json:"title"`
	Amount int            `json:"amount"`
	Note   string         `json:"note"`
	Tags   pq.StringArray `json:"tags" gorm:"type:text[]"`
}
