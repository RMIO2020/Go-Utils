package excel

import (
	"fmt"
	"github.com/RMIO2020/GO-PIN/config"
	utils "github.com/RMIO2020/Go-Utils"
	"github.com/tealeg/xlsx"
)

func WriteExcel(title string, field []string, data [][]interface{}) (saveName string, err error) {
	//创建表格
	file := xlsx.NewFile()
	//设置页面名称
	s, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 设置标题
	titleStyle := xlsx.NewStyle()
	titleStyle.Font.Size = 18
	titleStyle.Font.Bold = true
	titleStyle.Alignment.Horizontal = "center"
	titleStyle.Alignment.Vertical = "center"
	titleName := s.AddRow()
	titleName.SetHeightCM(1.5)
	cell := titleName.AddCell()
	cell.Value = title
	cell.Merge(len(field)-1, 0)
	cell.SetStyle(titleStyle)

	// 设置栏目名
	columnStyle := xlsx.NewStyle()
	columnStyle.Alignment.Horizontal = "center"
	columnStyle.Alignment.Vertical = "center"
	column := s.AddRow()
	column.SetHeightCM(1.5)
	for _, v := range field {
		cell := column.AddCell()
		cell.SetStyle(columnStyle)
		cell.Value = v
	}
	//设置内容
	for _, v1 := range data {
		row := s.AddRow()
		for _, v2 := range v1 {
			cell := row.AddCell()
			cell.SetValue(v2)
		}
	}
	fileName := title + "-" + utils.CreateOrderSn() + ".xlsx"
	filePath := config.Config.File.AbsolutePath + "/" + config.Config.File.RelativePath + "/" + fileName
	err = file.Save(filePath)
	saveName = config.Config.File.Domain + "/" + config.Config.File.RelativePath + "/" + fileName

	return
}
