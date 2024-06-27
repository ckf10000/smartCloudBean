// Package redis
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     redis.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/06/28 01:04:51
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package redis

import (
	"context"
	"smartCloudBean/internal/common/log"
	client "smartCloudBean/internal/infrastructure/middleware/redis"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
}

type redisClient struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisClient() RedisClient {
	return &redisClient{client: client.RedisCache, ctx: context.Background()}
}

func (r *redisClient) Set(key string, value interface{}, expiration time.Duration) error {
	log.Logger.Info("Setting key %s in Redis", key)
	return r.client.Set(r.ctx, key, value, expiration).Err()
}

func (r *redisClient) Get(key string) (string, error) {
	log.Logger.Info("Getting key %s from Redis", key)
	return r.client.Get(r.ctx, key).Result()
}
