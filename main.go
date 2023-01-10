package main

import (
	"fmt"
	"log"
	"net/http"
	"noh/go-order-items/controller"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("hello from Ordering Items Project")

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	RegisterProductRoutes(router)

	// Start the server
	godotenv.Load()
	log.Println(fmt.Sprintf("Starting Server on port %s", os.Getenv("PORT")))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("PORT")), router))
}

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/api/customer/order", controller.GetOrder).Methods("GET")
}
