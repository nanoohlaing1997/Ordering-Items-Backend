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

func CustomerOrderCreate(db *gorm.DB, customerOrder *CustomerOrder) (*CustomerOrder, error) {
	if res := db.Create(customerOrder); res.RowsAffected <= 0 {
		return nil, res.Error
	}
	return customerOrder, nil
}
