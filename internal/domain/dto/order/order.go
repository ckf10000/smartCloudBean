// Package order
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     order.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/04/15 10:00:30
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package order

type ResponseOrder struct {
	ID               int     `json:"id"`
	PreOrderID       int     `json:"pre_order_id"`
	PreSaleAmount    float64 `json:"pre_sale_amount"`
	Flight           string  `json:"flight"`
	Passenger        string  `json:"passenger"`
	AgeStage         string  `json:"age_stage"`
	CtripOrderID     string  `json:"ctrip_order_id"`
	PaymentAmount    float64 `json:"payment_amount"`
	PaymentMethod    string  `json:"payment_method"`
	ItineraryID      string  `json:"itinerary_id"`
	CtripUsername    string  `json:"ctrip_username"`
	UserPass         string  `json:"user_pass"`
	OutPF            string  `json:"out_pf"`
	OutTicketAccount string  `json:"out_ticket_account"`
	PayAccountType   string  `json:"pay_account_type"`
	PayAccount       string  `json:"pay_account"`
	PaymentTime      string  `json:"payment_time"`
}

// Response 结构体用于表示接收 JSON 响应的数据结构
type ResponsePaginationOrder struct {
	Data       []ResponseOrder `json:"data"`        // 订单数据列表
	TotalCount int             `json:"total_count"` // 数据总数
	Page       int             `json:"page"`        // 当前页码
	PageSize   int             `json:"page_size"`   // 每页数据条数
}
