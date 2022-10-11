package redis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

// TODO redis config ini

func init() {
	RedisPool = newPool("127.0.0.1:6379")
}

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle: 3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func () (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}

var (
	RedisPool *redis.Pool
	//redisServer = flag.String("redisServer", ":6379", "")
)
