package Models

import "time"

type Item struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Price     string    `json:"price"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updateed_at"`
}
