package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/adeemgoogle/Bookstore/pkg/cache"
	"github.com/adeemgoogle/Bookstore/pkg/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h handler) Getuser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var user models.User
	redis := cache.RedisCache{}
	name := redis.Get("id")
	if name != nil {
		if result := h.DB.Find(&user, id); result.Error != nil {
			fmt.Println(result.Error)
		}
	}
	redis.Set("id", name)
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}
