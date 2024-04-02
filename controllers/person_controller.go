package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/JuanSamuelArbelaez/GO_API/model"
	"github.com/JuanSamuelArbelaez/GO_API/services"
	"net/http"
)

func GetPersonDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("ID")
	fmt.Println(id)
	person, err := services.GetPerson(id)
	if err == nil {
		json.NewEncoder(w).Encode(person)
		return
	}
	json.NewEncoder(w).Encode(model.Person{})

}

func GetAllPersonsDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	inv, err := services.GetPeople()
	if err == nil {
		json.NewEncoder(w).Encode(inv)
		return
	}
	json.NewEncoder(w).Encode([]model.Person{})
}

func CountPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	size, err := services.GetNumberOfPeople()
	if err == nil {
		json.NewEncoder(w).Encode(size)
		return
	}
	json.NewEncoder(w).Encode(0)
}

func AddPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newPerson model.PersonRequest
	err := json.NewDecoder(r.Body).Decode(&newPerson)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := services.AddPerson(newPerson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(id)
}
