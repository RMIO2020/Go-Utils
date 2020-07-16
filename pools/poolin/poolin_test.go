package poolin

import (
	"fmt"
	"testing"
)

func initConf(Puid, Token string) {
	Init(Puid, Token)
}

func TestPoolIn_ListOfIncome(t *testing.T) {
	initConf("9034431", "wowCMVVgBKKP12QwfJ1weMS9NwpxkWeR0yaWZdcenHbQlPXSApDrEhgdUhmvsqDh")
	result, err := PIN.ListOfIncome(&IncomeListParams{
		Puid:      "9034431",
		CoinType:  "zec",
		Page:      0,
		StartDate: 20200714,
		EndDate:   20200716,
	})
	fmt.Printf("err is %+v \n", err)
	//fmt.Printf("result is %+v  \n", result.Data)
	for _, v := range result.Data {
		fmt.Printf("v is %+v  \n", v)
	}
}
