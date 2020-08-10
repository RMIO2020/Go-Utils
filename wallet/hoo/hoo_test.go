package hoo

import (
	"fmt"
	"github.com/RMIO2020/Go-Common/helper"
	"github.com/RMIO2020/Go-Wallet-Service/config"
	spm_wallet "github.com/RMIO2020/Go-Wallet-Service/link/spm-wallet"
	"github.com/RMIO2020/Go-Wallet-Service/logger"
	"go.uber.org/zap"
	"testing"
)

var H *Hoo

func initDb() {
	fmt.Println("init db")
	vHandle := config.New()
	path, _ := helper.GetProjectRoot()
	path = path + "/../../../"
	err := vHandle.InitConfig(path)
	fmt.Printf(" Config  %+v \n", vHandle)
	if err != nil {
		panic("init config failed:" + err.Error())
		return
	}
	// init logger
	if err := logger.InitLogger(&vHandle.Config.Log); err != nil {
		logger.Error("init logger failed", zap.String("error", err.Error()))
		panic("init logger failed:" + err.Error())
		return
	}

	// init MySQL
	if err := spm_wallet.InitMySQL(&vHandle.Config.SpmWallet); err != nil {
		panic("init mysql failed:" + err.Error())
		return
	}

	H = NewHoo(vHandle.Config.Hoo.ClientId, vHandle.Config.Hoo.ClientSecret, vHandle.Config.Hoo.HOST, vHandle.Config.Hoo.ApiHost)
}

func TestGetAccount(t *testing.T) {
	initDb()
	p := &AccountWhere{
		CoinName: "BTC",
	}
	fmt.Printf("params is %+v \n", p)
	Account, err := H.GetAccount(p)
	fmt.Println("err", err)
	fmt.Println("Account TokenName ", Account.TokenName)
	fmt.Println("Account Balance ", Account.Balance)
	fmt.Println("Account Frozen ", Account.Frozen)
}

func TestNewAddress(t *testing.T) {
	initDb()
	p := &AddressWhere{
		CoinName: "ZEC",
		Num:      1,
	}
	fmt.Printf("params is %+v \n", p)
	Address, err := H.NewAddress(p)
	fmt.Println("err", err)
	fmt.Println("Address ", Address)
}

func TestWithdraw(t *testing.T) {
	initDb()
	p := &WithdrawWhere{
		CoinName:        "USDT",
		TokenName:       "USDT-ERC20",
		OrderId:         "sup-min26934277t",
		Amount:          "0.01",
		ToAddress:       "0x00f49805af7c2e9c18f024b4cb3c94a49a1ff2d1",
		ContractAddress: "0xdAC17F958D2ee523a2206206994597C13D831ec7",
	}
	err, hooResp, hooData := H.SetWithdraw(p)
	fmt.Println("err", err)
	fmt.Println("Address ", hooResp)
	fmt.Println("Address ", hooData)
}

func TestGetTickersMarket(t *testing.T) {
	initDb()

	err, abc := H.GetTickersMarket()

	fmt.Println("err ", err)
	fmt.Printf("abc %+v", abc)
}
