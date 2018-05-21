package lru1

import (
	"container/list"
)

type Item struct {
	key string
	value string
}

type Cache struct {
	cap int
	data map[string]*list.Element
	l *list.List
}

func NewCache(cap int) *Cache {
	return &Cache{cap, make(map[string]*list.Element), list.New()}
}

func (c *Cache) Put(key, value string) {
	if len(c.data) == c.cap {
		delete(c.data, c.l.Back().Value.(*Item).key)
		c.l.Remove(c.l.Back())
	}

	c.data[key] = c.l.PushFront(&Item{key: key, value: value})
}

func (c *Cache) Get(key string) *Item {
	if c.data[key] != nil {
		c.l.MoveToFront(c.data[key])
		return c.data[key].Value.(*Item)
	}

	return nil
}
