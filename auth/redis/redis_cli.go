package redis

import (
	"github.com/gomodule/redigo/redis"
)

type RedisCli struct {
	conn redis.Conn
}

var redisCliInstance *RedisCli = nil

func Connect() (conn *RedisCli)  {
	if redisCliInstance == nil {
		redisCliInstance = new(RedisCli)
		var err error

		redisCliInstance.conn, err = redis.Dial("tcp", ":6379")
		if err != nil {
			panic(err)
		}
	}
	return redisCliInstance
}

func (redisCli *RedisCli) SetValue(key string, value string, expiration ...interface{}) error {
	_, err := redisCli.conn.Do("SET", key, value)

	if err != nil && expiration != nil {
		redisCli.conn.Do("EXPIRE", key, expiration[0])
	}
	return err
}

func (redisCli *RedisCli) GetValue(key string) (interface{}, error) {
	return redisCli.conn.Do("GET", key)
}