package main

import (
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
)

func CreatePoolRedis() *redis.Pool {
	pool := &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				log.Printf("ERROR: fail init redis pool: %s", err.Error())
				os.Exit(1)
			}
			return conn, err
		},
	}

	return pool
}
