package statusresponse

import (
	"github.com/motomux/skrill-go/currency"
	code "github.com/motomux/skrill-go/status_response/failed_reason_code"
	"github.com/motomux/skrill-go/status_response/status"
)

// StatusResponse describes a request body from Skrill API when changing status of payment
type StatusResponse struct {
	PayToEmail       string            `json:"pay_to_email"`
	PayFromEmail     string            `json:"pay_from_email"`
	MerchantID       string            `json:"merchant_id"`
	CustomerID       string            `json:"customer_id,omitempty"`
	TransactionID    string            `json:"transaction_id,omitempty"`
	MbTransactionID  string            `json:"mb_transaction_id"`
	MbAmount         float64           `json:"mb_amount,string"`
	MbCurrency       currency.Currency `json:"mb_currency"`
	Status           status.Status     `json:"status,string"`
	FailedReasonCode code.Code         `json:"failed_reason_code,string,omitempty"`
	Md5Sig           string            `json:"md5sig"`
	Sha2Sig          string            `json:"sha2sig"`
	Amount           float64           `json:"amount,string"`
	Currency         currency.Currency `json:"currency"`
	NetellerID       string            `json:"neteller_id,omitempty"`
	PaymentType      string            `json:"payment_type,omitempty"`
}
