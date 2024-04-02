package main

import (
	"fmt"
	"github.com/JuanSamuelArbelaez/GO_API/SQL"
	"github.com/JuanSamuelArbelaez/GO_API/Utils"
	"github.com/JuanSamuelArbelaez/GO_API/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	Utils.NotCountryCode("col")

	router := mux.NewRouter()

	SQL.InitDB()

	router.HandleFunc("/people/add", controllers.AddPerson).Methods("PUT")

	router.HandleFunc("/people/get-details", controllers.GetPersonDetails).Methods("GET")
	router.HandleFunc("/people/get-all", controllers.GetAllPersonsDetails).Methods("GET")
	router.HandleFunc("/people/count-all", controllers.CountPersons).Methods("GET")

	port := ":8087"

	servidor := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("Starting server on port %s. Press CTRL + C to escape\n", port)
	log.Fatal(servidor.ListenAndServe())
}
