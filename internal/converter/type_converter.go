// Package converter
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     type_converter.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/04/15 12:07:08
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package converter

import "strconv"

// parseQueryInt 解析 URL 查询参数为整数
func ConvertQueryInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	result, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return result
}
