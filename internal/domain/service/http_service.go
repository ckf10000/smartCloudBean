// Package service
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     http_service.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/04/09 15:25:06
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package service

import (
	"net/url"
	"smartCloudBean/internal/infrastructure/api"
)

// HTTPService is an HTTP request service
type HTTPService struct {
	httpClient *api.HTTPClient
	baseURL    string
}

// NewHTTPService creates a new HTTP request service with the specified baseURL
func NewHTTPService(baseURL string) *HTTPService {
	return &HTTPService{
		httpClient: api.NewHTTPClient(baseURL),
		baseURL:    baseURL,
	}
}

// Get sends a GET request and parses the JSON response
func (s *HTTPService) Get(path string, params url.Values, headers map[string]string, result interface{}) error {
	return s.httpClient.Get(path, params, headers, result)
}

// Post sends a POST request and parses the JSON response
func (s *HTTPService) Post(path string, params url.Values, headers map[string]string, body interface{}, result interface{}) error {
	return s.httpClient.Post(path, params, headers, body, result)
}
