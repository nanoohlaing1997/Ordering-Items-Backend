package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nanoohlaing1997/online-ordering-items/models"
)

type CreateItemRequest struct {
	Name       string  `json:"name" validate:"required"`
	Price      float64 `json:"price" validate:"required"`
	Quantity   int64   `json:"quantity" validate:"required"`
	CategoryID uint64  `json:"category_id" validate:"required"`
}

func (c *Controller) CreateItem(w http.ResponseWriter, r *http.Request) {
	var itemRequest CreateItemRequest
	if err := json.NewDecoder(r.Body).Decode(&itemRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Use Validator to perform struct validation
	if err := validate.Struct(itemRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	category, err := c.dbm.GetCategoryById(itemRequest.CategoryID)
	if err != nil {
		http.Error(w, "Item Category not found!!!", http.StatusInternalServerError)
		return
	}

	item := &models.Item{
		Name:       itemRequest.Name,
		Price:      itemRequest.Price,
		Quantity:   itemRequest.Quantity,
		CategoryID: category.ID,
	}

	itemObj, err := c.dbm.CreateItem(item)
	if err != nil {
		http.Error(w, "Item creation process failed!!!", http.StatusNotAcceptable)
		return
	}

	json.NewEncoder(w).Encode(itemObj)
}

func (c *Controller) GetItems(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stringCategoryID := vars["category_id"]
	categoryID, _ := StringToUint64(stringCategoryID)
	if categoryID <= 0 {
		http.Error(w, "Category Id is not valid", http.StatusInternalServerError)
		return
	}

	categoryWithItems, err := c.dbm.GetItemsByCategoryID(categoryID)
	if err != nil {
		http.Error(w, "Category not found!!!", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(categoryWithItems)
}
