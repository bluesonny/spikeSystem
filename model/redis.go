package model

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

//初始化redis连接池
func NewPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   10000,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "localhost:6379")
			if _, err := c.Do("AUTH", "1234"); err != nil {
				_ = c.Close()
				return nil, err
			}

			if err != nil {
				log.Printf("redis连接%v", err.Error())
			} else {
				log.Printf("redis连接成功了")
			}
			return c, err
		},
	}
}
