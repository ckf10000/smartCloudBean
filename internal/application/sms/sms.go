// Package sms
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     sms.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/06/28 00:44:48
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package sms

import (
	smsDto "smartCloudBean/internal/domain/dto/sms"
	smsRepo "smartCloudBean/internal/domain/repository/sms"
)

type SmsService struct {
	smsRepository smsRepo.SmsRepository
}

func NewSmsService() *SmsService {
	return &SmsService{
		smsRepository: smsRepo.NewSmsRepository(),
	}
}

func (s *SmsService) InsertSms(sms *smsDto.SmsInsert) error {
	return s.smsRepository.InsertSms(sms)
}

func (s *SmsService) FindSms(phoneNum string, searchValue string, page, limit int) ([]smsDto.ResponseSms, int, error) {
	return s.smsRepository.FindSms(phoneNum, searchValue, page, limit)
}
