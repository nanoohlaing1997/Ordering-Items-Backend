package models

import (
	"time"

	"gorm.io/gorm"
)

type ItemCategory struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func ItemCategoryCreate(db *gorm.DB, category *ItemCategory) (*ItemCategory, error) {
	if res := db.Create(&category); res.RowsAffected <= 0 {
		return nil, res.Error
	}
	return category, nil
}
