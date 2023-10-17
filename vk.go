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

type Body map[string]string

func (c *Client) Post(ctx context.Context, url string, body Body) (*http.Response, error) {

	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(data)

	req, err := http.NewRequestWithContext(ctx, "POST", url, reader)

	resp, err := c.Http.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
