// Package middleware
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     apollo.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/04/15 11:24:12
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package apollo

import (
	"smartCloudBean/internal/common/log"

	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/agcache"
	"github.com/apolloconfig/agollo/v4/env/config"
)

var ApolloCache agcache.CacheInterface

func init() {
	c := &config.AppConfig{
		AppID:          "org-system-smart-cloud-bean",
		Cluster:        "PRO",
		IP:             "http://192.168.3.232:8080",
		NamespaceName:  "application",
		IsBackupConfig: true,
		Secret:         "12436cfd7f054bb1a5366e3d228287f4",
	}

	client, _ := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	log.Logger.Info("初始化Apollo配置成功.")
	ApolloCache = client.GetConfigCache(c.NamespaceName)
}
