// Package service
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     qlv_service.go
* Description:  获取劲旅平台相关数据
* Author:       ckf10000
* CreateDate:   2024/04/09 16:27:17
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package service

import (
	"fmt"
	"net/url"
	"smartCloudBean/internal/domain/dto/qlv"
)

// LockOrder 发送 LockOrder 请求并解析响应
func (s *HTTPService) LockOrder(requestData qlv.LockOrderRequest) (interface{}, error) {
	path := s.baseURL + "/LockOrder.ashx"
	params := url.Values{}
	headers := map[string]string{}

	var result interface{}
	err := s.Post(path, params, headers, requestData, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to lock order: %w", err)
	}

	return result, nil
}
