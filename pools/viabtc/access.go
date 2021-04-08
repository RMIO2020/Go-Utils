package viabtc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

const (
	GET = "GET"
)

type Params map[string]string

type BaseReq struct {
	Errors  string      `json:"errors"`
	ErrNo   int64       `json:"err_no"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func CheckBase(Message string, result interface{}) (err error) {
	fmt.Printf("data is %+v \n", Message)
	_ = json.Unmarshal([]byte(Message), &result)
	return
}

func (P *F2) SortParams(params map[string]string) string {
	keys := make([]string, len(params))
	i := 0
	for k := range params {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	sorted := make([]string, len(params))
	i = 0
	for _, k := range keys {
		sorted[i] = k + "=" + url.QueryEscape(params[k])
		i++
	}
	return strings.Join(sorted, "&")
}

func (P *F2) Request(method string, path string, params map[string]string) (result string, err error) {
	client := &http.Client{}
	sorted := P.SortParams(params)
	var req *http.Request
	if method == "POST" {
		req, _ = http.NewRequest(method, P.HOST+path, strings.NewReader(sorted))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req, _ = http.NewRequest(method, P.HOST+path+"?"+sorted, strings.NewReader(""))
	}
	fmt.Println("F2 Client Do......")
	resp, err := client.Do(req)
	fmt.Println("F2 resp is", resp)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	result = string(body)
	return
}
