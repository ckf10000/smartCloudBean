// Package mysql
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     mysql.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/06/27 19:37:31
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package mysql

import (
	"fmt"
	"smartCloudBean/internal/common/log"
	"smartCloudBean/internal/infrastructure/middleware/apollo"
	"time"

	_ "github.com/go-sql-driver/mysql" // 导入 MySQL 驱动程序
	"github.com/jinzhu/gorm"
)

// 定义全局变量
var Db *gorm.DB

func init() {

	// MySQL 连接信息
	mysqlUser, _ := apollo.ApolloCache.Get("mysql.username")
	mysqlPassword, _ := apollo.ApolloCache.Get("mysql.password")
	mysqlHost, _ := apollo.ApolloCache.Get("mysql.host")

	mysqlPort, _ := apollo.ApolloCache.Get("mysql.port")
	mysqlDatabase, _ := apollo.ApolloCache.Get("mysql.db")

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDatabase)

	var err error
	// 初始化数据库连接
	Db, err = gorm.Open("mysql", mysqlURI)
	if err != nil {
		log.Logger.Error("failed to connect database: %s, reason: %s", mysqlURI, err)
		panic("failed to connect database")
	}

	// 设置连接池的最大连接数
	Db.DB().SetMaxOpenConns(10)
	// 设置连接的最大空闲时间
	Db.DB().SetConnMaxLifetime(time.Hour)

	// Ping 测试数据库连接
	err = Db.DB().Ping()
	if err != nil {
		str := fmt.Errorf("failed to ping mysql: %v", err)
		panic(str)
	}

	log.Logger.Info("连接Mysql数据库：%s 正常.", mysqlURI)
}
