package cache

import (
	"log"
	"sync"
)

func New() Cache {
	var c = &SimpleCache{make(map[string][]byte), sync.RWMutex{}}

	log.Println("ready to serve")
	return c
}

type Cache interface {
	Set(string, []byte) error
	Get(string) ([]byte, error)
	Del(string) error
}

type SimpleCache struct {
	c     map[string][]byte
	mutex sync.RWMutex
}

func (c *SimpleCache) Set(k string, v []byte) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.c[k] = v
	return nil
}

func (c *SimpleCache) Get(k string) ([]byte, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.c[k], nil
}

func (c *SimpleCache) Del(k string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	_, exist := c.c[k]
	if exist {
		delete(c.c, k)
		//c.del(k, v)
	}
	return nil
}
