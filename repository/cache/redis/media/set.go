package mediaredis

import (
	"context"
	"fmt"
	"time"
)

func (c *Cache) SetMediaURLByMediaID(ctx context.Context, mediaID uint, url string) error {

	key := fmt.Sprintf("media_url:%d", mediaID)
	ttl := 3600 * time.Second

	err := c.conn.Set(ctx, key, url, ttl).Err()
	if err != nil {
		return err
	}
	return nil
}
