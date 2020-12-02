package hoo

import (
	"fmt"
	"testing"
)

var H *Hoo

func initDb() {
	fmt.Println("init db")
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
