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
	client  *redis.Client
}

func NewRedisCache(host string, db int, pass string, expires time.Duration) PostCache {
	cache := &redisCache{
		host:    host,
		db:      db,
		pass:    pass,
		expires: expires,
	}
	cache.getClient()
	return cache
}

func (cache *redisCache) getClient() {
	log.Println("Trying to connect to redis")
	cache.client = redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: cache.pass,
		DB:       cache.db,
	})
}

func (cache *redisCache) Set(key string, value *db.GetFileRow) {
	json, err := json2.Marshal(value)
	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("Added to cache: %s", key)
	cache.client.Set(key, json, cache.expires*time.Minute)
}

func (cache *redisCache) Get(key string) *db.GetFileRow {

	val, err := cache.client.Get(key).Result()
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
