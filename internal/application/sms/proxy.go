// Package sms
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     proxy.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/06/28 00:50:41
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package sms

import (
	"regexp"
	smsRepo "smartCloudBean/internal/domain/repository/sms"
)

type SmsProxyService struct {
	SmsProxyRepository smsRepo.SmsProxyRepository
	SmsRepository      smsRepo.SmsRepository
}

func NewSmsProxyService() *SmsProxyService {
	return &SmsProxyService{
		SmsProxyRepository: smsRepo.NewSmsProxyRepository(),
		SmsRepository:      smsRepo.NewSmsRepository(),
	}
}

func (s *SmsProxyService) ParseSms(text string) (string, string, string) {
	// 定义正则表达式以匹配第一个和第二个\n之间的内容, 为短信内容
	re1 := regexp.MustCompile(`^[^\n]+\n(.*?)\n`)
	match1 := re1.FindStringSubmatch(text)
	var context string
	var phone string
	var code string
	if match1 == nil || len(match1) < 2 {
		return "", "", ""
	}
	context = match1[1]

	// 定义正则表达式以匹配手机号码
	// re2 := regexp.MustCompile(`\n[^_]*_(\d+)\n`)
	re2 := regexp.MustCompile(`SIM\d*_(\d{11})`)
	match2 := re2.FindStringSubmatch(text)
	if match2 == nil || len(match2) < 1 {
		phone = ""
	} else {
		phone = match2[1]
	}

	// 定义正则表达式以匹配验证码
	re3 := regexp.MustCompile(`验证码\D*(\d+)`)
	match3 := re3.FindStringSubmatch(text)
	if match3 == nil || len(match3) < 1 {
		code = ""
	} else {
		code = match3[1]
	}
	return context, phone, code
}
