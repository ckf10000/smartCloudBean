// Package log
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     log.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/06/27 19:14:55
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package log

import (
	"github.com/ckf10000/gologger/v3/log"
)

var logFile string = "web.log"
var logLevel string = "debug"
var logMessageTemplate string = "simple" // simple|standard

// 定义全局变量
var Logger *log.FileLogger

func get_logger() *log.FileLogger {
	// 获取可执行文件所在目录的路径
	exeDir := log.GetExecuteFilePath()
	if exeDir == "" {
		return nil
	}
	// projectPath := "./"
	log := log.NewLogger(logLevel, exeDir, logFile, logMessageTemplate, 50*1024*1024, true, true, true)
	return log
}

// init 函数在包初始化时运行
func init() {
	Logger = get_logger()
	if Logger == nil {
		Logger.Fatal("Logger initialization failed")
	}
}
