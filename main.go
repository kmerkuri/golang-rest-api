package main

import (
	"github.com/gorilla/mux"
	"github.com/kmerkuri/golang-crud-rest-api/database"
	"github.com/kmerkuri/golang-crud-rest-api/controllers"
	"log"
	"fmt"
	"net/http"
)

func RegisterProductRoutes(router *mux.Router){
	router.HandleFunc("/api/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/api/products/{id}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/products/{id}", controllers.DeleteProduct).Methods("DELETE")
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