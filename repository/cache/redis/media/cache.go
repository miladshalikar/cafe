package mediaredis

import (
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	conn *redis.Client
}

func New(r *redis.Client) *Cache {
	return &Cache{conn: r}
}
