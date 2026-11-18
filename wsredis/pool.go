package wsredis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

func NewPool(addr, pwd string) *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr, redis.DialPassword(pwd))
		},
		MaxIdle:     10,
		MaxActive:   50, // prevent connection pool from growing indefinitely.
		Wait:        true,
		IdleTimeout: 240 * time.Second,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < 60*time.Second {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}
