package skrill

// StatusResponseStatus represents Status of StatusResponse
type StatusResponseStatus int

// StatusResponseStatus list
const (
	StatusResponseStatusProcessed  StatusResponseStatus = 2
	StatusResponseStatusPending    StatusResponseStatus = 0
	StatusResponseStatusCancelled  StatusResponseStatus = -1
	StatusResponseStatusFailed     StatusResponseStatus = -2
	StatusResponseStatusChargeback StatusResponseStatus = -3
)

func (status StatusResponseStatus) String() string {
	switch status {
	case StatusResponseStatusProcessed:
		return "processed"
	case StatusResponseStatusPending:
		return "pending"
	case StatusResponseStatusCancelled:
		return "cancelled"
	case StatusResponseStatusFailed:
		return "failed"
	case StatusResponseStatusChargeback:
		return "chargeback"
	default:
		return "unkown"
	}
}

// StatusResponse describes a request body from Skrill API when changing status of payment
type StatusResponse struct {
	TransactionID   string               `json:"transaction_id"`
	MbAmount        float64              `json:"mb_amount,string"`
	Amount          float64              `json:"amount,string"`
	MbTransactionID string               `json:"mb_transaction_id"`
	MbCurrency      string               `json:"mb_currency"`
	PayFromEmail    string               `json:"pay_from_email"`
	Md5Sig          string               `json:"md5sig"`
	PayToEmail      string               `json:"pay_to_email"`
	Currency        string               `json:"currency"`
	MerchantID      string               `json:"merchant_id"`
	Status          StatusResponseStatus `json:"status,string"`
}
