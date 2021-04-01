package poolin

import (
	"fmt"
	"testing"
)

func initConf(Puid, Token string) {
	Init(Puid, Token)
}

func TestPoolIn_ListOfIncome(t *testing.T) {
	//initConf("9153598", "wowrAsFD2hBxdOWwyDqpNEq0MWa8jKPl2skfAv0Svtm3x4nL7MLG2Aqns6UeoiVU")
	initConf("9199959", "wowJUxTE5nbzFlrhLNEYascBvLOiyDqazQQ96UesNpUr0t1jugzzoLS8rCWAMLDf")
	result, err := PIN.ListOfIncome(&IncomeListParams{
		CoinType: "eth",
		//Page:      0,
		//StartDate: ,
		//EndDate:   ,
	})
	fmt.Printf("err is %+v \n", err)
	//fmt.Printf("result is %+v  \n", result.Data)
	for _, v := range result {
		fmt.Printf("v is %+v  \n", v)
	}
}
