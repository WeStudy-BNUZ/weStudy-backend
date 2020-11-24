package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"strings"
	"time"
)

var Redis redisPool

type redisPool struct {
	pool *redis.Pool
}

func (r *redisPool) Do(db int, common string, args ...interface{}) (interface{}, error) {
	conn := r.pool.Get()
	defer conn.Close()
	_, _ = conn.Do("select", db)
	if strings.ToLower(common) == "set" && len(args) > 2 {
		args = append(args, strconv.Itoa(args[2].(int)))
		args[2] = "EX"
	}

	return conn.Do(common, args...)
}

func InitRedisPool() {
	Redis = redisPool{
		pool: &redis.Pool{
			MaxIdle:     3,
			MaxActive:   50,
			IdleTimeout: 240 * time.Second,
			Dial:        dial,
		},
	}
}

func dial() (redis.Conn, error) {

	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", "reids", "6379"))
	if err != nil {
		fmt.Println("[!] redis缓存服务连接异常，尝试重新链接...", err)
		time.Sleep(time.Second * 5)
		return dial()
	}

	if _, err := conn.Do("AUTH", "123456"); err != nil {
		_ = conn.Close()
		panic(err)
	}
	return conn, nil
}
