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

	if products, err := SQL.CountProducts(); err != nil {
		fmt.Println(err)
	} else if products <= 0 {
		SQL.InsertDataSet()
	}

	router.HandleFunc("/GetProductDetails", controllers.GetProductDetails).Methods("GET")
	//router.HandleFunc("/", controllers.SetProductName)
	port := ":8087"

	servidor := &http.Server{
		Handler:      router,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("Starting server on port %s. Press CTRL + C to escape\n", port)
	log.Fatal(servidor.ListenAndServe())
}
