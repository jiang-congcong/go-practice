package redis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func getRedisConnection() redis.Conn {
	do := redis.DialPassword("123456")

	conn, err := redis.Dial("tcp", "localhost:6379", do)
	if err != nil {
		fmt.Println("connection error: ", err)
	}

	return conn
}

func Get(key string) interface{} {
	conn := getRedisConnection()
	result, err := conn.Do("Get", key)

	if err != nil {
		fmt.Println("redis get error, key : ", key)
		return nil
	}

	conn.Close()

	return result
}

func Set(key string, value interface{}) bool {
	conn := getRedisConnection()

	_, err := conn.Do("Set", key, value)
	if err != nil {
		return false
	}

	conn.Close()

	return true
}
