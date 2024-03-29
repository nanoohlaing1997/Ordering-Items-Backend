package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/nanoohlaing1997/online-ordering-items/api"
	"github.com/nanoohlaing1997/online-ordering-items/env"
	"github.com/nanoohlaing1997/online-ordering-items/log"
	"github.com/nanoohlaing1997/online-ordering-items/models"
)

var environ = env.GetEnviroment()

func main() {
	fmt.Println("hello from Ordering Items Project")
	log := log.GetLogger()

	// Register Routes
	router := RegisterRoute()

	// Start the server
	godotenv.Load()
	// log.Println(fmt.Sprintf("Starting Server on port %s", environ.RestPort))
	log.Printf("Starting Server on port %s", environ.RestPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", environ.RestPort), router))
}

func RegisterRoute() *mux.Router {
	dbm := models.NewDatabaseManager()
	controller := api.NewControllerManager(dbm)

	router := mux.NewRouter()

	// Register API route
	// Public routes
	publicRouter := router.PathPrefix("/order").Subrouter()
	publicRouter.HandleFunc("/signup", controller.SignUpHandler).Methods("POST")
	publicRouter.HandleFunc("/signin", controller.SignInHandler).Methods("GET")
	publicRouter.HandleFunc("/token/refresh", controller.TokenRefreshHandler).Methods("GET")

	// Authentated routes
	authRouter := router.PathPrefix("/order/auth").Subrouter()
	authRouter.Use(api.AuthMiddleWare)
	authRouter.HandleFunc("/categories", controller.GetCategories).Methods("GET")
	authRouter.HandleFunc("/category/{category_id}/items", controller.GetItems).Methods("GET")
	authRouter.HandleFunc("/invoice", controller.CreateOrder).Methods("POST")
	authRouter.HandleFunc("/invoice/user", controller.GetOrderByUserID).Methods("GET")
	authRouter.HandleFunc("/invoice", controller.GetOrderByInvoiceID).Methods("GET")

	// Admin routes
	adminRouter := router.PathPrefix("/order/admin").Subrouter()
	adminRouter.Use(api.AuthMiddleWare)
	adminRouter.Use(api.AdminMiddleware)
	adminRouter.HandleFunc("/category/{user_id}", controller.CreateCategory).Methods("POST")
	adminRouter.HandleFunc("/category/{user_id}/remove", controller.DeleteCategory).Methods("DELETE")
	adminRouter.HandleFunc("/item/{user_id}/create", controller.CreateItem).Methods("POST")
	adminRouter.HandleFunc("/invoice/{user_id}/delivery", controller.UpdateDeliveryStatus).Methods("PUT")

	// Testing healthz route
	router.HandleFunc("/healthz", controller.HealthzHandler).Methods("GET")
	return router
}
