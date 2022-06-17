package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/tmnhs/fginx/server/global"
	"github.com/tmnhs/fginx/server/internal/model"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type RedisService struct{}

func (redisService *RedisService) GetIntFromRedis(ctx context.Context, key string) (int64, error) {
	s, err := redisService.GetStringFromRedis(ctx, key)
	if err != nil {
		return 0, err
	}
	appletID, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return appletID, nil
}

func (redisService *RedisService) HgetIntFromRedis(ctx context.Context, key string, value string) (int64, error) {
	s, err := redisService.HgetStringFromRedis(ctx, key, value)
	if err != nil {
		return 0, err
	}
	id, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (redisService *RedisService) GetStringFromRedis(ctx context.Context, key string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 3000*time.Millisecond)
	defer cancel()

	result, err := global.GV_REDIS.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", model.ErrRedisNotFound
		}
		return "", err
	}
	return result, nil
}

func (redisService *RedisService) HgetStringFromRedis(ctx context.Context, key string, value string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 3000*time.Millisecond)
	defer cancel()

	result, err := global.GV_REDIS.HGet(ctx, key, value).Result()
	if err != nil {
		if err == redis.Nil {
			return "", model.ErrRedisNotFound
		}
		return "", err
	}
	return result, nil
}

func (redisService *RedisService) GetFromRedis(ctx context.Context, key string, object interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, 3000*time.Millisecond)
	defer cancel()

	result, err := global.GV_REDIS.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return model.ErrRedisNotFound
		}
		return err
	}

	global.GV_LOG.Debug("key:%s result to unmarshal: %s", zap.String("redis", key), zap.String("redis", result))
	if err := json.Unmarshal([]byte(result), object); err != nil {
		return fmt.Errorf("unmarshal error")
	}
	return nil
}

func (redisService *RedisService) SetToRedis(ctx context.Context, key string, object interface{}, expire time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, 3000*time.Millisecond)
	defer cancel()

	buf, err := json.Marshal(object)
	if err != nil {
		global.GV_LOG.Error("marshal error:%s", zap.Error(err))
		return err
	}
	_, err = global.GV_REDIS.Set(ctx, key, buf, expire).Result()
	if err != nil {
		return err
	}
	return nil
}

func (redisService *RedisService) DelFromRedis(ctx context.Context, key string) error {
	ctx, cancel := context.WithTimeout(ctx, 3000*time.Millisecond)
	defer cancel()

	_, err := global.GV_REDIS.Del(ctx, key).Result()
	return err
}

func (redisService *RedisService) GetIntArrayFromRedis(ctx context.Context, key string) ([]int64, error) {
	var a = make([]int64, 0)
	err := redisService.GetFromRedis(ctx, key, &a)
	if err == nil { // Redis Found
		return a, nil
	}
	if err != model.ErrRedisNotFound { // Redis Error
		return nil, err
	}
	// Redis Not Found
	return a, nil
}
