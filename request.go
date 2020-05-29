package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type client struct {
	httpClient *http.Client
	url        string
	in         interface{}
}

// NewClient 实例化一个客户端
func NewClient(url string, in interface{}) *client {
	return &client{
		httpClient: &http.Client{},
		url:        url,
		in:         in,
	}
}

// Response rm响应
type RmResponse struct {
	RetCode int         `json:"retCode"`
	RetData interface{} `json:"retData"`
	RetMsg  string      `json:"retMsg"`
}

// PostCurl post请求
func (c *client) PostCurl() (out *RmResponse, err error) {
	out = &RmResponse{}
	inJson, _ := json.Marshal(c.in)
	req, err := http.NewRequest(http.MethodPost, c.url, bytes.NewBuffer(inJson))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	rsp, err := c.httpClient.Do(req)
	if err != nil {
		return
	}
	body := rsp.Body
	defer body.Close()

	if rsp.StatusCode != http.StatusOK {
		err = errors.New("http status is error")
		return
	}
	err = json.NewDecoder(rsp.Body).Decode(out)
	return
}
