package skrill

// StatusResponse describes a request body from Skrill API when changing status of payment
type StatusResponse struct {
	TransactionID   string  `json:"transaction_id"`
	MbAmount        float64 `json:"mb_amount,string"`
	Amount          float64 `json:"amount,string"`
	MbTransactionID string  `json:"mb_transaction_id"`
	MbCurrency      string  `json:"mb_currency"`
	PayFromEmail    string  `json:"pay_from_email"`
	Md5Sig          string  `json:"md5sig"`
	PayToEmail      string  `json:"pay_to_email"`
	Currency        string  `json:"currency"`
	MerchantID      string  `json:"merchant_id"`
	Status          int     `json:"status,string"`
}
