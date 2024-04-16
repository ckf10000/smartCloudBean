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
	"net/http"
	"smartCloudBean/internal/application/services"
	"smartCloudBean/internal/converter"
	"smartCloudBean/internal/domain/dto"

	"github.com/ckf10000/gologger/v3/log"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func StartServer(log *log.FileLogger, db *gorm.DB) {
	r := gin.Default()

	// 设置静态文件路由
	r.Static("/static", "./static")

	orderService := services.NewOrderService(db, log) // 初始化应用服务

	r.GET("/api/v1/orders", func(c *gin.Context) {
		passenger := c.DefaultQuery("passenger", "")
		preOrderID := converter.ConvertQueryInt(c.DefaultQuery("pre_order_id", "-1"), -1)
		page := converter.ConvertQueryInt(c.DefaultQuery("page", "1"), 1)
		limit := converter.ConvertQueryInt(c.DefaultQuery("limit", "10"), 10)

		orders, total_count, err := orderService.FindOrders(preOrderID, passenger, page, limit, log)
		paginationDta := dto.ResponsePaginationOrder{
			Data:       orders,
			TotalCount: total_count,
			Page:       page,
			PageSize:   limit,
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 400, "message": err.Error(), "data": paginationDta})
			log.Error("request for '/api/v1/orders' interface failed, reason: %s", err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "查询成功", "data": paginationDta})
	})

	r.Run(":8080")
}
