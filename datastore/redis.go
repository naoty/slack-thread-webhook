package datastore

import (
	"github.com/gomodule/redigo/redis"
)

const redisKey = "hooks"

// Redis is a client for redis.
type Redis struct {
	URL string
}

// Get returns a value matched to a given key.
func (client *Redis) Get(key string) (string, error) {
	conn, err := redis.DialURL(client.URL)
	if err != nil {
		return "", err
	}

	return redis.String(conn.Do("HGET", redisKey, key))
}

// Set stores a value with a key.
func (client *Redis) Set(key, value string) error {
	conn, err := redis.DialURL(client.URL)
	if err != nil {
		return err
	}

	_, err = conn.Do("HSET", redisKey, key, value)
	return err
}

// Ping checks connection to redis with PING command.
func (client *Redis) Ping() error {
	conn, err := redis.DialURL(client.URL)
	if err != nil {
		return err
	}

	_, err = conn.Do("PING")
	return err
}
