package main

import (
	"fmt"
	"github.com/adeemgoogle/Bookstore/pkg/db"
	"github.com/adeemgoogle/Bookstore/pkg/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	fmt.Println("helloooowowoowo")
	DB := db.Init()
	h := handlers.New(DB)
	r := mux.NewRouter()

	r.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	r.HandleFunc("/books/{id}", h.GetBook).Methods(http.MethodGet)
	r.HandleFunc("/books", h.AddBook).Methods(http.MethodPost)
	r.HandleFunc("/books/{id}", h.UpdateBook).Methods(http.MethodPut)
	r.HandleFunc("/books/{id}", h.DeleteBook).Methods(http.MethodDelete)
	r.HandleFunc("/names", h.Byname).Methods(http.MethodGet)
	r.HandleFunc("/names", h.AddUser).Methods(http.MethodPost)
	r.HandleFunc("/names/{id}", h.Getuser).Methods(http.MethodGet)

	http.ListenAndServe(":8000", r)

}
