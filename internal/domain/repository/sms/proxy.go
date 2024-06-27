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
	"io/ioutil"
	"net/http"
	"net/url"
	"smartCloudBean/internal/common/log"
	"smartCloudBean/internal/domain/repository/redis"
	"smartCloudBean/internal/infrastructure/http_client"
	"time"
)

type SmsProxyRepository interface {
	GetToken(username, password string) (string, error)
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

func (s *smsProxyRepository) GetToken(username, password string) (string, error) {
	tokenKey := "jlx:" + username + ":token"

	// 尝试从 Redis 获取 token
	token, err := s.redisClient.Get(tokenKey)
	if err == nil {
		log.Logger.Info("Token retrieved from Redis: %s", token)
		return token, nil
	}

	log.Logger.Info("Token not found in Redis, requesting new token with username: %s", username)

	urlAddr := "http://ticket.jiulvxing.com/ticket/auto/getApiUserToken"
	form := url.Values{}
	form.Set("userName", username)
	form.Set("password", password)
	formData := []byte(form.Encode())

	response, err := s.httpClient.Post(urlAddr, "application/x-www-form-urlencoded", formData, nil)
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
	url := "http://ticket.jiulvxing.com/purchaseApi/sms"
	requestData := map[string]string{"context": context, "phomeNum": phoneNum}
	jsonData, _ := json.Marshal(requestData)

	headers := map[string]string{"Authorization": token}

	response, err := s.httpClient.Post(url, "application/json", jsonData, headers)
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
