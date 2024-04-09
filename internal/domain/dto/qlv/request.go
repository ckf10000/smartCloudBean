// Package qlv
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     qlv_request.go
* Description:  劲旅平台api请求结构体
* Author:       ckf10000
* CreateDate:   2024/04/09 16:15:58
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package qlv

type CommonRequest struct {
	UserID  string `json:"UserId"`  // 合作伙伴用户编号
	SignKey string `json:"SignKey"` // 签名key
}

// LockOrderRequest 请求结构体
type LockOrderRequest struct {
	CommonRequest
	RequestData LockOrderRequestData `json:"requestData"`
}

// LockOrderRequestData 请求数据结构体
type LockOrderRequestData struct {
	PolicyName  string `json:"policyName"`  // 政策名称关键字
	OrderSrcCat string `json:"orderSrcCat"` //订单类别，国内、国际
	Oper        string `json:"oper"`        // 机器人名字,如果在系统中配置了机器人,则会根据进单分配的机器人名称匹配订单,如果在系统中没有配置则会更具政策关键字,航司等条件锁单
}

// UnlockOrderRequest 请求结构体
type UnlockOrderRequest struct {
	CommonRequest
	RequestData UnlockOrderRequestData `json:"requestData"`
}

// UnlockOrderRequestData 请求数据结构体
type UnlockOrderRequestData struct {
	OrderID       string `json:"orderID"`       // 劲旅订单号
	Oper          string `json:"oper"`          // 操作人,机器人名称
	OrderState    string `json:"orderState"`    // 出票状态，0失败，1成功
	OrderLoseType string `json:"orderLoseType"` // 出票失败类别，用户自定义失败类别,比如:政策,系统等等
}

// OrderLogWriteNew 请求数据结构体
type OrderLogWriteNewRequest struct {
	CommonRequest
	RequestData OrderLogWriteNewRequestData `json:"requestData"`
}

// OrderLogWriteNewRequestData 请求数据结构体
type OrderLogWriteNewRequestData struct {
	OrderID string `json:"orderID"` // 劲旅订单号
	Oper    string `json:"oper"`    // 操作人,机器人名称
	Logs    string `json:"logs"`    // 日志内容
}

// SaveOrderPayInfo 请求数据结构体
type SaveOrderPayInfoRequest struct {
	CommonRequest
	RequestData SaveOrderPayInfoRequestData `json:"requestData"`
}

// OrderLogWriteNewRequestData 请求数据结构体
type SaveOrderPayInfoRequestData struct {
}
