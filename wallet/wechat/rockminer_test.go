package wechat

import (
	"fmt"
	"github.com/RMIO2020/Go-Common/helper"
	"github.com/RMIO2020/Go-Wallet-Service/common/helper/rsa"
	"github.com/RMIO2020/Go-Wallet-Service/config"
	"testing"
)

func Init() {
	vHandle := config.New()
	path, _ := helper.GetProjectRoot()
	path = path + "/../../../"
	err := vHandle.InitConfig(path)
	if err != nil {
		panic(err)
	}
	rsa.InitTRSACrypt(vHandle.Config.RMRsa)
}

func TestRMPay(t *testing.T) {
	params := &ThirdPayToRM{
		Platform:    "pin-min",
		OrderSn:     "test2",
		Token:       "eN8B8ZZ2xZfLUi8NGlDSeDoGtM4pVOGd",
		PaymentType: 11,
		PayAmount:   1,
		OrderType:   1,
		Lang:        "en",
		Client:      1,
		ProcessType: 3,
	}

	result, err := GetPayUrl(params)
	fmt.Printf("err is %+v \n", err)
	fmt.Printf("result is %+v \n", result)
}
