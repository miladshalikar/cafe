package mediaredis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func (c *Cache) MGetMediaURLs(ctx context.Context, mediaIDs []uint) (map[uint]string, error) {

	result := make(map[uint]string)

	keys := make([]string, len(mediaIDs))
	keyToID := make(map[string]uint, len(mediaIDs))
	for i, id := range mediaIDs {
		key := fmt.Sprintf("media_url:%d", id)
		keys[i] = key
		keyToID[key] = id
	}

	values, err := c.conn.MGet(ctx, keys...).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}

	for i, v := range values {
		if v == nil {
			continue
		}
		strVal, ok := v.(string)
		if !ok {
			continue
		}
		id := keyToID[keys[i]]
		result[id] = strVal
	}

	return result, nil
}
