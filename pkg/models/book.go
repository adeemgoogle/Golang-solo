package models

import (
	"sync"
	"time"
)

type Book struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title" `
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

type User struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Cache struct {
	sync.RWMutex
	DefaultExpiration time.Duration
	CleanupInterval   time.Duration
	Items             map[string]Book
}
