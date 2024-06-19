// Package services
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     order_service.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/04/15 10:00:25
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package services

import (
	"smartCloudBean/internal/domain/dto"
	"smartCloudBean/internal/infrastructure/repository"

	"github.com/ckf10000/gologger/v3/log"
	"github.com/jinzhu/gorm"
)

type OrderService struct {
	orderRepository repository.OrderRepository
}

type SmsService struct {
	smsRepository repository.SmsRepository
}

func NewOrderService(db *gorm.DB, log *log.FileLogger) *OrderService {
	return &OrderService{
		orderRepository: repository.NewOrderRepository(db, log),
	}
}

func NewSmsService(db *gorm.DB, log *log.FileLogger) *SmsService {
	return &SmsService{
		smsRepository: repository.NewSmsRepository(db, log),
	}
}

func (s *OrderService) FindOrders(preOrderID int, passenger string, page, limit int, log *log.FileLogger) ([]dto.ResponseOrder, int, error) {
	return s.orderRepository.FindOrders(preOrderID, passenger, page, limit, log)
}

func (s *SmsService) FindSms(phoneNum string, page, limit int, log *log.FileLogger) ([]dto.ResponseSms, int, error) {
	return s.smsRepository.FindSms(phoneNum, page, limit, log)
}
