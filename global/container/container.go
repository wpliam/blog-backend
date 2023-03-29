package container

import (
	"blog-backend/internal/api"
	"fmt"
	"sync"
)

var (
	containerMap     sync.Map
	DefaultContainer = &container{}
)

type Container interface {
	Get(key string) api.Client
	Set(key string, client api.Client)
}

type container struct {
}

func (c *container) Get(key string) api.Client {
	value, ok := containerMap.Load(key)
	if ok {
		return value.(api.Client)
	}
	return nil
}

func (c *container) Set(key string, client api.Client) {
	_, ok := containerMap.Load(key)
	if ok {
		panic(fmt.Sprintf("key %s already exist", key))
	}
	containerMap.Store(key, client)
}
