// +build !prod

package testing

import (
	"log"
	"testing"

	"github.com/go-redis/redis"
)

func TestRedistSetValue(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:7891",
		DB:   0, // use default DB
	})

	err := client.Set("mylove", "nuke", 0).Err()
	if err != nil {
		panic(err)
	}

}

func TestRedisGetValue(t *testing.T) {

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:7891",
		DB:   0, // use default DB
	})

	val, err := client.Get("mylove").Result()
	if err != nil {
		panic(err)
	}
	log.Fatal("mylove:", val)
}
