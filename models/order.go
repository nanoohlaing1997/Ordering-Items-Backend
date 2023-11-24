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

func (order *Order) TableName() string {
	return "orders"
}

func OrderManager(db *gorm.DB) *OrderDB {
	return &OrderDB{
		db:    db,
		model: &Order{},
	}
}

type DeliveryStatus int32

const (
	Default DeliveryStatus = iota - 1
	Pending
	Processing
	Done
	Cancel
)

func (odb *OrderDB) CreateOrder(order *Order) (*Order, error) {
	if res := odb.db.Create(order); res.RowsAffected <= 0 {
		return nil, res.Error
	}
	return order, nil
}

func (odb *OrderDB) GetOrdersByUserID(userID uint64) ([]*Order, error) {
	orders := make([]*Order, 0)
	if res := odb.db.Where("user_id = ?", userID).Find(&orders); res.Error != nil {
		return nil, res.Error
	}
	return orders, nil
}

func (odb *OrderDB) GetOrderByInvoice(invoice string) (*Order, error) {
	var order Order
	if res := odb.db.Where("invoice_id = ?", invoice).First(&order); res.Error != nil {
		return nil, res.Error
	}
	return &order, nil
}

func (odb *OrderDB) UpdateOrder(invoice string, status int32) (bool, error) {
	var orderToUpdate Order
	if err := odb.db.First(&orderToUpdate, "invoice_id = ?", invoice).Error; err != nil {
		return false, err
	}
	orderToUpdate.DeliveryStatus = status

	if err := odb.db.Save(&orderToUpdate).Error; err != nil {
		return false, err
	}
	return true, nil
}
