package common

import (
	"github.com/gomodule/redigo/redis"
	"time"
	"github.com/robfig/config"
	"os"
	"fmt"
)

var RedisPool *redis.Pool

func init() {
	c, err := config.ReadDefault("conf.ini")
	if err != nil {
		os.Exit(1)
	}
	host, err := c.String("database", "host")
	if err != nil {
		os.Exit(1)
	}
	port, err := c.Int("database", "port")
	if err != nil {
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
