package main

import (
	"fmt"
	"github.com/adeemgoogle/Bookstore/pkg/db"
	"github.com/adeemgoogle/Bookstore/pkg/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	fmt.Println("hello baby")
	DB := db.Init()
	h := handlers.New(DB)
	r := mux.NewRouter()

	r.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	r.HandleFunc("/books/{id}", h.GetBook).Methods(http.MethodGet)
	r.HandleFunc("/books", h.AddBook).Methods(http.MethodPost)
	r.HandleFunc("/books/{id}", h.UpdateBook).Methods(http.MethodPut)
	r.HandleFunc("/books/{id}", h.DeleteBook).Methods(http.MethodDelete)

	http.ListenAndServe(":8000", r)

}
