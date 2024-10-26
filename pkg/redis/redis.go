package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type RedisCache interface {
	Connect() error
	Get(ctx context.Context, key string) (string, bool, error)
	Set(ctx context.Context, key string, value string, expiration int) error
	Delete(ctx context.Context, key string)
}

type RedisConfig struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type cache struct {
	config *RedisConfig
	conn   *redis.Client
}

func NewCache(cfg *RedisConfig) RedisCache {
	return &cache{config: cfg}
}

func (c *cache) Connect() error {
	c.conn = redis.NewClient(&redis.Options{
		Addr:     c.config.Address,
		Password: c.config.Password,
		DB:       c.config.DB,
	})

	_, err := c.conn.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Не удалось подключиться к Redis: %v", err)
	}
	return err
}

// return value,bool isValie in cache and error
func (c *cache) Get(ctx context.Context, key string) (string, bool, error) {
	val, err := c.conn.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		return "", false, err
	}
	return val, true, nil
}

func (c *cache) Set(ctx context.Context, key string, value string, expiration int) error {
	err := c.conn.Set(ctx, key, value, time.Duration(expiration)).Err()
	return err
}
func (c *cache) Delete(ctx context.Context, key string) {
	c.conn.Del(ctx, key)
}
