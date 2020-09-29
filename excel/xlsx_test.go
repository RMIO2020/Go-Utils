package excel

import (
	"github.com/RMIO2020/GO-PIN/config"
	"testing"
)

func TestListOfIncome(t *testing.T) {
	config.Config = &config.AppConfig{}
	config.Config.File = config.File{
		Domain:       "http://127.0.0.1:8091/static",
		AbsolutePath: "E:/project/file-service/static/excel",
		RelativePath: "",
	}
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
