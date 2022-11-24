package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/adeemgoogle/Bookstore/pkg/models"
	"net/http"
)

func (h handler) Byname(w http.ResponseWriter, r *http.Request) {

	var name []models.User
	if result := h.DB.Find(&name); result.Error != nil {
		fmt.Println(result.Error)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(name)

}
