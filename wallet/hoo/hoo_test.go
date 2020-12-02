package hoo

import (
	"fmt"
	"github.com/RMIO2020/Go-Utils/wallet/hoo"
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

func TestSignEcc(t *testing.T) {
	s := hoo.NewHoo("UvBeD3EamTaGLM6fx32ZESyGZJnLzN", "Cibs1MK4paoM3ArkkQ4TS5h6n6WcadWZqPJmsBcx38BZQoDhFJ", "https://www.hoozh.com", "https://api.hoolgd.com")
	Params := map[string]string{
		"sign":           "96bd4770e9b66fafc39fb11077baefee9f25805926e216ad1eac3e7db7b26d0e",
		"chain_name":     "ETH",
		"coin_name":      "USDT-ERC20",
		"alias":          "USDT",
		"trad_type":      "2",
		"block_height":   "11364984",
		"transaction_id": "0x8b957da4b462fdd2febe0021773cafab2548aea69be967e11573a0f50df95c14",
		"trx_n":          "0", "confirmations": "6",
		"from_address": "0x008d44a22df16518ea360dacdea5fb8cf6717a9e",
		"to_address":   "0x00a4d4a5ddc78860cdce8e051f2f094fe388a406",
		"memo":         "", "amount": "2825.782881",
		"fee": "0", "contract_address": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
		"outer_order_no": "USDT-ERC20_00890502",
		"confirm_time":   "1606806616",
		"message":        "success", "Status": "",
		"Currency": "",
		"Protocol": "",
		"WalletId": "",
		"Uid":      "",
		"Platform": "",
	}
	str := s.SignEcc(Params)
	fmt.Println(str)
}
