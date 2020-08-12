package hoo

import (
	"encoding/json"
)

const (
	TickersMarketUrl = "/open/v1/tickers/market"
)

type TickersMarketResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    []TickerMarket
}

type TickerMarket struct {
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

func (H *Hoo) GetTickersMarket() (err error, data TickersMarketResp) {
	hooResp, err := H.Request(GET, TickersMarketUrl, map[string]string{
		"ApiUrl":     "true",
		"NullParams": "true",
	})
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(hooResp), &data)
	return
}
