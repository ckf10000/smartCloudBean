// Package sms
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     sms.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/06/27 21:08:36
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package sms

import (
	"fmt"
	"smartCloudBean/internal/common/log"
	smsDto "smartCloudBean/internal/domain/dto/sms"
	"smartCloudBean/internal/infrastructure/middleware/mysql"

	"github.com/jinzhu/gorm"
)

type SmsRepository interface {
	FindSms(PhoneNum string, searchValue string, page int, limit int) ([]smsDto.ResponseSms, int, error)
	InsertSms(sms *smsDto.SmsInsert) error
}

type smsRepository struct {
	db *gorm.DB
}

func NewSmsRepository() SmsRepository {
	return &smsRepository{db: mysql.Db}
}

func (r *smsRepository) InsertSms(sms *smsDto.SmsInsert) error {
	if err := r.db.Table("t_sms").Create(sms).Error; err != nil {
		return fmt.Errorf("failed to insert sms: %v", err)
	}
	return nil
}

func (r *smsRepository) FindSms(phoneNum string, searchValue string, page int, limit int) ([]smsDto.ResponseSms, int, error) {
	var sms []smsDto.ResponseSms
	var totalRows int64

	query := r.db.Table("t_sms").Select("id, phone_num, context, digital_verification_code, DATE_FORMAT(create_time, '%Y-%m-%d %H:%i:%s') AS create_time")
	if phoneNum != "" {
		query = query.Where("phone_num = ?", phoneNum)
	}

	if searchValue != "" {
		query = query.Where("context like ?", "%"+searchValue+"%")
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
	// sql := query.Offset(offset).Limit(limit).Order("create_time DESC").Debug().Find(&sms)

	// 打印 SQL 查询语句
	// fmt.Printf("Executing SQL: %s\n", sql.Debug().Get("sql"))

	err := query.Offset(offset).Limit(limit).Order("create_time DESC").Scan(&sms).Error
	if err != nil {
		log.Logger.Error("failed to fetch orders: %v", err)
		return nil, 0, err
	}

	return sms, int(totalRows), nil
}
