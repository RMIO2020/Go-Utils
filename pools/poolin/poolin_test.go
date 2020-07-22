package poolin

import (
	"fmt"
	"testing"
)

func initConf(Puid, Token string) {
	Init(Puid, Token)
}

func TestPoolIn_ListOfIncome(t *testing.T) {
	initConf("9114639", "wowzEqrkJOW5AJNoo9Aj72NyATX3ve9Ye85A0oVG1q9hLYJDQsSgec0M69d59gSK")
	result, err := PIN.ListOfIncome(&IncomeListParams{
		CoinType:  "zec",
		Page:      0,
		StartDate: 20200720,
		EndDate:   20200720,
	})
	fmt.Printf("err is %+v \n", err)
	//fmt.Printf("result is %+v  \n", result.Data)
	for _, v := range result {
		fmt.Printf("v is %+v  \n", v)
	}
}
