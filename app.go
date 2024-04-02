// Package main
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     app.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/04/02 10:43:04
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	// "github.com/ckf10000/gologger/core"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 设置日志输出格式
	gin.DefaultWriter = io.MultiWriter(os.Stdout)

	// 设置 Gin 的日志输出格式
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 自定义日志格式
		return fmt.Sprintf("%s - [%s] - [%s] - [%s] - %s - %s\n",
			param.TimeStamp.Format("2006-01-02 15:04:05.000"),
			param.Request.URL,
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.ErrorMessage,
		)
	}))
	// log := core.NewLogger("debug", "", "", 0, false, true, false)
	// log.Info("开启启动web服务...")

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	r.POST("/post-example", func(c *gin.Context) {
		var request struct {
			Name string `json:"name"`
		}

		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Hello " + request.Name})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
