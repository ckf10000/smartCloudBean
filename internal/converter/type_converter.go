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

import (
	"strconv"
	"time"
)

// parseQueryInt 解析 URL 查询参数为整数
func ConvertQueryInt(value string, nullDefault, exceptionDefault int) int {
	if value == "" {
		return nullDefault
	}
	result, err := strconv.Atoi(value)
	if err != nil {
		return exceptionDefault
	}
	return result
}

// RFC3339格式（例如：2024-04-15T09:12:53+08:00）时间字符串转本地本地时间字符串(2016-01-02 15:04:02)
func ConvertRFC3339StrToLocalStr(dateTimeStr string) string {
	// 指定日期时间字符串的格式
	layout := "2006-01-02T15:04:05-07:00"
	// 使用 Parse 函数将字符串解析为 time.Time 类型
	t, err := time.Parse(layout, dateTimeStr)
	if err != nil {
		return "Error parsing time"
	}
	// 使用 Format 方法将时间转换为指定格式的字符串
	timeStr := t.Format("2006-01-02 15:04:02")
	return timeStr
}
