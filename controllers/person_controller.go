package controllers

import (
	"encoding/json"
	"github.com/JuanSamuelArbelaez/GO_API/model"
	"github.com/JuanSamuelArbelaez/GO_API/services"
	"net/http"
)

func GetPersonDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("ID")
	person, err := services.GetPerson(id)
	if err == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(person)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(model.Person{})

}

func GetAllPersonsDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	inv, err := services.GetPeople()
	if err == nil {
		json.NewEncoder(w).Encode(inv)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode([]model.Person{})
}

func CountPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	size, err := services.GetNumberOfPeople()
	if err == nil {
		json.NewEncoder(w).Encode(size)
		return
	}
	w.WriteHeader(http.StatusOK)
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(id)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("ID")

	state, err := services.RemovePerson(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(state)
}

func RecoverPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("ID")

	state, err := services.RecoverPerson(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(state)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var person model.Person
	err := json.NewDecoder(r.Body).Decode(&person)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := services.UpdatePerson(person); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(person)
}
