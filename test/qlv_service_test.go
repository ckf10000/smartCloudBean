// Package test
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     qlv_service_test.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/04/09 18:17:04
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package test

import (
	"fmt"
	"smartCloudBean/internal/domain/dto/qlv"
	"smartCloudBean/internal/domain/service"
	"testing"
)

func TestQlvLockOrder(t *testing.T) {
	httpService := service.NewHTTPService("https://pekdazhicheng.qlv88.com")
	// 构造请求参数
	requestData := qlv.LockOrderRequest{
		CommonRequest: qlv.CommonRequest{
			UserID:  "186",
			SignKey: "9a68295ec90b1fc10ab94331c882bab9",
		},
		RequestData: qlv.LockOrderRequestData{
			PolicyName:  "测试-XC",
			OrderSrcCat: "国内",
			Oper:        "测试-XC",
		},
	}
	response, err := httpService.LockOrder(requestData)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Lock order response:", response)
}
