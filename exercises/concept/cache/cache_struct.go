package cache

import (
	"errors"
	"time"
)

// Don't change this file

type Cache struct {
	ResponseTime time.Duration
	Data         map[string]string
}

var ErrKeyNotExist = errors.New("key error")

func (c *Cache) Get(key string) (string, error) {
	time.Sleep(c.ResponseTime)
	if c.Data[key] != "" {
		return c.Data[key], nil
	} else {
		return "", ErrKeyNotExist
	}
}

func (c *Cache) Set(key string, value string) error {
	time.Sleep(c.ResponseTime)
	c.Data[key] = value

	return nil
}
