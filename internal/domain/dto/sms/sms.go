// Package sms
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     sms.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/06/27 21:02:12
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package sms

// SMSInsert 表示入库短信对象
type SmsInsert struct {
	PhoneNum                string `json:"phone_num"`
	Context                 string `json:"context"`
	DigitalVerificationCode string `json:"digital_verification_code"`
}

func (SmsInsert) TableName() string {
	return "t_sms"
}

type ResponseSms struct {
	ID                      int    `json:"id"`
	PhoneNum                int    `json:"phone_num"`
	Context                 string `json:"context"`
	DigitalVerificationCode string `json:"digital_verification_code"`
	CreateTime              string `json:"create_time"`
}

// Response 结构体用于表示接收 JSON 响应的数据结构
type ResponsePaginationSms struct {
	Data       []ResponseSms `json:"data"`        // 短信数据列表
	TotalCount int           `json:"total_count"` // 数据总数
	Page       int           `json:"page"`        // 当前页码
	PageSize   int           `json:"page_size"`   // 每页数据条数
}
