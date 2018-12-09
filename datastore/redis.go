package datastore

import (
	"github.com/gomodule/redigo/redis"
)

const redisKey = "hooks"

// Redis is a client for redis.
type Redis struct {
	host string
}

// NewRedis returns a Redis.
func NewRedis(host string) *Redis {
	return &Redis{host: host}
}

// Get returns a value matched to a given key.
func (client *Redis) Get(key string) (string, error) {
	conn, err := redis.Dial("tcp", client.host)
	if err != nil {
		return "", err
	}

	value, err := conn.Do("HGET", redisKey, key)
	if err != nil {
		return "", err
	}

	return value.(string), nil
}

// Set stores a value with a key.
func (client *Redis) Set(key, value string) error {
	conn, err := redis.Dial("tcp", client.host)
	if err != nil {
		return err
	}

	_, err = conn.Do("HSET", redisKey, key, value)
	if err != nil {
		return err
	}

	return nil
}
