package database

import (
	"context"
	"crypto/tls"
	"os"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:      os.Getenv("DB_ADDR"),
		Password:  os.Getenv("DB_PASS"),
		DB:        dbNo,
		TLSConfig: &tls.Config{InsecureSkipVerify: false},
	})
	return rdb
}
