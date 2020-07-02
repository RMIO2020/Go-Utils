package wechat

type RmCallback struct {
	OrderNo     string `json:"order_no"`
	PayAmount   string `json:"pay_amount"`
	ProcessType string `json:"process_type"`
	PayType     int64  `json:"payment_type"`
}
