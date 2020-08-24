package f2

import (
	"fmt"
	"testing"
)

func TestListOfIncome(t *testing.T) {
	Init()
	//user := "bitcoin/user"
	user := "nervos/ckb07p"
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
