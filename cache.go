package cache

import "time"

type Cache struct {
	data       map[string]string
	expiration map[string]time.Time
}

func NewCache() Cache {
	return Cache{data: make(map[string]string),
		expiration: make(map[string]time.Time)}
}

func (c Cache) Get(key string) (string, bool) {
	if deadline, ok := c.expiration[key]; !ok || time.Now().Before(deadline) {
		value, ok := c.data[key]
		return value, ok
	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	c.data[key] = value
}

func (c Cache) Keys() []string {
	keys := make([]string, 0, len(c.data))
	for key := range c.data {
		if deadline, ok := c.expiration[key]; !ok || time.Now().Before(deadline) {
			keys = append(keys, key)
		}
	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.data[key] = value
	c.expiration[key] = deadline
}
