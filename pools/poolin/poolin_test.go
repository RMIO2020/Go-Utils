package poolin

import (
	"fmt"
	"testing"
)

func initConf(Puid, Token string) {
	Init(Puid, Token)
}

func TestPoolIn_ListOfIncome(t *testing.T) {
	initConf("9123974", "wowA8GUkHvzF5kHKHQv6k3SExPu1nhBL8iEwfxNQGxaTaNZnLFikSR64aAUQaNm5")
	result, err := PIN.ListOfIncome(&IncomeListParams{
		CoinType:  "btc",
		Page:      0,
		StartDate: 20200701,
		EndDate:   20200831,
	})
	fmt.Printf("err is %+v \n", err)
	//fmt.Printf("result is %+v  \n", result.Data)
	for _, v := range result {
		fmt.Printf("v is %+v  \n", v)
	}
}
