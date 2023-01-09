package main

import (
	"fmt"
	"log"
	"net/http"
	"noh/go-order-items/cmd"
	"noh/go-order-items/controller"
	"noh/go-order-items/database"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("hello from Ordering Items Project")
	cmd.Execute()
	godotenv.Load()
	// dbName := os.Getenv("GO_ORDER")

	// connection := dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// }

	// itemCategory := &models.ItemCategory{
	// 	Name: "medicine",
	// }
	// category, _ := models.ItemCategoryCreate(db, itemCategory)

	// item := &models.Item{
	// 	Name:       "Sayar Kho",
	// 	Price:      "1000",
	// 	CategoryID: category.ID,
	// }

	// itemDetail, _ := models.ItemCreate(db, item)

	// customerOrder := &models.CustomerOrder{
	// 	CustomerId: 1,
	// 	Attribute:  "items_id",
	// 	Value:      itemDetail.ID,
	// }

	// order, _ := models.CustomerOrderCreate(db, customerOrder)
	// fmt.Println("Done process.........................")
	// fmt.Println(*order)

	database.DBManger()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	RegisterProductRoutes(router)

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", os.Getenv("PORT")))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("PORT")), router))
}

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/api/customer/order", controller.GetOrder).Methods("GET")
}
