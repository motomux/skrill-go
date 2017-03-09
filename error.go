package skrill

import "fmt"

// ErrSkrill describes error response from Skrill API
type ErrSkrill struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e ErrSkrill) Error() string {
	return fmt.Sprintf("Skrill returns error response. code=%q message=%q", e.Code, e.Message)
}
