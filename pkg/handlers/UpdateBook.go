package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/adeemgoogle/Bookstore/pkg/models"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func (h handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)

	}
	var updateBook models.Book
	json.Unmarshal(body, &updateBook)

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Updated")

}

