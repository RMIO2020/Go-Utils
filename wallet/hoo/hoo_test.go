package hoo

import (
	"fmt"
	"testing"
)

const (
	ClientId     = "UvBeD3EamTaGLM6fx32ZESyGZJnLzN"
	ClientSecret = "Cibs1MK4paoM3ArkkQ4TS5h6n6WcadWZqPJmsBcx38BZQoDhFJ"
	Host         = "https://hoo.com"
	HotAddress   = "0x8aAeDDf30A3Fd059574FDc8Bc1936072b182A1C4"
)

var H = NewHoo(ClientId, ClientSecret, Host)

func TestGetAccount(t *testing.T) {
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
	p := &AddressWhere{
		CoinName: "ETH",
		Num:      1,
	}
	fmt.Printf("params is %+v \n", p)
	Address, err := H.NewAddress(p)
	fmt.Println("err", err)
	fmt.Println("Address ", Address)
}

func TestWithdraw(t *testing.T) {
	p := &WithdrawWhere{
		CoinName:  "ETH",
		TokenName: "ETH",
		OrderId:   "Test003",
		Amount:    "0.001",
		ToAddress: HotAddress,
	}
	Address, err := H.SetWithdraw(p)
	fmt.Println("err", err)
	fmt.Println("Address ", Address)
}
