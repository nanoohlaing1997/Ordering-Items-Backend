package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID         uint64    `json:"id"`
	Name       string    `json:"name"`
	Price      string    `json:"price"`
	CategoryID uint64    `gorm:"column:category_id"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

type ItemDB struct {
	db *gorm.DB
}

func ItemManager(db *gorm.DB) *ItemDB {
	return &ItemDB{
		db: db,
	}
}

func ItemCreate(db *gorm.DB, item *Item) (*Item, error) {
	if res := db.Create(&item); res.RowsAffected <= 0 {
		return nil, res.Error
	}
	return item, nil
}
