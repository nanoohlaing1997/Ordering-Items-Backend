package main

import (
	"fmt"
	"noh/go-order-items/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("hello from Ordering Items Project")
	godotenv.Load()
	dbName := os.Getenv("GO_ORDER")

	connection := dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	itemCategory := &models.ItemCategory{
		Name: "medicine",
	}
	category, _ := models.ItemCategoryCreate(db, itemCategory)

	item := &models.Item{
		Name:       "Sayar Kho",
		Price:      "1000",
		CategoryID: category.ID,
	}

	itemDetail, _ := models.ItemCreate(db, item)

	customerOrder := &models.CustomerOrder{
		CustomerId: 1,
		Attribute:  "items_id",
		Value:      itemDetail.ID,
	}

	order, _ := models.CustomerOrderCreate(db, customerOrder)
	fmt.Println("Done process.........................")
	fmt.Println(*order)
}
