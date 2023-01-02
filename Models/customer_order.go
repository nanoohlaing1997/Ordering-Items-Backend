package Models

import "time"

type CustomerOrder struct {
	CustomerId uint64    `json:"customer_id"`
	Attribute  string    `json:"attribute"`
	Value      string    `josn:"value"`
	CreatedAt  time.Time `gorm:"Column:created_at"`
	UpdatedAt  time.Time `gorm:"Column:Updated_at"`
}
