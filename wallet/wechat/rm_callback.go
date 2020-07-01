package wechat

type RmCallback struct {
	OrderNo   string  `json:"order_no"`
	Status    string  `json:"status"`
	PayAmount float64 `json:"pay_amount"`
	PayTime   string  `json:"pay_time"`
	PayType   int64   `json:"pay_type"`
}
