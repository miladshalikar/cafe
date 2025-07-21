package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type Config struct {
	Host     string `koanf:"host"`
	Port     string `koanf:"port"`
	Password string `koanf:"password"`
	DB       int    `koanf:"db"`
}

type Redis struct {
	cfg    Config
	client *redis.Client
}

func (r Redis) Conn() *redis.Client {
	return r.client
}

func New(cfg Config) *Redis {

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("❌ Redis connection failed: %v", err)
	}

	log.Println("✅ Redis connected successfully")

	return &Redis{cfg: cfg, client: client}
}

//func (r *Redis) Set(ctx context.Context, key string, value string, ttlSeconds int) error {
//	return r.client.Set(ctx, key, value, time.Duration(ttlSeconds)*time.Second).Err()
//}
//
//func (r *Redis) Get(ctx context.Context, key string) (string, error) {
//	val, err := r.client.Get(ctx, key).Result()
//	if err == redis.Nil {
//		return "", nil
//	}
//	return val, err
//}
//
//func (r *Redis) Del(ctx context.Context, key string) error {
//	return r.client.Del(ctx, key).Err()
//}
