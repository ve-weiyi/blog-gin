package initialize

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/ve-weiyi/ve-blog-gin/config"
)

func Redis(cfg *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
}
