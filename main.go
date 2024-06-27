// Package main
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     main.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/04/02 10:43:04
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package main

import (
	"smartCloudBean/internal/infrastructure/middleware/mysql"
	"smartCloudBean/internal/web"
)

func main() {

	defer mysql.Db.Close()
	web.StartServer()
}
