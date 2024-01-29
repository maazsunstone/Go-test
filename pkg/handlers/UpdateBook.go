package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"main.go/pkg/mocks"
	"main.go/pkg/models"
)

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	fmt.Println(id)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var updatedBookVars models.Book
	json.Unmarshal(body, &updatedBookVars)
	for index, i := range mocks.Book {
		if i.Id == id {
			i.Title = updatedBookVars.Title
			i.Author = updatedBookVars.Author
			i.Desc = updatedBookVars.Desc
			mocks.Book[index] = i
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Updated")
			break

		}
	}

}
