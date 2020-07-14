package poolin

import (
	"fmt"
	"github.com/RMIO2020/Go-Common/helper"
	"github.com/RMIO2020/Go-Wallet-Service/config"
	"testing"
)

func initConf() {
	fmt.Println("init db")
	vHandle := config.New()
	path, _ := helper.GetProjectRoot()
	path = path + "/../../../"
	fmt.Println("conf", path)
	err := vHandle.InitConfig(path)
	if err != nil {
		panic("init config failed:" + err.Error())
		return
	}
	Init(vHandle.Config.PoolIn.Puid, vHandle.Config.PoolIn.Token)
}

func TestPoolIn_ListOfIncome(t *testing.T) {
	initConf()
	result, err := PIN.ListOfIncome(&IncomeListParams{
		Puid:      "9034431",
		CoinType:  "zec",
		Page:      1,
		StartDate: 20200701,
		EndDate:   20200708,
	})
	fmt.Printf("err is %+v \n", err)
	fmt.Printf("result is %+v  \n", result.Data)
}
