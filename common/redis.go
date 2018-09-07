package common

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
	"time"
)

var RedisPool *redis.Pool

func init() {

	host, err := Conf.String("redis", "host")
	if err != nil {
		fmt.Println("Config redis->host not found")
		os.Exit(1)
	}
	port, err := Conf.Int("redis", "port")
	if err != nil {
		fmt.Println("Config redis->port not found")
		os.Exit(1)
	}

	RedisPool = &redis.Pool{
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
			if err != nil {
				return nil, err
			}
			c.Do("SELECT", 0)
			return c, nil
		},
	}
}
