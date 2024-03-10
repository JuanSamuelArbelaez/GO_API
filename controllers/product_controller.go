package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/JuanSamuelArbelaez/GO_API/model"
	"github.com/JuanSamuelArbelaez/GO_API/services"
	"net/http"
)

func GetProductDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("ID")
	fmt.Println(id)
	pro, err := services.GetProduct(id)
	if err == nil {
		json.NewEncoder(w).Encode(pro)
		return
	}
	json.NewEncoder(w).Encode(model.Product{})

}

func GetAllProductDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	inv, err := services.GetInventory()
	if err == nil {
		json.NewEncoder(w).Encode(inv)
		return
	}
	json.NewEncoder(w).Encode([]model.Product{})
}

func CountProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	size, err := services.GetInventorySize()
	if err == nil {
		json.NewEncoder(w).Encode(size)
		return
	}
	json.NewEncoder(w).Encode(0)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newProduct model.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.AddProduct(newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Product added successfully")
}
