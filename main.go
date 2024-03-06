package main

import (
    "fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func nomain() {
	router := mux.NewRouter()
	
	port := ":8087"

	servidor := &http.Server{
		Handler: router,
		Addr:    port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Starting server on port %s. Press CTRL + C to escape\n", port)
	log.Fatal(servidor.ListenAndServe())
}
