// +build !prod

package config

import "github.com/go-redis/redis"

const RedisHost = "localhost:7891"

var RedisOptions = redis.Options{
	Addr: RedisHost,
	DB:   0, // use default DB
}
