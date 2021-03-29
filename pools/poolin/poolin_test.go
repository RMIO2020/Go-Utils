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
	initConf("9114639", "wowapZmxqTiyHBTokTjn4NcjJ1AYIlYp52DLeN3ifHEghrPZ7sJutWCsI73U8RtU")
	result, err := PIN.ListOfIncome(&IncomeListParams{
		CoinType: "zec",
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
