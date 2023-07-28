package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kmerkuri/golang-crud-rest-api/database"
	"github.com/kmerkuri/golang-crud-rest-api/entities"
)

func GetSales(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var sales []entities.Sales
	database.Instance.Find(&sales)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sales)
}
func GetSalesID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	salesid := mux.Vars(r)["id"]
	if Checkifsaleexists(salesid) == false {
		json.NewEncoder(w).Encode("Sale not found")
	}
	var sales entities.Sales
	database.Instance.First(&sales, salesid)
	json.NewEncoder(w).Encode(sales)

}

func UpdateSales(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	salesid := mux.Vars(r)["id"]
	if Checkifsaleexists(salesid) == false {
		json.NewEncoder(w).Encode("Cant update sales")
		w.WriteHeader(http.StatusNotFound)
	}
	var sales entities.Sales
	database.Instance.First(&sales, salesid)
	json.NewDecoder(r.Body).Decode(&sales)
	database.Instance.Save(&sales)
	json.NewEncoder(w).Encode(sales)
}
func DeleteSales(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	salesid := mux.Vars(r)["id"]
	if checkifProductExists(salesid) == false {
		json.NewEncoder(w).Encode("CAnt find sales")
		return
	}
	var sales entities.Sales
	database.Instance.Delete(&sales, salesid)
	json.NewEncoder(w).Encode("Sale deleted successfully")

}

func Checkifsaleexists(salesid string) bool {
	var sales entities.Sales
	database.Instance.Find(&sales, salesid)
	if sales.ID == 0 {
		return false
	}
	return true
}

func CreateSale(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	salesid := mux.Vars(r)["id"]
	if checkifProductExists(salesid) == false {
		json.NewEncoder(w).Encode("CAnt find sales")
		return
	}
	var sales entities.Sales
	database.Instance.Create(&sales)
	json.NewEncoder(w).Encode("Sale creted")
}
