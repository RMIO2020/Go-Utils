package f2

import "testing"

func TestListOfIncome(t *testing.T) {
	Init()
	user := "bitcoin/user"
	//user := "bitcoin/rmt19p07"
	_, _ = PIN.ListOfIncome(user)
}
