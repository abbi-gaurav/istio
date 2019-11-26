package db

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis"

	"istio.io/istio/mixer/adapter/mygrpcadapter/internal/options"
)

type DB struct {
	client *redis.Client
}

func New() *DB {
	client := redis.NewClient(
		&redis.Options{
			Addr:     options.GlobalConfig.RedisURL,
			Password: "",
			DB:       0,
		})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	if err != nil {
		log.Panicf("failed to connect to Redis %s", options.GlobalConfig.RedisURL)
	}
	return &DB{client: client}
}

func (db *DB) Store(key string, obj interface{}) error {
	ba, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	sc := db.client.Set(key, string(ba), 0)
	if err = sc.Err(); err != nil {
		log.Printf("Got err while saving %+v", err)
		return err
	}

	return nil
}
