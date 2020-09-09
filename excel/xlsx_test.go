package excel

import (
	"testing"
)

func TestListOfIncome(t *testing.T) {
	title := "测试"
	field := []string{
		"栏目名1",
		"栏目名2",
		"栏目名3",
		"栏目名4",
		"栏目名5",
	}
	data := [][]interface{}{
		{"1", "2", "3", "4", 5},
		{"6", "7", "8", "9", 10},
		{"11", "12", "13", "14", 15},
	}
	saveName, err := WriteExcel(title, field, data)
	if err != nil {
		t.Error("fail!")
	}
	t.Log("success! saveName is ", saveName)
}
