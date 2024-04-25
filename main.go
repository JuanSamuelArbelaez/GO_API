package main

import (
	"fmt"
	"github.com/JuanSamuelArbelaez/GO_API/SQL"
	"github.com/JuanSamuelArbelaez/GO_API/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter()

	SQL.InitDB()
	router.HandleFunc("/misc/country/validate", controllers.ValidateCountryCode).Methods("GET")
	router.HandleFunc("/people/add", controllers.AddPerson).Methods("PUT")
	router.HandleFunc("/people/get-details", controllers.GetPersonDetails).Methods("GET")
	router.HandleFunc("/people/get-all", controllers.GetAllPersonsDetails).Methods("GET")
	router.HandleFunc("/people/count-all", controllers.CountPersons).Methods("GET")
	router.HandleFunc("/people/remove", controllers.DeletePerson).Methods("DELETE")
	router.HandleFunc("/people/recover", controllers.RecoverPerson).Methods("POST")
	router.HandleFunc("/people/update", controllers.UpdatePerson).Methods("POST")

	port := ":8088"

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("Starting server on port %s. Press CTRL + C to escape\n", port)
	log.Fatal(server.ListenAndServe())
}
