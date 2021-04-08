package viabtc

const ApiUrl = "https://www.viabtc.com"

var PIN *Viabtc

type Viabtc struct {
	HOST string `json:"host"`
}

func New() *Viabtc {
	return PIN
}

func Init() *Viabtc {
	PIN = &Viabtc{
		HOST: ApiUrl,
	}
	return PIN
}
