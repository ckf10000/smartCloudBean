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
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {

	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("你好，我的世界！")
	})
	s.Run()
}
