package cache

import (
	json2 "encoding/json"
	"log"
	"time"

	db "github.com/Tech-With-Tim/cdn/db/sqlc"
	"github.com/go-redis/redis/v7"
)

type redisCache struct {
	host    string
	db      int
	pass    string
	expires time.Duration
}

func NewRedisCache(host string, db int, pass string, expires time.Duration) PostCache {
	return &redisCache{
		host:    host,
		db:      db,
		pass:    pass,
		expires: expires,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	log.Println("Trying to connect to redis")
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: cache.pass,
		DB:       cache.db,
	})
}

func (cache *redisCache) Set(key string, value *db.GetFileRow) {
	client := cache.getClient()
	json, err := json2.Marshal(value)
	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("Added to cache: %s", key)
	client.Set(key, json, cache.expires*time.Minute)
}

func (cache *redisCache) Get(key string) *db.GetFileRow {
	client := cache.getClient()

	val, err := client.Get(key).Result()
	if err != nil {
		return nil
	}
	fileRow := db.GetFileRow{}
	err = json2.Unmarshal([]byte(val), &fileRow)
	if err != nil {
		log.Println(err.Error())
	}
	return &fileRow
}
