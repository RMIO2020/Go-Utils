package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/RMIO2020/Go-User-Service/logger"
	"go.uber.org/zap"
	"io/ioutil"
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

type clientV1 struct {
	httpClient *http.Client
	url        string
	in         interface{}
	pattern    string
	header     map[string]string
}

// NewClientV1 实例化一个客户端
func NewClientV1(url string, in interface{}, pattern string, header map[string]string) *clientV1 {
	return &clientV1{
		httpClient: &http.Client{},
		url:        url,
		in:         in,
		pattern:    pattern,
		header:     header,
	}
}

// GoResponse go响应
type GoResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"message"`
}

// GoJavaResponse java响应
type GoJavaResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"message"`
}

// CurlGet 通用的curl的get请求
func (c *clientV1) CurlGet(query map[string]string) (out []byte, err error) {
	req, err := http.NewRequest(c.pattern, c.url, nil)
	if err != nil {
		return
	}
	if len(query) > 0 {
		q := req.URL.Query()
		for k, v := range query {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}
	for k, v := range c.header {
		req.Header.Set(k, v)
	}
	rsp, err := c.httpClient.Do(req)
	if err != nil {
		return
	}
	body := rsp.Body
	defer body.Close()

	if rsp.StatusCode != http.StatusOK {
		tmp, err1 := ioutil.ReadAll(rsp.Body)
		logger.Warn("curl get:", zap.String("url", c.url), zap.String("pattern", c.pattern), zap.String("error message", string(tmp)))
		if err1 != nil {
			err = errors.New("remote: status " + string(rsp.StatusCode) + "http status is error")
			return
		}
		data := ToGo(tmp)
		if len(data.Msg) > 0 {
			err = errors.New("remote: " + data.Msg)
		} else {
			err = errors.New("remote: status " + string(rsp.StatusCode) + "http status is error")
		}
		return
	}
	out, err = ioutil.ReadAll(rsp.Body)
	return
}

// Curl 通用的curl请求
func (c *clientV1) Curl() (out []byte, err error) {
	inJson, _ := json.Marshal(c.in)
	req, err := http.NewRequest(c.pattern, c.url, bytes.NewBuffer(inJson))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	for k, v := range c.header {
		req.Header.Set(k, v)
	}
	rsp, err := c.httpClient.Do(req)
	if err != nil {
		return
	}
	body := rsp.Body
	defer body.Close()
	fmt.Println(c.pattern, c.url)
	out, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		return
	}
	fmt.Println(c.pattern, c.url, string(out))
	if rsp.StatusCode != http.StatusOK {
		data := ToGo(out)
		if len(data.Msg) > 0 {
			err = errors.New("remote: " + data.Msg)
		} else {
			//err = errors.New("remote: status " + string(rsp.StatusCode) + " http status is error")
			err = errors.New(fmt.Sprintf("remote: status %d http status is error", rsp.StatusCode))
		}
	}
	return
}

// ToRm 转成rm结构体
func ToRm(body []byte) (out *RmResponse) {
	out = &RmResponse{}
	_ = json.Unmarshal(body, out)
	return
}

// ToRm 转成go结构体
func ToGo(body []byte) (out *GoResponse) {
	out = &GoResponse{}
	_ = json.Unmarshal(body, out)
	return
}

// To20Java 转成钱包2.0java结构体
func To20Java(body []byte) (out *GoJavaResponse) {
	out = &GoJavaResponse{}
	_ = json.Unmarshal(body, out)
	return
}
