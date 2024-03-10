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
	p, err := services.GetProduct(id)
	if err == nil {
		fmt.Println(p)
		json.NewEncoder(w).Encode(p)
		return
	}
	json.NewEncoder(w).Encode(model.Product{})

}
