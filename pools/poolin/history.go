package poolin

import (
	"github.com/RMIO2020/Go-Wallet-Service/common/helper/request"
	"strconv"
)

const (
	HistoryUrl = "/api/public/v2/payment/history"
)

type IncomeListParams struct {
	Puid      string `json:"puid"`
	CoinType  string `json:"coin_type"`
	Page      int    `json:"page"`
	StartDate int    `json:"start_date"`
	EndDate   int    `json:"end_date"`
}

type IncomeListRet struct {
	CurrentPage  int64
	Data         interface{}
	FirstPageUrl string
	From         int64
	LastPage     int64
	LastPageUrl  string
	NextPageUrl  string
	Path         string
	PerPage      int64
	PrevPageUrl  string
	To           int64
	Total        int64
}

func (P *PoolIn) ListOfIncome(Params *IncomeListParams) (IncomeList IncomeListRet, err error) {
	params := request.ReqParams{}

	params["puid"] = Params.Puid
	params["coin_type"] = Params.CoinType
	params["Page"] = strconv.Itoa(Params.Page)
	params["start_date"] = strconv.Itoa(Params.StartDate)
	params["end_date"] = strconv.Itoa(Params.EndDate)

	result, err := P.Request(GET, HistoryUrl, params)
	if err != nil {
		return
	}
	IncomeList = IncomeListRet{}
	err = CheckBase(result, &IncomeList)
	return
}
