package f2

import (
	"fmt"
	"testing"
)

func TestListOfIncome(t *testing.T) {
	Init()
	//user := "bitcoin/user"
	user := "ethereum/rmp07a10"
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
