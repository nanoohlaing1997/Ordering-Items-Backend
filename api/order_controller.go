package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nanoohlaing1997/online-ordering-items/models"
)

type CreateOrderRequest struct {
	UserID   uint64 `json:"user_id" validate:"required"`
	ItemID   uint64 `json:"item_id" validate:"required"`
	Quantity int32  `json:"quantity" validate:"required"`
}

type GetOrdersRequest struct {
	UserID uint64 `json:"user_id" validate:"required"`
}

type GetOrderRequest struct {
	InvoiceID string `json:"invoice" validate:"required"`
}

type UpdateDeliveryRequest struct {
	InvoiceID      string `json:"invoice" validate:"required"`
	DeliveryStatus string `json:"delivery_status" validate:"required"`
}

func (c *Controller) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var orderRequest CreateOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&orderRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Use Validator to perform struct validation
	if err := validate.Struct(orderRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := c.dbm.GetUserByID(orderRequest.UserID)
	if err != nil {
		http.Error(w, "User not found!!!", http.StatusNotFound)
		return
	}

	invoiceID := GenerateInvoiceID()

	order := &models.Order{
		InvoiceID:      invoiceID,
		UserID:         user.ID,
		ItemID:         orderRequest.ItemID,
		Quantity:       int64(orderRequest.Quantity),
		DeliveryStatus: 0,
	}

	orderObj, err := c.dbm.CreateOrder(order)
	if err != nil {
		http.Error(w, "Creating order process failed!!!", http.StatusNotAcceptable)
		return
	}

	json.NewEncoder(w).Encode(orderObj)
}

func (c *Controller) GetOrderByUserID(w http.ResponseWriter, r *http.Request) {
	var orderRequest GetOrdersRequest

	if err := json.NewDecoder(r.Body).Decode(&orderRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Use Validator to perform struct validation
	if err := validate.Struct(orderRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	orders, err := c.dbm.GetOrdersByUserID(orderRequest.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	if len(orders) <= 0 {
		http.Error(w, "Invoice not found!!!", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

func (c *Controller) GetOrderByInvoiceID(w http.ResponseWriter, r *http.Request) {
	var orderRequest GetOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&orderRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Use Validator to perform struct validation
	if err := validate.Struct(orderRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order, err := c.dbm.GetOrderByInvoice(orderRequest.InvoiceID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	json.NewEncoder(w).Encode(order)
}

func (c *Controller) UpdateDeliveryStatus(w http.ResponseWriter, r *http.Request) {
	var deliveryRequest UpdateDeliveryRequest

	if err := json.NewDecoder(r.Body).Decode(&deliveryRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Use Validator to perform struct validation
	if err := validate.Struct(deliveryRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	status := checkDeliveryStatus(deliveryRequest.DeliveryStatus)

	_, err := c.dbm.UpdateOrder(deliveryRequest.InvoiceID, int32(status))
	if err != nil {
		http.Error(w, "Update delivery status failed", http.StatusNotAcceptable)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Updateing devlivery status succeeded!!!")
}

func checkDeliveryStatus(status string) models.DeliveryStatus {
	statusMap := map[string]models.DeliveryStatus{
		"Pending":       models.Pending,
		"Processing":    models.Processing,
		"Delivery done": models.Done,
		"Cancel":        models.Cancel,
	}

	if val, ok := statusMap[status]; ok {
		return val
	}

	return models.Default
}
