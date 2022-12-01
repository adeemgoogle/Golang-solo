package cache

import (
	"github.com/adeemgoogle/Bookstore/pkg/models"
	"github.com/go-redis/redis/v7"
	"github.com/goccy/go-json"
	"time"
)

type RedisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, expires time.Duration) PostCache {
	return &RedisCache{
		host:    "6379",
		db:      1,
		expires: 4,
	}
}

func (cache *RedisCache) getUser() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *RedisCache) Set(key string, value *models.User) {
	client := cache.getUser()
	add, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	client.Set(key, add, cache.expires*time.Second)
}

func (cache *RedisCache) Get(key string) *models.User {
	client := cache.getUser()

	val, err := client.Get(key).Result()
	if err != nil {
		return nil
	}
	name := models.User{}
	err = json.Unmarshal([]byte(val), &name)
	if err != nil {
		panic(err)
	}
	return &name
}
