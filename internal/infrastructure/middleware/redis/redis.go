// Package redis
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     redis.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/06/27 19:11:21
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package redis

import (
	"context"
	"fmt"
	"smartCloudBean/internal/common/log"
	"smartCloudBean/internal/infrastructure/middleware/apollo"
	"strconv"

	"github.com/go-redis/redis/v8"
)

// 定义全局变量
var RedisCache *redis.Client

func init() {
	// Redis 连接信息
	addrInterface, _ := apollo.ApolloCache.Get("redis.host")
	passInterface, _ := apollo.ApolloCache.Get("redis.password")
	dbInterface, _ := apollo.ApolloCache.Get("redis.db.jlx_token")

	dbInt, err := strconv.ParseInt(dbInterface.(string), 10, 64) // 10 表示十进制，64 表示 int64 类型
	if err != nil {
		log.Logger.Error("Error converting string to integer: %+v", err)
		panic("connect to Redis failed")
	}

	// 构建 Redis URI
	redisURI := fmt.Sprintf("redis://%s@%s/%s", passInterface, addrInterface, dbInterface)

	// 设置密码
	options := &redis.Options{
		Password: passInterface.(string),
		Addr:     addrInterface.(string),
		DB:       int(dbInt),
	}

	// 初始化 Redis 客户端
	RedisCache = redis.NewClient(options)

	// 可选：设置连接池参数
	RedisCache.Options().MaxRetries = 3
	RedisCache.Options().MinIdleConns = 10

	// 可选：测试连接
	if err := RedisCache.Ping(context.Background()).Err(); err != nil {
		log.Logger.Error("connect to Redis failed: %+v", err)
		panic("connect to Redis failed")
	}
	log.Logger.Info("连接Redis：%s 正常.", redisURI)
}
