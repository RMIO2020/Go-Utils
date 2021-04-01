package poolin

import (
	"github.com/RMIO2020/Go-Common/helper/request"
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
	Data         []IncomeDetails
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

type IncomeDetails struct {
	Date             int
	CoinType         string `mapstructure:"coin_type"`
	Amount           float64
	Change           float64
	ShareAccept      float64
	ShareUnit        string
	RejectRate       float64
	Address          string
	WalletName       string
	WalletNameEn     string
	AddressCreatedAt float64
	Txhash           string
	UnpaidReason     string
	PaymentStatus    string
	PaidAt           float64
	CreatedAt        float64
	Rate             string
	Coinexchange     string
	BillType         float64
}

func (P *PoolIn) ListOfIncome(Params *IncomeListParams) (List []IncomeDetails, err error) {
	params := request.ReqParams{}

	if Params.Puid == "" {
		params["puid"] = P.Puid
	}
	params["coin_type"] = Params.CoinType
	params["page"] = strconv.Itoa(Params.Page)
	params["start_date"] = strconv.Itoa(Params.StartDate)
	params["end_date"] = strconv.Itoa(Params.EndDate)

	result, err := P.Request(GET, HistoryUrl, params)
	if err != nil {
		return
	}
	IncomeList := IncomeListRet{}
	err = CheckBase(result, &IncomeList)
	List = IncomeList.Data
	return
}
