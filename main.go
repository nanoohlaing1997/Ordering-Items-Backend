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
	authRouter.HandleFunc("/healthz", controller.HealthzHandler).Methods("GET")

	// Testing healthz route
	router.HandleFunc("/healthz", controller.HealthzHandler).Methods("GET")
	return router
}
