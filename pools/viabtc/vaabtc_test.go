package viabtc

import (
	"fmt"
	"testing"
)

func TestListOfIncome(t *testing.T) {
	Init()
	user := "07dfb1b2eab13248d4d1d3764f42d0c1"
	List, err := PIN.ListOfIncome(user)
	if err != nil {
		fmt.Println("err", err.Error())
	}
	if len(List) > 0 {
		for k, v := range List {
			fmt.Println(k, ":", v)
		}
	}
}
