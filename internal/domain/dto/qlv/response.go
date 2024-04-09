// Package qlv
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     qlv_response.go
* Description:  劲旅平台api响应结构体
* Author:       ckf10000
* CreateDate:   2024/04/09 16:15:31
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package qlv

// LockOrderResponse 是从服务端返回的订单信息结构体
type LockOrderResponse struct {
	Code         int           `json:"Code"`
	Message      string        `json:"Message"`
	DataInfoJson *OrderDetails `json:"DataInfoJson"`
}

// OrderDetails 包含订单详细信息的结构体
type OrderDetails struct {
	ID            int         `json:"ID"`
	IOrderNo      string      `json:"IOrderNo"`
	TripType      string      `json:"TripType"`
	StatOrder     string      `json:"StatOrder"`
	StatOpration  string      `json:"StatOpration"`
	ReceiptStat   string      `json:"ReceiptStat"`
	Receipted     float64     `json:"Receipted"`
	PaySum        float64     `json:"PaySum"`
	SourceType    string      `json:"SourceType"`
	SourceCat     string      `json:"SourceCat"`
	SourceOTA     string      `json:"SourceOTA"`
	SourceName    string      `json:"SourceName"`
	DatCreate     string      `json:"DatCreate"`
	Contact       string      `json:"Contact"`
	ContactMobile string      `json:"ContactMobile"`
	ContactEmail  string      `json:"ContactEmail"`
	OrderRemark   string      `json:"OrderRemark"`
	PolicyTypeOTA string      `json:"PolicyTypeOTA"`
	PolicyID      string      `json:"PolicyID"`
	PolicyCode    string      `json:"PolicyCode"`
	PolicyName    string      `json:"PolicyName"`
	PolicyType    string      `json:"PolicyType"`
	RuleChng      string      `json:"RuleChng"`
	IssueRemark   string      `json:"IssueRemark"`
	Flights       []Flight    `json:"Flights"`
	Peoples       []Passenger `json:"Peoples"`
}

// Flight 是航班信息结构体
type Flight struct {
	OrderID   int     `json:"OrderID"`
	FlightIdx int     `json:"FlightIdx"`
	AirCoCode string  `json:"AirCoCode"`
	CodeDep   string  `json:"CodeDep"`
	CodeArr   string  `json:"CodeArr"`
	DatArr    string  `json:"DatArr"`
	DatDep    string  `json:"DatDep"`
	FlightNo  string  `json:"FlightNo"`
	PlaneType string  `json:"PlaneType"`
	Cabin     string  `json:"Cabin"`
	PriceStd  float64 `json:"PriceStd"`
	CityDep   string  `json:"CityDep"`
	CityArr   string  `json:"CityArr"`
	AirName   string  `json:"AirName"`
}

// Passenger 是乘客信息结构体
type Passenger struct {
	OrderID  int    `json:"OrderID"`
	PName    string `json:"PName"`
	PrName   string `json:"PrName"`
	PType    string `json:"PType"`
	IDType   string `json:"IDType"`
	IDNo     string `json:"IDNo"`
	Gender   string `json:"Gender"`
	BirthDay string `json:"BirthDay"`
	Mobile   string `json:"Mobile"`
}
