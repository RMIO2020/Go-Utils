package hoo

import (
	"fmt"
	"testing"
)

const (
	ClientId     = "UvBeD3EamTaGLM6fx32ZESyGZJnLzN"
	ClientSecret = "Cibs1MK4paoM3ArkkQ4TS5h6n6WcadWZqPJmsBcx38BZQoDhFJ"
	Host         = "https://hoo.com"
	HotAddress   = "0x000bb4e5694e241a06e9f42c583205d073a046fc"
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
		OrderId:   "sup-min165074399724041972",
		Amount:    "0.001",
		//ToAddress: HotAddress,
		ToAddress: "0x000bb4e5694e241a06e9f42c583205d073a046fc",
	}
	err, hooResp, hooData := H.SetWithdraw(p)
	fmt.Println("err", err)
	fmt.Println("Address ", hooResp)
	fmt.Println("Address ", hooData)
}

func TestHoo_GetKline(t *testing.T) {
	where := &KlineWhere{
		Symbol: "BTC_USDT",
		Type:   "5Min",
	}

	Klin, err := H.GetKline(where)

	fmt.Println("err ", err)
	fmt.Println("Klin ", Klin)
}
