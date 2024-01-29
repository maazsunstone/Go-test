package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"main.go/pkg/mocks"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars((r))
	id, _ := strconv.Atoi(vars["id"])

	for _, i := range mocks.Book {
		if i.Id == id {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusFound)
			json.NewEncoder(w).Encode(i)
		}
	}

}
