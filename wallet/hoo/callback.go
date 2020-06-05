package hoo

type Transaction struct {
	Sign            string `json:"sign" binding:"required"`
	ChainName       string `json:"chain_name" binding:"required"`
	CoinName        string `json:"coin_name" binding:"required"`
	Alias           string `json:"alias" binding:"required"`
	TradType        string `json:"trad_type" binding:"required"`
	BlockHeight     string `json:"block_height" binding:"required"`
	TransactionId   string `json:"transaction_id" binding:"required"`
	TrxN            string `json:"trx_n" binding:"required"`
	Confirmations   string `json:"confirmations" binding:"required"`
	FromAddress     string `json:"from_address" binding:"required"`
	ToAddress       string `json:"toAddress" binding:"required"`
	Memo            string `json:"memo" binding:"required"`
	Amount          string `json:"amount" binding:"required"`
	Fee             string `json:"fee" binding:"required"`
	ContractAddress string `json:"contractAddress" binding:"required"`
	OuterOrderNo    string `json:"outerOrderNo" binding:"required"`
	ConfirmTime     string `json:"confirmTime" binding:"required"`
	Message         string `json:"message" binding:"required"`
	Currency        string
	Protocol        string
}
