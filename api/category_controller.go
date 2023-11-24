package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nanoohlaing1997/online-ordering-items/models"
)

type CreateCategoryRequest struct {
	CategoryName string `json:"category_name" validate:"required"`
}

type DeteleCategoryRequest struct {
	CategoryID string `json:"category_id" validate:"required"`
}

func (c *Controller) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var categoryRequest CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&categoryRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Use Validator to perform struct validation
	if err := validate.Struct(categoryRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	category := c.dbm.CreateItemCategory(&models.Category{
		Name: categoryRequest.CategoryName,
	})

	json.NewEncoder(w).Encode(category)
}

func (c *Controller) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := c.dbm.GetCategories()
	if err != nil {
		http.Error(w, "Item Category not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(categories)
}

func (c *Controller) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	// user, ok := r.Context().Value(authUser).(*models.User)
	// if !ok {
	// 	http.Error(w, "Fail to get user from context", http.StatusInternalServerError)
	// 	return
	// }

	var categoryRequest DeteleCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&categoryRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(categoryRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	categoryID, _ := StringToUint64(categoryRequest.CategoryID)
	res, _ := c.dbm.DeleteCategory(categoryID)
	if !res {
		http.Error(w, "Removing category process failed!!!", http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Removing category process succeeded!!!")
}
