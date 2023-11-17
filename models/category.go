package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type CategoryDB struct {
	db    *gorm.DB
	model *Category
}

func CategoryManager(db *gorm.DB) *CategoryDB {
	return &CategoryDB{
		db:    db,
		model: &Category{},
	}
}

func (icm *CategoryDB) ItemCategoryCreate(category *Category) (*Category, error) {
	if res := icm.db.Create(&category); res.RowsAffected <= 0 {
		return nil, res.Error
	}
	return category, nil
}
