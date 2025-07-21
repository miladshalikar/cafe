package mediaredis

import (
	"context"
	"fmt"
)

func (c *Cache) GetMediaURLByMediaID(ctx context.Context, mediaID uint) (string, error) {
	key := fmt.Sprintf("media_url:%d", mediaID)
	return c.conn.Get(ctx, key).Result()
}
