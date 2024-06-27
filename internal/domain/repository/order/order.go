// Package order
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     order.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/04/15 10:14:39
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package order

import (
	orderDto "smartCloudBean/internal/domain/dto/order"

	"smartCloudBean/internal/common/log"
	"smartCloudBean/internal/infrastructure/middleware/mysql"

	"github.com/jinzhu/gorm"
)

type OrderRepository interface {
	FindOrders(preOrderID int, passenger string, page int, limit int) ([]orderDto.ResponseOrder, int, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository() OrderRepository {
	return &orderRepository{db: mysql.Db}
}

func (r *orderRepository) FindOrders(preOrderID int, passenger string, page int, limit int) ([]orderDto.ResponseOrder, int, error) {
	var orders []orderDto.ResponseOrder
	var totalRows int64

	query := r.db.Table("t_orders").Select("id, pre_order_id, pre_sale_amount, flight, passenger, age_stage, ctrip_order_id, payment_amount, payment_method, itinerary_id, ctrip_username, user_pass, out_pf, out_ticket_account, pay_account_type, pay_account, DATE_FORMAT(payment_time, '%Y-%m-%d %H:%i:%s') AS payment_time")
	if preOrderID != -1 {
		query = query.Where("pre_order_id = ?", preOrderID)
	}
	if passenger != "" {
		query = query.Where("passenger = ?", passenger)
	}

	// 查询满足条件的总数量
	if err := query.Count(&totalRows).Error; err != nil {
		log.Logger.Error("Error counting rows: %v\n", err)
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
		log.Logger.Error("failed to fetch orders: %v", err)
		return nil, 0, err
	}

	return orders, int(totalRows), nil
}
