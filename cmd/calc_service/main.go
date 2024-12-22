package main

import (
	"log"
	"net/http"

	"calculator/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/calculate", handlers.CalculateHandler).Methods("POST")
	
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
