package viabtc

import (
	"encoding/json"
	"github.com/RMIO2020/GO-PIN/logger"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type IncomeRet struct {
	Code int `json:"code"`
	Data DataRet
}

type DataRet struct {
	Code     int             `json:"code"`
	CurrPage int             `json:"curr_page"`
	Data     []IncomeListRet `json:"Data"`
}

type IncomeListRet struct {
	Coin        string `json:"coin"`
	Date        string `json:"date"`
	PplnsProfit string `json:"pplns_profit"`
	PpsProfit   string `json:"pps_profit"`
	SoloProfit  string `json:"solo_profit"`
	TotalProfit string `json:"total_profit"`
}

type PayoutHistory struct {
	Date   time.Time
	TxId   string
	Income float64
}

func (P *Viabtc) ListOfIncome(xApiKey string, currency string) (List []*PayoutHistory, err error) {
	req, err := http.NewRequest(http.MethodGet, P.HOST+"/res/openapi/v1/profit/history", nil)
	if err != nil {
		return
	}
	q := req.URL.Query()
	q.Add("coin", currency)
	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return
	}
	body := rsp.Body
	defer body.Close()

	if rsp.StatusCode != http.StatusOK {
		tmp, err := ioutil.ReadAll(rsp.Body)
		logger.Warn("curl get viabtc data, fail:", zap.String("url", P.HOST+"/res/openapi/v1/profit/history"), zap.String("pattern", http.MethodGet), zap.String("error message", string(tmp)))
		if err != nil {
			return
		}
		return
	}
	out, err := ioutil.ReadAll(rsp.Body)
	ret := IncomeRet{}
	_ = json.Unmarshal(out, &ret)
	List = ret.GetPayoutHistory()
	return
}

func (P *IncomeRet) GetPayoutHistory() (PayoutHistoryData []*PayoutHistory) {
	if P.Code != 0 {
		return
	}
	if P.Data.Data == nil || len(P.Data.Data) == 0 {
		return
	}
	list := P.Data.Data
	for _, v := range list {
		date, err := time.ParseInLocation("2006-01-02", v.Date, time.Local)
		if err != nil {
			continue
		}
		var (
			income float64
		)
		income, _ = strconv.ParseFloat(v.TotalProfit, 64)
		txid := v.Date
		if income == 0 {
			continue
		}
		history := &PayoutHistory{
			Date:   date,
			TxId:   txid,
			Income: income,
		}
		PayoutHistoryData = append(PayoutHistoryData, history)
	}
	return
}
