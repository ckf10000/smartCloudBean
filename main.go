// Package main
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     main.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/04/02 10:43:04
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package main

import (
	"fmt"
	"smartCloudBean/internal/web"

	"smartCloudBean/internal/infrastructure/middleware"

	"github.com/ckf10000/gologger/v3/log"
	_ "github.com/go-sql-driver/mysql" // 导入 MySQL 驱动程序
	"github.com/jinzhu/gorm"
)

func main() {
	// 获取可执行文件所在目录的路径
	exeDir := log.GetExecuteFilePath()
	if exeDir == "" {
		return
	}

	// projectPath := "./"
	log := log.NewLogger("debug", exeDir, "web.log", "simple", 50*1024*1024, true, true, true)
	cache := middleware.GetApolloCache(log)

	// MySQL 连接信息
	mysqlUser, _ := cache.Get("mysql.username")
	mysqlPassword, _ := cache.Get("mysql.password")
	mysqlHost, _ := cache.Get("mysql.host")

	mysqlPort, _ := cache.Get("mysql.port")
	mysqlDatabase, _ := cache.Get("mysql.db")

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDatabase)

	// 初始化数据库连接
	db, err := gorm.Open("mysql", mysqlURI)
	if err != nil {
		log.Error("failed to connect database: %s, reason: %s", mysqlURI, err)
		panic("failed to connect database")
	}
	log.Info("连接Mysql数据库：%s 正常.", mysqlURI)
	defer db.Close()
	web.StartServer(log, db)
}
