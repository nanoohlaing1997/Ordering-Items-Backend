package models

import (
	"time"

	"gorm.io/gorm"
)

type CustomerOrder struct {
	CustomerId uint64    `json:"customer_id"`
	Attribute  string    `json:"attribute"`
	Value      uint64    `json:"value"`
	CreatedAt  time.Time `gorm:"Column:created_at"`
	UpdatedAt  time.Time `gorm:"Column:Updated_at"`
}

type CustomerOrderDB struct {
	db *gorm.DB
}

func CustomerOrderManager(db *gorm.DB) *CustomerOrderDB {
	return &CustomerOrderDB{
		db: db,
	}
}

func CustomerOrderCreate(db *gorm.DB, customerOrder *CustomerOrder) (*CustomerOrder, error) {
	if res := db.Create(customerOrder); res.RowsAffected <= 0 {
		return nil, res.Error
	}
	return customerOrder, nil
}

func (cdb *CustomerOrderDB) GetCustomerOrder() (*CustomerOrder, error) {
	var customerOrder *CustomerOrder
	if res := cdb.db.First(&customerOrder); res.Error != nil {
		return nil, res.Error
	}
	return customerOrder, nil
}
