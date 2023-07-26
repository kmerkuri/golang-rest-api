package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kmerkuri/golang-crud-rest-api/database"
	"github.com/kmerkuri/golang-crud-rest-api/entities"
)

func GetProducts(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var product []entities.Product
	database.Instance.Find(&product)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var product entities.Product
	json.NewDecoder(r.Body).Decode(&product)
	database.Instance.Create(&product)
	json.NewEncoder(w).Encode(product)

}
func GetProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	productID ;= mux.Vars("id")
	if checkifProductExists(productID) == false {
		json.NewEncoder(w).Encode("Product not found")
	}
	var product entities.Product
	database.Instance.First(&product, productId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)


}
func checkifProductExists(productID) {
	w.Header().Set("Content-Type","application/json")
	var product entities.Product
	database.Instance.Find(&product,productID)
	if product.ID == 0 {
		return false
	}
	return true
}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	productID ;= mux.Vars("id")
	if checkifProductExists(productID) == false {
		json.NewEncoder(w).Encode("Product not found")
	}
	var product entities.Product
	database.Instance.First(&product, productId)
	json.NewDecoder(r.Body).Decode(&product)
	database.Instance.Save(&product)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	productId := mux.Vars(r)["id"]
	if checkIfProductExists(productId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product entities.Product
	database.Instance.Delete(&product, productId)
	json.NewEncoder(w).Encode("Product Deleted Successfully!")
}