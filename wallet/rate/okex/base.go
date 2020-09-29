package okex

import (
	"fmt"
	"github.com/RMIO2020/Go-Utils/wallet/rate"
)

const (
	Host = "https://api.huobi.pro"

	MarketTickersUrl = "/market/tickers"
)

func GetMarketTickers() {
	url := Host + MarketTickersUrl

	data, err := rate.Request(rate.GET, url, rate.Params{})
	if err != nil {
		fmt.Println("err is ", err)
	}

	fmt.Println("data is ", data)
}
