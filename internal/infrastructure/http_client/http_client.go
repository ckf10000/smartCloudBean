// Package http_client
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     http_client.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/06/27 20:59:55
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package http_client

import (
	"bytes"
	"net/http"
	"smartCloudBean/internal/common/log"
)

type HttpClient interface {
	Post(url string, contentType string, body []byte, headers map[string]string) (*http.Response, error)
}

type httpClient struct{}

func NewHttpClient() HttpClient {
	return &httpClient{}
}

func (c *httpClient) Post(url string, contentType string, body []byte, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Logger.Error("Error creating request: %v", err)
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	log.Logger.Info("Sending POST request to %s with body: %s and headers: %v", url, string(body), req.Header)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Logger.Error("Error sending request: %v", err)
		return nil, err
	}
	log.Logger.Info("Received response with status code: %d", response.StatusCode)
	return response, nil
}
