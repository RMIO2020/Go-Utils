package wechat

import (
	"fmt"
	"github.com/RMIO2020/Go-Common/helper"
	"github.com/RMIO2020/Go-Wallet-Service/common/helper/rsa"
	"github.com/RMIO2020/Go-Wallet-Service/config"
	"testing"
)

func Init() {
	vHandle := config.New()
	path, _ := helper.GetProjectRoot()
	path = path + "/../../../"
	err := vHandle.InitConfig(path)
	if err != nil {
		panic(err)
	}
	rsa.InitTRSACrypt(vHandle.Config.RMRsa)
}

func TestRMPay(t *testing.T) {
	Init()
	RSA := rsa.NewRSACrypt()
	params := &ThirdPayToRM{
		OrderSn:     "test2",
		UserId:      "$2y$13$W6xCSp72rHldAqIXLbuyFuLkVny/q/NRLpWOQGQgtUm.SIRxuAiOO",
		PaymentType: 11,
		PayAmount:   1,
		OrderType:   2,
		Lang:        "en",
		Client:      1,
	}

	result, err := RMPay(params, RSA)
	fmt.Printf("err is %+v \n", err)
	fmt.Printf("result is %+v \n", result)
}

func Test_Decrypt(t *testing.T) {
	Init()
	RSA := rsa.NewRSACrypt()
	data := "{\"data\":\"RcWZiHi5tXi96V3a6fajykWZZ56YscOJLEYpcQ08+XPpRQEIMNA8g9xhsH1bGlBgmvvPlyHWJSE+kM0ZFPHpvUEd6aoOVu4vc4vA72TCfYgM+jeNGwFEutM+jjhOuivIly1d7Ko74xQS3LM\\/gY2yPkr4pWX8jrFTznb+oxlRpa3FgItXolbw0lxp7NnEGfp7jzWIM8Z0GuOK8LcJkMgvCVMYHXqGYlhjQJyfuvUDpnkPz3hIrlsoMYKza0j5IEXBB1e06588yQiD1N8dimEIPDl5p8MaOikLgLvJO5gSNtZwcmrQq1i3VlzjLFdKJjEGR7\\/JVk9U5\\/GuHMT1KYeNr5kpDgcGmvF2DDsWtVGrPHccCI9r+Gx8lBpNSSGnNPtkFdB798kzYkmgGno9FUJXNnrNEs7oUTaM2TPeHkzbhC2OvmrpX1EaIGhKGb0Zn7xcr+t0D+UVAYMRR4\\/FN6yB+4A+ocBWpxdurn+9BzI8aCmarzrIUo8rt8dB+h2LXyfaDuzf4Cc3WyplQ8ydou0ftOCk5UilnNoxnR\\/N1keuCwHH6kTSds\\/AtyeXq4R8EuqdNV5EXDO2zdtJcUNA9dJoiX3JWBemtP8MgklPkVGVK5KFgpIkeye+Tw5WsGbh40oVJ3xrwTp9W0J7DxWHJiQk97rNqQZ04FJ5HTz0VWa9ijI=\",\"sign\":\"c9pDHPsOwzaxc\\/ilOsWjA6LbWu8rggoO5R9NIrSqBlG3PSD\\/U2mdRMbmL55jjxIl8DcA\\/PccrlO0t0gHF+AV9tNYQsCgY8wTa5HPxTQbOZW5q4iROjHG5S89xHbGhGSa728i5nJmZ9CSbqYpg3UDpyyUgtDQ1+qtjNzoo4QcZb8vAN0IlTSPRUHCkKnXzbJoy2UNO51qYgNw\\/Y8xI66FHZNIeatxpMeZonHGSVLiWKKbwegLwWkO+ZhHTnu7q3a6znr\\/\\/uzGt\\/AWq+5SZvz7INmUyNNf3mR\\/oKW9Kvj\\/qji3FkKfXZ1wcVrZ+G64EhkdYiFXbl1azs0+9NwTSP49VyonQdfwTMX9+YoHzoKZ3U8tb94+WM8eIyWIrw1qQb\\/oAanHBfzwwHfyexFWhSvgnose94re2YHpRFOzgr4SfMo2QdaDOs6YFsJtNpuoGlMU4\\/FMNFq8smeDxfG2ittV\\/5CwsNnfA5dtpFnEd8k3dG3P9bHLqzm741c7djVl7Cya6lKxPvWz2MwNcJbC+5Ptsn+rZozChb9cF+sZEX9MD1nUJqT8jAmllRfZR9Ao5sMRR7\\/P3J9xmYcIXSPT7igVZw8BlnsdHonmZZohOGriC0diAkM+xgAKmCxw7GpQlzqT90j1fo9B3K0HqsIixXGrUMO04G+k3mrrpF4kR\\/j2gHk=\"}"

	result, err := checkBase(data, RSA)
	fmt.Printf("err is %+v \n", err)
	fmt.Printf("result is %+v \n", result)
}
