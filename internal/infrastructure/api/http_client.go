// Package api
/***********************************************************************************************************************
* ProjectName:  smartCloudBean
* FileName:     http_client.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/04/09 14:53:12
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// HTTPClient is a generic HTTP client
type HTTPClient struct {
	BaseURL string // Base URL
}

// NewHTTPClient creates a new HTTP client
func NewHTTPClient(baseURL string) *HTTPClient {
	return &HTTPClient{BaseURL: baseURL}
}

// Get sends a GET request and parses the JSON response
func (c *HTTPClient) Get(path string, params url.Values, headers map[string]string, result interface{}) error {
	reqURL := c.BaseURL + path
	if params != nil {
		reqURL += "?" + params.Encode()
	}

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("GET request <%s> failed with status code: %d", reqURL, resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(result)
}

// Post sends a POST request and parses the JSON response
func (c *HTTPClient) Post(path string, params url.Values, headers map[string]string, body interface{}, result interface{}) error {
	reqURL := c.BaseURL + path
	if params != nil {
		reqURL += "?" + params.Encode()
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("POST request <%s> failed with status code: %d", reqURL, resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(result)
}
