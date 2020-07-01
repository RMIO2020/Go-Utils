package wechat

type RmCallback struct {
	OrderNo   string `json:"order_no"`
	Status    string `json:"status"`
	PayAmount string `json:"pay_amount"`
	PayTime   string `json:"pay_time"`
	PayType   string `json:"pay_type"`
}
