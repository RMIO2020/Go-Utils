package f2

const ApiUrl = "http://api.f2pool.com"

var PIN *F2

type F2 struct {
	HOST string `json:"host"`
}

func New() *F2 {
	return PIN
}

func Init() *F2 {
	PIN = &F2{
		HOST: ApiUrl,
	}
	return PIN
}
