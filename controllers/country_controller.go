package controllers

import (
	"encoding/json"
	"github.com/JuanSamuelArbelaez/GO_API/services"
	"net/http"
)

func ValidateCountryCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	code := r.URL.Query().Get("countryCode")
	notValid := services.NotCountryCode(code)
	json.NewEncoder(w).Encode(!notValid)
	return
}

//env test
