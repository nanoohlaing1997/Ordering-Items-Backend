package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID             uint64    `json:"id"`
	InvoiceID      string    `json:"invoice_id"`
	UserID         uint64    `json:"user_id"`
	ItemID         uint64    `json:"item_id"`
	Quantity       int64     `json:"quantity"`
	DeliveryStatus int32     `json:"delivery_status"`
	CreatedAt      time.Time `gorm:"Column:created_at"`
	UpdatedAt      time.Time `gorm:"Column:Updated_at"`
}

type OrderDB struct {
	db    *gorm.DB
	model *Order
}

func OrderManager(db *gorm.DB) *OrderDB {
	return &OrderDB{
		db:    db,
		model: &Order{},
	}
}

func (om *OrderDB) CustomerOrderCreate(order *Order) (*Order, error) {
	if res := om.db.Create(order); res.RowsAffected <= 0 {
		return nil, res.Error
	}
	return order, nil
}

func (om *OrderDB) GetCustomerOrder() (*Order, error) {
	var order *Order
	if res := om.db.First(&order); res.Error != nil {
		return nil, res.Error
	}
	return order, nil
}
