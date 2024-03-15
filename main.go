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

	router.HandleFunc("/products/add", controllers.AddProduct).Methods("PUT")

	router.HandleFunc("/products/get-details", controllers.GetProductDetails).Methods("GET")
	router.HandleFunc("/products/get-all", controllers.GetAllProductDetails).Methods("GET")
	router.HandleFunc("/products/count-all", controllers.CountProducts).Methods("GET")

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
