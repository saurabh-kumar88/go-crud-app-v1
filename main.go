package main

import (
	"fmt"
	database "golang-crud-rest-api/databases"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/api/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/api/products/{id}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/products/{id}", controllers.DeleteProduct).Methods("DELETE")
}

func main() {
	// Load Configuration from config.json using Viper
	LoadAppConfig()

	// Initialize the Database
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register routers
	RegisterProductRoutes(router)

	// start ther server.

	log.Println(fmt.Sprintf("Starting server on port %vs", AppConfig.Port))
	log.Println(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))

}
