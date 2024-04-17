// Package repository
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     order_repository.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/04/15 10:14:39
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package repository

import (
	"smartCloudBean/internal/domain/dto"

	"github.com/ckf10000/gologger/v3/log"
	"github.com/jinzhu/gorm"
)

type OrderRepository interface {
	FindOrders(preOrderID int, passenger string, page int, limit int, log *log.FileLogger) ([]dto.ResponseOrder, int, error)
}

type orderRepository struct {
	db  *gorm.DB
	log *log.FileLogger
}

func NewOrderRepository(db *gorm.DB, log *log.FileLogger) OrderRepository {
	return &orderRepository{db, log}
}

func (r *orderRepository) FindOrders(preOrderID int, passenger string, page int, limit int, log *log.FileLogger) ([]dto.ResponseOrder, int, error) {
	var orders []dto.ResponseOrder
	var totalRows int64

	query := r.db.Table("t_orders").Select("id, pre_order_id, pre_sale_amount, flight, passenger, age_stage, ctrip_order_id, payment_amount, payment_method, itinerary_id, ctrip_username, user_pass, out_pf, out_ticket_account, pay_account_type, pay_account, DATE_FORMAT(create_time, '%Y-%m-%d %H:%i:%s') AS payment_time")
	if preOrderID != -1 {
		query = query.Where("pre_order_id = ?", preOrderID)
	}
	if passenger != "" {
		query = query.Where("passenger = ?", passenger)
	}

	// 查询满足条件的总数量
	if err := query.Count(&totalRows).Error; err != nil {
		log.Error("Error counting rows: %v\n", err)
	}

	offset := (page - 1) * limit
	if offset < 0 {
		offset = 0
	}

	// 获取 SQL 查询语句
	// sql := query.Offset(offset).Limit(limit).Order("create_time DESC").Debug().Find(&orders)

	// 打印 SQL 查询语句
	// fmt.Printf("Executing SQL: %s\n", sql.Debug().Get("sql"))

	err := query.Offset(offset).Limit(limit).Order("create_time DESC").Scan(&orders).Error
	if err != nil {
		log.Error("failed to fetch orders: %v", err)
		return nil, 0, err
	}

	return orders, int(totalRows), nil
}
