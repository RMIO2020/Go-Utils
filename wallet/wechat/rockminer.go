package wechat

import (
	"github.com/RMIO2020/Go-Wallet-Service/common/helper/request"
	"strconv"
)

type ThirdPayToRM struct {
	OrderSn     string  `json:"order_sn"`
	UserId      string  `json:"user_id"`
	PaymentType int64   `json:"payment_type"`
	PayAmount   float64 `json:"pay_amount"`
	OrderType   int64   `json:"order_type"`
	Lang        string  `json:"lang"`
	Client      int64   `json:"client"`
}

func RMPay(Params *ThirdPayToRM) {
	var params = request.ReqParams{}

	params["order_sn"] = Params.OrderSn
	params["user_id"] = Params.UserId
	params["payment_type"] = strconv.FormatInt(Params.PaymentType, 10)
	params["pay_amount"] = strconv.FormatFloat(Params.PayAmount, 'G', -1, 64)
	params["order_type"] = strconv.FormatInt(Params.OrderType, 10)
	params["lang"] = Params.Lang
	params["client"] = strconv.FormatInt(Params.Client, 10)
}
