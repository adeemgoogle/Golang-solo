package cache

import "github.com/adeemgoogle/Bookstore/pkg/models"

type PostCache interface {
	Set(key string, value *models.User)
	Get(key string) *models.User
}
