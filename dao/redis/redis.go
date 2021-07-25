package redis

import (
	"bluebellAPI/settings"
	"fmt"
	"github.com/go-redis/redis"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

// Init 初始化redis连接
func Init(cfg *settings.RedisConfig) (err error) {
	// 初始化客户端配置
	client = redis.NewClient(&redis.Options{
		// host:port
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})
	// 连接
	_, err = client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

// Close 关闭redis连接
func Close() {
	_ = client.Close()
}
