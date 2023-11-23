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
	Items     []Item    `gorm:"foreignkey:CategoryID"`
}

type CategoryDB struct {
	db    *gorm.DB
	model *Category
}

func (c *Category) TableName() string {
	return "item_categories"
}

// Model relationship
func (c *Category) GetItems() []Item {
	return c.Items
}

func CategoryManager(db *gorm.DB) *CategoryDB {
	return &CategoryDB{
		db:    db,
		model: &Category{},
	}
}

func (cdb *CategoryDB) CreateItemCategory(category *Category) *Category {
	cdb.db.Where(Category{Name: category.Name}).FirstOrCreate(&category)
	return category
}

func (cdb *CategoryDB) GetCategories() ([]*Category, error) {
	ca := make([]*Category, 0)
	if res := cdb.db.Find(&ca); res.Error != nil {
		return nil, res.Error
	}
	return ca, nil
}

func (cdb *CategoryDB) DeleteCategory(categoryID uint64) (bool, error) {
	if res := cdb.db.Delete(&Category{ID: categoryID}); res.Error != nil {
		return false, res.Error
	}
	return true, nil
}

func (cdb *CategoryDB) GetCategoryById(categoryID uint64) (*Category, error) {
	var category Category
	if res := cdb.db.First(&category, "id = ?", categoryID); res.Error != nil {
		return nil, res.Error
	}
	return &category, nil
}
