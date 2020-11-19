package f2

import (
	"regexp"
	"strings"
	"time"
)

type IncomeListRet struct {
	Balance                     float64            `json:"balance"`
	HashesLastDay               float64            `json:"hashes_last_day"`
	Value                       float64            `json:"value"`
	StaleHashesRejectedLastDay  float64            `json:"stale_hashes_rejected_last_day"`
	Workers                     [][]interface{}    `json:"workers"`
	ValueWorkers                map[string]float64 `json:"value_workers"`
	ValueLastDay                float64            `json:"value_last_day"`
	HashrateHistory             map[string]float64 `json:"hashrate_history"`
	StaleHashesRejectedLastHour float64            `json:"stale_hashes_rejected_last_hour"`
	Paid                        uint               `json:"paid"`
	HashesLastHour              float64            `json:"hashes_last_hour"`
	WorkerLengthOnline          float64            `json:"worker_length_online"`
	PayoutHistory               [][]interface{}    `json:"payout_history"`
	PayoutHistoryFee            [][]interface{}    `json:"payout_history_fee"`
	WorkerLength                int                `json:"worker_length"`
	Hashrate                    float64            `json:"hashrate"`
}

type PayoutHistory struct {
	Date   time.Time
	TxId   string
	Income float64
}

func (P *F2) ListOfIncome(user string) (List []*PayoutHistory, err error) {
	path := "/" + strings.TrimLeft(user, "/")
	result, err := P.Request(GET, path, nil)
	if err != nil {
		return
	}
	IncomeList := IncomeListRet{}
	err = CheckBase(result, &IncomeList)
	if err != nil {
		return
	}
	List = IncomeList.GetPayoutHistory()
	return
}

func (P *IncomeListRet) GetPayoutHistory() (PayoutHistoryData []*PayoutHistory) {
	PayoutHistoryRet := P.PayoutHistory
	PayoutHistoryFee := P.PayoutHistoryFee
	dataPattern := `\d{4}.\d{1,2}.\d{1,2}`
	reg := regexp.MustCompile(dataPattern)
	for k, v := range PayoutHistoryRet {
		if len(v) < 3 {
			// 若长度小于3则跳过
			continue
		}
		timeStr, ok := v[0].(string)
		if !ok {
			continue
		}
		matchDate := reg.FindString(timeStr)
		date, err := time.ParseInLocation("2006-01-02", matchDate, time.Local)
		if err != nil {
			continue
		}
		var (
			txid   string
			income float64
		)
		txid, _ = v[1].(string)
		income, _ = v[2].(float64)
		history := &PayoutHistory{
			Date:   date,
			TxId:   txid,
			Income: income,
		}
		if len(PayoutHistoryFee) > 0 {
			timeStr, ok = PayoutHistoryFee[k][0].(string)
			if !ok {
				continue
			}
			matchDate = reg.FindString(timeStr)
			date2, err := time.ParseInLocation("2006-01-02", matchDate, time.Local)
			if err != nil {
				continue
			}
			if date == date2 {
				history.Income += PayoutHistoryFee[k][2].(float64)
			}
		}

		PayoutHistoryData = append(PayoutHistoryData, history)
	}
	return
}
