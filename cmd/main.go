package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"main.go/pkg/handlers"
)

func main() {
	fmt.Println("hello world")
	router := mux.NewRouter()
	router.HandleFunc("/books", handlers.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books", handlers.AddBook).Methods(http.MethodPost)

	router.HandleFunc("/books/{id}", handlers.GetBook).Methods(http.MethodGet)

	router.HandleFunc("/books/{id}", handlers.UpdateBook).Methods(http.MethodPut)

	router.HandleFunc("/books/{id}", handlers.GetAllBooks).Methods(http.MethodDelete)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello World!")
	})
	log.Println("API is running")
	http.ListenAndServe(":4000", router)
}
