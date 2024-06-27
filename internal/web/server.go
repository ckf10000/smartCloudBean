// Package web
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     server.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/04/15 10:01:18
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package web

import (
	orderApp "smartCloudBean/internal/application/order"
	smsApp "smartCloudBean/internal/application/sms"

	"smartCloudBean/internal/handler"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	// 设置静态文件路由
	r.Static("/static", "./static")
	findSmsHandler := &handler.FindSmsHandler{
		SmsService: smsApp.NewSmsService(), // 初始化短信服务
	}

	findOrdersHandler := &handler.FindOrdersHandler{
		OrderService: orderApp.NewOrderService(), // 初始化订单服务
	}

	sendSmsHandler := &handler.SendSmsHandler{
		SmsService:      smsApp.NewSmsService(),      // 初始化短信服务
		SmsProxyService: smsApp.NewSmsProxyService(), // 初始化短信代理服务
	}

	r.GET("/api/v1/healthCheck", handler.HealthCheck)
	r.GET("/api/v1/orders", findOrdersHandler.FindOrdersHandler)
	r.GET("/api/v1/sms", findSmsHandler.FindSmsHandler)
	r.POST("/api/v1/smsProxy", sendSmsHandler.SendSmsHandler)
	r.Run(":8081")
}
