// Package handler
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     handler.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/06/28 02:43:52
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package handler

import (
	"net/http"
	orderApp "smartCloudBean/internal/application/order"
	smsApp "smartCloudBean/internal/application/sms"
	"smartCloudBean/internal/common/log"
	"smartCloudBean/internal/converter"
	orderDto "smartCloudBean/internal/domain/dto/order"
	smsDto "smartCloudBean/internal/domain/dto/sms"

	"github.com/gin-gonic/gin"
)

type FindSmsHandler struct {
	SmsService *smsApp.SmsService
}

type SendSmsHandler struct {
	SmsService      *smsApp.SmsService
	SmsProxyService *smsApp.SmsProxyService
}

type FindOrdersHandler struct {
	OrderService *orderApp.OrderService
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "查询成功", "data": struct{}{}})
}

func (h *FindSmsHandler) FindSmsHandler(c *gin.Context) {
	phoneNum := c.DefaultQuery("phone_num", "")
	searchValue := c.DefaultQuery("search_value", "")
	page := converter.ConvertQueryInt(c.DefaultQuery("page", "1"), 1, 1)
	limit := converter.ConvertQueryInt(c.DefaultQuery("limit", "10"), 10, 10)
	sms, total_count, err := h.SmsService.FindSms(phoneNum, searchValue, page, limit)
	paginationDta := smsDto.ResponsePaginationSms{
		Data:       sms,
		TotalCount: total_count,
		Page:       page,
		PageSize:   limit,
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 400, "message": err.Error(), "data": paginationDta})
		log.Logger.Error("request for '/api/v1/sms' interface failed, reason: %s", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "查询成功", "data": paginationDta})
}

func (h *SendSmsHandler) SendSmsHandler(c *gin.Context) {
	var requestBody struct {
		Text string `json:"text"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		log.Logger.Error("Error decoding request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Logger.Info("Received request with text: \n%s\n", requestBody.Text)

	// 解析文本内容
	context, phone, code := h.SmsProxyService.ParseSms(requestBody.Text)
	if context != "" {
		sms := &smsDto.SmsInsert{
			PhoneNum:                phone,
			Context:                 context,
			DigitalVerificationCode: code,
		}

		err := h.SmsService.InsertSms(sms)
		if err != nil {
			log.Logger.Error("failed to insert sms: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "code": 400, "message": err.Error()})
			return
		}
		log.Logger.Info("SMS context insert to the database success")
		if code != "" {
			token, err := h.SmsProxyService.SmsProxyRepository.GetToken("yundou", "jlx123")
			if err != nil {
				log.Logger.Error("Error getting token: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"success": false, "code": 400, "message": err.Error()})
				return
			}

			err = h.SmsProxyService.SmsProxyRepository.SendSms(token, context, phone)
			if err != nil {
				log.Logger.Error("Error sending SMS: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"success": false, "code": 400, "message": err.Error()})
				return
			}

			log.Logger.Info("SMS content forward jlx api success")
			c.JSON(http.StatusOK, gin.H{"success": true, "code": 0, "message": "success"})
		} else {
			log.Logger.Warning("SMS content does not involve ticket channel login verification code, ignored")
			c.JSON(http.StatusOK, gin.H{"success": true, "code": 0, "message": "success"})
		}
	} else {
		log.Logger.Warning("SMS content parse exception")
		c.JSON(http.StatusOK, gin.H{"success": false, "code": 500, "message": "SMS content parse exception"})
	}
}

func (h *FindOrdersHandler) FindOrdersHandler(c *gin.Context) {
	passenger := c.DefaultQuery("passenger", "")
	preOrderID := converter.ConvertQueryInt(c.Query("pre_order_id"), -1, 0)
	page := converter.ConvertQueryInt(c.DefaultQuery("page", "1"), 1, 1)
	limit := converter.ConvertQueryInt(c.DefaultQuery("limit", "10"), 10, 10)
	orders, total_count, err := h.OrderService.FindOrders(preOrderID, passenger, page, limit)
	paginationDta := orderDto.ResponsePaginationOrder{
		Data:       orders,
		TotalCount: total_count,
		Page:       page,
		PageSize:   limit,
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 400, "message": err.Error(), "data": paginationDta})
		log.Logger.Error("request for '/api/v1/orders' interface failed, reason: %s", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "查询成功", "data": paginationDta})
}
