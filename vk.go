package vkapi

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

// Response todo:
type VkResponse struct {
	Json string
}

// Client todo: init patterns with default params
type Client struct {
	Http       *http.Client
	BaseURL    string
	Retry      bool // should we auto retry?
	RetryCount int  // what times should we retry
}

type Params map[string]string

func (c *Client) Post(ctx context.Context, url string, queryParams, body Params) (*http.Response, error) {

	var req *http.Request
	var err error

	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reader := bytes.NewReader(data)

		req, err = http.NewRequestWithContext(ctx, "POST", c.BaseURL+url, reader)
		if err != nil {
			return nil, err
		}

	} else {
		req, err = http.NewRequestWithContext(ctx, "POST", c.BaseURL+url, nil)
		if err != nil {
			return nil, err
		}
	}

	if queryParams != nil {
		q := req.URL.Query()
		for k, v := range queryParams {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	req.Header.Set("Content-Type", "application/x-www-form-encoded")

	resp, err := c.Http.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) Get(ctx context.Context, url string, headers, queryParams Params) (*http.Response, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", c.BaseURL+url, nil)

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if queryParams != nil {
		q := req.URL.Query()
		for k, v := range queryParams {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.Http.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type VkAPI struct {
	Client  *Client
	Token   string
	Version string
}

// Call todo: return VkResponse?
//
// Generic method for calling VK API with specified query parameters
// https://dev.vk.com/ru/api/api-requests
func (vk *VkAPI) Call(ctx context.Context, method string, params Params) (*http.Response, error) {

	headers := Params{
		"Authorization": "Bearer " + vk.Token,
	}

	params["v"] = vk.Version

	// todo: Если вы передаёте большие данные (больше 2 Кбайт), используйте запрос POST и формат multipart/form-data.
	resp, err := vk.Client.Get(ctx, method, headers, params)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// todo: return VkResponse
// SendMessage https://dev.vk.com/ru/method/messages.send
func (vk *VkAPI) SendMessage(ctx context.Context, params Params) (*http.Response, error) {

	resp, err := vk.Call(ctx, "messages.send/", params)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
