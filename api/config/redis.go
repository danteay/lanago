package config

import (
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

// GetRedis return a new Redis connection manager. This connection takes it
// configuration from the next ENV vars.
//
// REDIS_HOST: Redis host and port (localhost:12345)
// REDIS_PASS: Redis pasword
// REDIS_DBAS: Redis index or database
func GetRedis() *redis.Client {
	var db int64
	var err error

	if auxdb := os.Getenv("REDIS_DBAS"); auxdb != "" {
		db, err = strconv.ParseInt(auxdb, 0, 64)

		if err != nil {
			panic(err)
		}
	} else {
		db = 0
	}

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       int(db),
	})

	return client
}
