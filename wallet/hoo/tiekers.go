package hoo

import (
	"encoding/json"
)

const (
	TickersUrl = "/open/v1/tickers/market"
)

type TickerReq struct {
	Code    int
	Message string
	Data    []TickerData
}

type TickerData struct {
	Amount string
	AmtNum int64
	Change string
	High   string
	Low    string
	Price  string
	QtyNum int64
	Symbol string
	Volume string
}

func (H *Hoo) GetTickers() (BaseData TickerReq, err error) {
	hooResp, err := H.Request(GET, TickersUrl, map[string]string{
		"is_api_host": "true",
		"null_params": "true",
	})
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(hooResp), &BaseData)
	return
}
