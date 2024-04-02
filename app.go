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
	"net/http"

	"github.com/ckf10000/gologger/core"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	log := core.NewLogger("debug", "", "", 0, false, true, false)
	log.Info("开启启动web服务...")

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
		log.Fatal("Server failed to start: %s", err)
	}
}
