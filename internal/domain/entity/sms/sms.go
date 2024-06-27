// Package sms
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     sms.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/06/28 00:35:56
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package sms

import "time"

// SMS 表示短信对象
type Sms struct {
	ID                      int       `json:"id"`
	PhoneNum                string    `json:"phone_num"`
	Context                 string    `json:"context"`
	DigitalVerificationCode string    `json:"digital_verification_code"`
	IsDeleted               int       `json:"is_deleted"`
	CreateTime              time.Time `json:"create_time"`
	UpdateTime              time.Time `json:"update_time"`
	DeletedTime             int16     `json:"deleted_time"`
}

func (Sms) TableName() string {
	return "t_sms"
}
