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
		UserId:      "$2y$13$nchMil0i.HsMEI92UVR3ruqUJCC7F3N6FKq4RvKec6ivbk6KodMEa",
		PaymentType: 11,
		PayAmount:   0.01,
		OrderType:   2,
		Lang:        "en",
		Client:      2,
	}

	result, err := RMPay(params, RSA)
	fmt.Printf("err is %+v \n", err)
	fmt.Printf("result is %+v \n", result)
}

func Test_Decrypt(t *testing.T) {
	Init()
	RSA := rsa.NewRSACrypt()
	data := "{\"data\":\"3MepupOK0mfy3SxzAeqApHgkTY7Yddymv45NEBIL7CBTBy36DzwntiHMT+8VPBxfBDjEPP1H2yadUovjPQFNjFVghi1sXF95PZvxUOW3COJec3bn3BZmicURlSH6GRxSuMfUs2ouxFt6J31+sREzK9NDe2ELaE8Mr4V1g6jnDFCbmLLKOJJnPpAB04SMD9NFAYD7ZictIwCvH8jTzrJHtq+lSR1PokRtHj9yrZWnRKxlvk\\/ST7lGDg3ZuqDCSlTkyRW+MWBBmXKmD487RXnOHvJP7LKdJCJk8Z3YFdS8+qanbmsHgbWGCjhRjN5gF2vrLe1lOAbxk43kqkY4nVXRwBTiYmIYqJvhvdgf9uhhcXdaydwRLokjetRE1wpaiTT1PPd3KWxT9T5UoJ1rXiq8KNAg+Z8DLTnjeWtht3vfGQUeZqIeGNrZRnF675NwO1GFsl8Y5pJ7xuV\\/CNYDROKTlvc3WoOcUChy4vfXgfd0NFtk4oIYl+kop0LfcIGqL7DUtX5JktwjNCpbu2lXMJsk5vJo4W\\/gAS7wkDQilDm88F\\/f4KoS5B1OX+du3+GPRDwok6pVlN4qeZzta\\/Za8W5JOWLoR9L8+WCZwK\\/BtU5cKJMHUA7VKKAa9XvBe0f8eLjkU6Nonp\\/IsK58vPq2Q3aYPeZaVvSHapEANGGfMvvGyQ0=\",\"sign\":\"c9pDHPsOwzaxc\\/ilOsWjA6LbWu8rggoO5R9NIrSqBlG3PSD\\/U2mdRMbmL55jjxIl8DcA\\/PccrlO0t0gHF+AV9tNYQsCgY8wTa5HPxTQbOZW5q4iROjHG5S89xHbGhGSa728i5nJmZ9CSbqYpg3UDpyyUgtDQ1+qtjNzoo4QcZb8vAN0IlTSPRUHCkKnXzbJoy2UNO51qYgNw\\/Y8xI66FHZNIeatxpMeZonHGSVLiWKKbwegLwWkO+ZhHTnu7q3a6znr\\/\\/uzGt\\/AWq+5SZvz7INmUyNNf3mR\\/oKW9Kvj\\/qji3FkKfXZ1wcVrZ+G64EhkdYiFXbl1azs0+9NwTSP49VyonQdfwTMX9+YoHzoKZ3U8tb94+WM8eIyWIrw1qQb\\/oAanHBfzwwHfyexFWhSvgnose94re2YHpRFOzgr4SfMo2QdaDOs6YFsJtNpuoGlMU4\\/FMNFq8smeDxfG2ittV\\/5CwsNnfA5dtpFnEd8k3dG3P9bHLqzm741c7djVl7Cya6lKxPvWz2MwNcJbC+5Ptsn+rZozChb9cF+sZEX9MD1nUJqT8jAmllRfZR9Ao5sMRR7\\/P3J9xmYcIXSPT7igVZw8BlnsdHonmZZohOGriC0diAkM+xgAKmCxw7GpQlzqT90j1fo9B3K0HqsIixXGrUMO04G+k3mrrpF4kR\\/j2gHk=\"}"

	result, err := checkBase(data, RSA)
	fmt.Printf("err is %+v \n", err)
	fmt.Printf("result is %+v \n", result)
}
