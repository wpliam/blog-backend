package container

import (
	"blog-backend/internal/common/client"
	"fmt"
	"sync"
)

var containerMap sync.Map

type Factory interface {
	Get(key string) client.Client
	Set(key string, client client.Client)
}

// New 创建一个容器
func New() Factory {
	return &container{}
}

type container struct {
}

func (c *container) Get(key string) client.Client {
	value, ok := containerMap.Load(key)
	if ok {
		return value.(client.Client)
	}
	return nil
}

func (c *container) Set(key string, client client.Client) {
	_, ok := containerMap.Load(key)
	if ok {
		panic(fmt.Sprintf("key %s already exist", key))
	}
	containerMap.Store(key, client)
}
