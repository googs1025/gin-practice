package redis

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"testing"
)

func TestRedisOperation(t *testing.T) {

	conn := RedisPool.Get()
	res, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		log.Println(err)
	}
	log.Println(res)
}
