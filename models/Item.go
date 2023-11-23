package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID         uint64    `json:"id"`
	Name       string    `json:"name"`
	Price      float64   `json:"price"`
	Quantity   int64     `json:"quantity"`
	CategoryID uint64    `json:"category_id"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

type ItemDB struct {
	db    *gorm.DB
	model *Item
}

func (item *Item) TableName() string {
	return "items"
}

func ItemManager(db *gorm.DB) *ItemDB {
	return &ItemDB{
		db:    db,
		model: &Item{},
	}
}

func (idb *ItemDB) CreateItem(item *Item) (*Item, error) {
	if res := idb.db.Create(&item); res.RowsAffected <= 0 {
		return nil, res.Error
	}
	return item, nil
}

func (idb *ItemDB) GetItemsByCategoryID(categoryID uint64) (*Category, error) {
	var ca Category
	if res := idb.db.Preload("Items").First(&ca, categoryID); res.Error != nil {
		return nil, res.Error
	}
	return &ca, nil
}
