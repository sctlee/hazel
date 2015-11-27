package db

import (
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
)

var RedisPool *redis.Pool

type RedisConfig struct {
	Name     string
	Host     string
	Port     string
	Password string
}

func StartRedisPool(config RedisConfig) {
	RedisPool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp",
				strings.Join([]string{config.Host, config.Port}, ":"))
			if err != nil {
				return nil, err
			}
			if len(config.Password) > 0 {
				if _, err := c.Do("AUTH", config.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
