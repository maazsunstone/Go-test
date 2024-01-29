package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"

	"main.go/pkg/mocks"
	"main.go/pkg/models"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(body, &book)
	book.Id = rand.Intn(100)
	mocks.Book = append(mocks.Book, book)
	fmt.Println(mocks.Book)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Created")

}
