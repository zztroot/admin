package redis

import (
	"context"
	"fmt"
	"time"

	"goods/config"

	"github.com/gomodule/redigo/redis"
)

var rdsPool *redis.Pool

type ZRedis struct {
	db  redis.Conn
	ctx context.Context
}

func getRedisDB(ctx context.Context) *ZRedis {
	return &ZRedis{
		ctx: ctx,
		db:  rdsPool.Get(),
	}
}

func InitRedis(conf *config.AppConfig) {
	address := fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port)
	rdsPool = &redis.Pool{
		Wait:        true,
		MaxIdle:     conf.Redis.MaxIdle,
		MaxActive:   conf.Redis.MaxActive,
		IdleTimeout: time.Duration(conf.Redis.IdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", address, redis.DialPassword(conf.Redis.Password))
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
}

// 关闭连接
func (z *ZRedis) Close() {
	_ = z.db.Close()
}
