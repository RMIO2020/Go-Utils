package f2

const ApiUrl = "https://api-prod.poolin.com"

var PIN *PoolIn

type PoolIn struct {
	UserName string `json:"puid"`
	HOST     string `json:"host"`
}

func New() *PoolIn {
	return PIN
}

func Init(UserName string) *PoolIn {
	PIN = &PoolIn{
		UserName: UserName,
		HOST:     ApiUrl,
	}
	return PIN
}
