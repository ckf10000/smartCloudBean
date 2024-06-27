// Package order
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     order.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/04/15 10:00:25
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package order

import (
	orderDto "smartCloudBean/internal/domain/dto/order"
	orderRepo "smartCloudBean/internal/domain/repository/order"
)

type OrderService struct {
	orderRepository orderRepo.OrderRepository
}

func NewOrderService() *OrderService {
	return &OrderService{
		orderRepository: orderRepo.NewOrderRepository(),
	}
}

func (s *OrderService) FindOrders(preOrderID int, passenger string, page, limit int) ([]orderDto.ResponseOrder, int, error) {
	return s.orderRepository.FindOrders(preOrderID, passenger, page, limit)
}
