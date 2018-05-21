package lru1

import "time"

type Item struct {
	key string
	value string
	last time.Time
}

type Cache struct {
	cap int
	data map[string]*Item
}

func NewCache(cap int) *Cache {
	return &Cache{cap, make(map[string]*Item),}
}

func (c *Cache) makeSpace() {
	old := &Item{ last: time.Now() }
	var key string
	for k, v := range c.data {
		if v.last.Before(old.last) {
			old = v
			key = k
		}
	}

	delete(c.data, key)
}

func (c *Cache) Put(key, value string) {
	if len(c.data) == c.cap {
		c.makeSpace()
	}

	c.data[key] = &Item{ value: value, last: time.Now() }
}

func (c *Cache) Get(key string) *Item {
	if c.data[key] != nil {
		c.data[key].last = time.Now()
		return c.data[key]
	}

	return nil
}
