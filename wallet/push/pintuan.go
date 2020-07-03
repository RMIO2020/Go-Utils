package push

import (
	"github.com/RMIO2020/Go-Wallet-Service/common/helper/request"
	"strconv"
)

const (
	PayStatusUrl = "/pin/order/notice"
)

var PinC *PinTuan

type PinTuan struct {
	Protocol string
	Host     string
}

func NewPinTuan() *PinTuan {
	return PinC
}

func InitPinTuan(Protocol, Host string) *PinTuan {
	PinC = &PinTuan{
		Protocol: Protocol,
		Host:     Host,
	}
	return PinC
}

type ToPinTuan struct {
	OrderNo     string  `json:"order_no"`
	PayCurrency string  `json:"pay_currency"`
	PayProtocol string  `json:"pay_protocol"`
	PayType     string  `json:"pay_type"`
	PayStatus   string  `json:"pay_status"`
	PayAmount   float64 `json:"pay_amount"`
}

func (P *PinTuan) GetPinTuanParams(Params *ToPinTuan) request.ReqParams {
	var params = request.ReqParams{}
	params["order_no"] = Params.OrderNo
	params["pay_currency"] = Params.PayCurrency
	params["pay_protocol"] = Params.PayProtocol
	params["pay_type"] = Params.PayType
	params["pay_status"] = Params.PayStatus
	params["pay_amount"] = strconv.FormatFloat(Params.PayAmount, 'G', -1, 64)
	return params
}

func (P *PinTuan) PushPayStatus(Params *ToPinTuan) (result string, err error) {
	params := P.GetPinTuanParams(Params)
	url := P.Protocol + P.Host + PayStatusUrl
	result, err = request.Request(request.PUT, url, params, request.ContentTypFormUrl)
	result = result + " | " + url
	return
}
