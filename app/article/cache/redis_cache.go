package cache

import (
	"encoding/json"
	"log"
	"time"

	"github.com/fjasper13/endpoint-article/app/article/entities"
	"github.com/go-redis/redis"
)

type redisCache struct {
	host     string
	db       int
	password string
	expires  time.Duration
}

func NewRedisCache(host string, db int, pass string, expires time.Duration) ArticleCache {
	return &redisCache{
		host:     host,
		db:       db,
		password: pass,
		expires:  expires,
	}
}

func (c *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.host,
		Password: c.password,
		DB:       c.db,
	})
}

func (c *redisCache) Set(key string, value *entities.Article) {
	client := c.getClient()

	json, err := json.Marshal(value)
	if err != nil {
		log.Fatal(err)
	}

	client.Set(key, json, c.expires*time.Second)
}

func (c *redisCache) Get(key string) *entities.Article {
	client := c.getClient()

	val, err := client.Get(key).Result()
	if err != nil {
		return nil
	}

	article := entities.Article{}
	err = json.Unmarshal([]byte(val), &article)
	if err != nil {
		log.Fatal(err)
	}

	return &article
}
