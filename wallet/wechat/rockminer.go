package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/RMIO2020/Go-Wallet-Service/common/helper/request"
	"github.com/RMIO2020/Go-Wallet-Service/common/helper/rsa"
	"github.com/yuchenfw/gocrypt"
	"strconv"
)

const (
	RMProtocol = "https://"
	RMHOST     = "api.rockminer.com"
	RMPayUrl   = "/app/pay/confirm-pay"
)

type ThirdPayToRM struct {
	OrderSn     string  `json:"order_sn"`
	Token       string  `json:"token"`
	PaymentType int64   `json:"payment_type"`
	PayAmount   float64 `json:"pay_amount"`
	OrderType   int64   `json:"order_type"`
	Lang        string  `json:"lang"`
	Client      int64   `json:"client"`
}

type BaseReq struct {
	Data string
	Sign string
}

type ReqData struct {
	RetCode int
	RetData string
	RetMsg  string
}

func GetPayUrl(Params *ThirdPayToRM, RSA *rsa.Crypt) (PayUrl string, err error) {
	params := getParams(Params)
	RepParams, err := getRepParams(params, RSA)
	if err != nil {
		return
	}
	url := RMProtocol + RMHOST + RMPayUrl
	sorted := request.SortParams(RepParams)
	PayUrl = url + "?" + sorted
	return
}

func RMPay(Params *ThirdPayToRM, RSA *rsa.Crypt) (result string, err error) {
	params := getParams(Params)

	RepParams, err := getRepParams(params, RSA)
	if err != nil {
		return
	}

	url := RMProtocol + RMHOST + RMPayUrl
	result, err = request.Request(request.GET, url, RepParams)
	fmt.Printf("result %+v \n", result)
	if err != nil {
		return
	}
	result, err = checkBase(result, RSA)
	return
}

func getParams(Params *ThirdPayToRM) request.ReqParams {
	var params = request.ReqParams{}
	params["order_sn"] = Params.OrderSn
	params["token"] = Params.Token
	params["payment_type"] = strconv.FormatInt(Params.PaymentType, 10)
	params["pay_amount"] = strconv.FormatFloat(Params.PayAmount, 'G', -1, 64)
	params["order_type"] = strconv.FormatInt(Params.OrderType, 10)
	params["lang"] = Params.Lang
	params["client"] = strconv.FormatInt(Params.Client, 10)

	return params
}

func getRepParams(params request.ReqParams, RSA *rsa.Crypt) (request.ReqParams, error) {
	var RsaParams = request.ReqParams{}
	data, _ := json.Marshal(params)

	sign, err := RSA.Sign(string(data), gocrypt.SHA256, gocrypt.Base64)
	if err != nil {
		return nil, err
	}
	RsaParams["sign"] = sign

	data2, err := RSA.Encrypt(string(data), gocrypt.Base64)
	if err != nil {
		return nil, err
	}
	RsaParams["data"] = data2
	return RsaParams, nil
}

func checkBase(Message string, RSA *rsa.Crypt) (ReData string, err error) {
	var data BaseReq
	err = json.Unmarshal([]byte(Message), &data)
	if err != nil {
		return
	}

	data2, err := RSA.Decrypt(data.Data, gocrypt.Base64)
	if err != nil {
		return
	}

	verifySign, err := RSA.VerifySign(data2, gocrypt.SHA256, data.Sign, gocrypt.Base64)
	if err != nil {
		return
	}

	if !verifySign {
		err = errors.New("verifySign err")
		return
	}

	var data3 ReqData
	err = json.Unmarshal([]byte(data2), &data3)
	if err != nil {
		return
	}
	if data3.RetCode != 0 {
		err = errors.New(data3.RetMsg)
		return
	}
	ReData = data3.RetData
	return
}
