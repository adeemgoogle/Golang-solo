package db

import (
	"github.com/adeemgoogle/Bookstore/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:123@localhost:5432"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)

	}
	db.AutoMigrate(&models.Book{}, &models.User{})
	return db
}
