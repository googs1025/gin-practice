package redis

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"testing"
)

func TestRedisOperation(t *testing.T) {

	conn := RedisPool.Get()
	// 可以Do执行原生命令
	res, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		log.Println(err)
	}
	log.Println(res)
}
