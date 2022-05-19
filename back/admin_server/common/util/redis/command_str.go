package redis

import (
	"context"

	"github.com/gomodule/redigo/redis"
)

// 设置缓存
func Set(ctx context.Context, key string, value interface{}, expire int) error {
	rds := getRedisDB(ctx)
	defer rds.Close()
	_, err := rds.db.Do("SET", key, value, "EX", expire)
	return err
}

// 获取数据
func Get(ctx context.Context, key string) (interface{}, error) {
	rds := getRedisDB(ctx)
	defer rds.Close()
	reply, err := rds.db.Do("GET", key)
	if err == redis.ErrNil {
		return nil, nil
	}
	return reply, err
}

// key是否存在
func Exists(ctx context.Context, key string) (bool, error) {
	rds := getRedisDB(ctx)
	defer rds.Close()
	return redis.Bool(rds.db.Do("EXISTS", key))
}

// 获取字节切片
func Bytes(ctx context.Context, key string) ([]byte, error) {
	return redis.Bytes(Get(ctx, key))
}
