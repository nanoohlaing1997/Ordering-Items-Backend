package controller

import (
	"encoding/json"
	"net/http"
	"noh/go-order-items/database"
)

func GetOrder(w http.ResponseWriter, r *http.Request) {
	db := database.DBManger().CustomerOrderDB
	res, _ := db.GetCustomerOrder()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
