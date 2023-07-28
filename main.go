package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kmerkuri/golang-crud-rest-api/controllers"
	"github.com/kmerkuri/golang-crud-rest-api/database"
	"log"
	"net/http"
)

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/api/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/api/products/{id}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/products/{id}", controllers.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/api/sales", controllers.GetSales).Methods("GET")
	router.HandleFunc("/api/sales/{id}", controllers.GetSalesID).Methods("GET")
	router.HandleFunc("/api/sales", controllers.CreateSale).Methods("POST")
	router.HandleFunc("/api/sales/{id}", controllers.UpdateSales).Methods("PUT")
	router.HandleFunc("/api/sales/{id}", controllers.DeleteSales).Methods("DELETE")
}
func main() {
	//    configs.LoadAppConfig()
	var constring string = "root:root@tcp(127.0.0.1:3306)/crud_demo"
	var conport int = 8080
	database.Connect(constring)
	database.Migrate()
	router := mux.NewRouter()
	// Register Routes
	RegisterProductRoutes(router)
	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", conport))
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}
