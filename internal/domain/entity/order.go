// Package entity
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     order_entity.go
* Description:  订单模型
* Author:       ckf10000
* CreateDate:   2024/04/16 09:12:07
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package entity

import "time"

type OrderEntity struct {
	ID                int       `json:"id"`
	PreOrderID        int       `json:"pre_order_id"`
	DepartureCity     string    `json:"departure_city"`
	ArriveCity        string    `json:"arrive_city"`
	DepartureTime     time.Time `json:"departure_time"`
	PreSaleAmount     float64   `json:"pre_sale_amount"`
	Flight            string    `json:"flight"`
	Passenger         string    `json:"passenger"`
	AgeStage          string    `json:"age_stage"`
	CardType          string    `json:"card_type"`
	CardId            string    `json:"card_id"`
	InternalPhone     string    `json:"internal_phone"`
	PassengerPhone    string    `json:"passenger_phone"`
	CtripOrderID      string    `json:"ctrip_order_id"`
	PaymentAmount     float64   `json:"payment_amount"`
	PaymentMethod     string    `json:"payment_method"`
	ItineraryID       string    `json:"itinerary_id"`
	DepartureCityName string    `json:"departure_city_name"`
	ArriveCityName    string    `json:"arrive_city_name"`
	ArriveTime        time.Time `json:"arrive_time"`
	CtripUsername     string    `json:"ctrip_username"`
	UserPass          string    `json:"user_pass"`
	OutPF             string    `json:"out_pf"`
	OutTicketAccount  string    `json:"out_ticket_account"`
	PayAccountType    string    `json:"pay_account_type"`
	PayAccount        string    `json:"pay_account"`
	Oper              string    `json:"oper"`
	IsDeleted         int       `json:"is_deleted"`
	PaymentTime       time.Time `json:"payment_time"`
	CreateTime        time.Time `json:"create_time"`
	UpdateTime        time.Time `json:"update_time"`
}

// SMS 表示短信对象
type SMSEntity struct {
	ID                      int
	PhoneNum                string
	Context                 string
	DigitalVerificationCode string
	CreateTime              time.Time
}
