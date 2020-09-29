package poolin

const ApiUrl = "https://api-prod.poolin.com"

var PIN *PoolIn

type PoolIn struct {
	Puid  string `json:"puid"`
	HOST  string `json:"host"`
	Token string `json:"token"`
}

func New() *PoolIn {
	return PIN
}

func Init(Puid, Token string) *PoolIn {
	PIN = &PoolIn{
		Puid:  Puid,
		HOST:  ApiUrl,
		Token: Token,
	}
	return PIN
}
