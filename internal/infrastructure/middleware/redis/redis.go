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
	"fmt"
	"smartCloudBean/internal/common/log"
	"smartCloudBean/internal/infrastructure/middleware/apollo"

	"github.com/go-redis/redis/v8"
)

// 定义全局变量
var RedisCache *redis.Client

func init() {

	// Redis 连接信息
	addr, _ := apollo.ApolloCache.Get("redis.host")
	password, _ := apollo.ApolloCache.Get("redis.password")
	db, _ := apollo.ApolloCache.Get("redis.db.jlx_token")

	// 构建 Redis URI
	redisURI := fmt.Sprintf("redis://%s@%s/%s", password, addr, db)

	options, err := redis.ParseURL(redisURI)
	if err != nil {
		log.Logger.Error("failed to parse redis URI: %v", err)
	}

	RedisCache = redis.NewClient(options)
	log.Logger.Info("Connected to Redis successfully!")

}
