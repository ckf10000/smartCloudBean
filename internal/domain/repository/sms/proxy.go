// Package sms
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     proxy.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/06/28 00:57:13
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package sms

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"smartCloudBean/internal/common/log"
	"smartCloudBean/internal/domain/repository/redis"
	"smartCloudBean/internal/infrastructure/http_client"
	"smartCloudBean/internal/infrastructure/middleware/apollo"
	"time"
)

type SmsProxyRepository interface {
	GetToken() (string, error)
	SendSms(token, context, phoneNum string) error
}

type smsProxyRepository struct {
	httpClient  http_client.HttpClient
	redisClient redis.RedisClient
}

func NewSmsProxyRepository() SmsProxyRepository {
	return &smsProxyRepository{
		httpClient:  http_client.NewHttpClient(),
		redisClient: redis.NewRedisClient(),
	}
}

func (s *smsProxyRepository) GetToken() (string, error) {
	user, _ := apollo.ApolloCache.Get("jlx.api.auth_user")
	password, _ := apollo.ApolloCache.Get("jlx.api.auth_password")
	urlAddr, _ := apollo.ApolloCache.Get("jlx.api.get_token_url")

	// 构建 Redis URI
	tokenKey := fmt.Sprintf("jlx:%s:token", user)

	// 尝试从 Redis 获取 token
	token, err := s.redisClient.Get(tokenKey)
	if err == nil {
		log.Logger.Info("Token retrieved from Redis: %s", token)
		return token, nil
	}

	log.Logger.Info("Token not found in Redis, requesting new token with username: %s", user)

	form := url.Values{}
	form.Set("userName", user.(string))
	form.Set("password", password.(string))
	formData := []byte(form.Encode())

	response, err := s.httpClient.Post(urlAddr.(string), "application/x-www-form-urlencoded", formData, nil)
	if err != nil {
		log.Logger.Error("Error getting token: %v", err)
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Logger.Error("Failed to get token, status code: %d", response.StatusCode)
		return "", errors.New("failed to get token")
	}

	var result struct {
		Msg     string `json:"msg"`
		Code    int    `json:"code"`
		Success bool   `json:"success"`
		Data    struct {
			Token      string `json:"token"`
			UserName   string `json:"userName"`
			LoginTime  int64  `json:"loginTime"`
			ExpireTime int64  `json:"expireTime"`
		} `json:"data"`
	}

	body, _ := ioutil.ReadAll(response.Body)
	if err := json.Unmarshal(body, &result); err != nil {
		log.Logger.Error("Error parsing token response: %v", err)
		return "", err
	}

	if !result.Success {
		log.Logger.Info("Token request unsuccessful, Reason is: %s", result.Msg)
		return "", errors.New("request not successful")
	}

	token = result.Data.Token
	// 将 token 存储到 Redis 并设置有效期为 12 小时
	if err := s.redisClient.Set(tokenKey, token, 12*time.Hour); err != nil {
		log.Logger.Error("Error setting token in Redis: %v", err)
		return "", err
	}

	log.Logger.Info("Successfully obtained and stored token: %s", token)
	return token, nil
}

func (s *smsProxyRepository) SendSms(token, context, phoneNum string) error {
	urlAddr, _ := apollo.ApolloCache.Get("jlx.api.send_sms")
	requestData := map[string]string{"context": context, "phomeNum": phoneNum}
	jsonData, _ := json.Marshal(requestData)

	headers := map[string]string{"Authorization": token}

	response, err := s.httpClient.Post(urlAddr.(string), "application/json", jsonData, headers)
	if err != nil {
		log.Logger.Error("Error sending SMS: %v", err)
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Logger.Error("Failed to send SMS, status code: %d", response.StatusCode)
		return errors.New("failed to send sms")
	}

	var result struct {
		Code    int    `json:"code"`
		Success bool   `json:"success"`
		Msg     string `json:"msg"`
	}

	body, _ := ioutil.ReadAll(response.Body)
	if err := json.Unmarshal(body, &result); err != nil {
		log.Logger.Error("Error parsing SMS response: %v", err)
		return err
	}

	if !result.Success {
		log.Logger.Warning("SMS request unsuccessful, Reason is: %s", result.Msg)
		return errors.New("request not successful")
	}

	// common.Logger.Info("Successfully sent SMS, SMS context: %s, Send Phone: %s", context, phoneNum)
	return nil
}
