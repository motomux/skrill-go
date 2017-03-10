package skrill

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/motomux/skrill-go/currency"
)

const defaultURL = "https://pay.skrill.com/"

// Client is a Skrill client
type Client struct {
	url string
}

// Config is configration to initiate Skrill client
type Config struct {
	URL string
}

// New initiates Skrill client
func New(configs ...Config) *Client {
	client := &Client{
		url: defaultURL,
	}

	if len(configs) != 0 {
		config := configs[0]
		client.url = config.URL
	}
	return client
}

// PrepareParam describes describes the payment source used to make Prepare
type PrepareParam struct {
	PayToEmail    string            `json:"pay_to_email"`
	PrepareOnly   string            `json:"prepare_only"`
	Amount        int               `json:"amount"`
	Currency      currency.Currency `json:"currency"`
	TransactionID string            `json:"transaction_id,omitempty"`
	ReturnURL     string            `json:"return_url,omitempty"`
	StatusURL     string            `json:"status_url,omitempty"`
}

// Prepare make a request to prepare payment and returns redirect url
func (c *Client) Prepare(param PrepareParam) (redirectURL string, err error) {
	param.PrepareOnly = "1"
	b := &bytes.Buffer{}
	json.NewEncoder(b).Encode(param)
	res, err := http.Post(c.url, "application/json", b)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var err ErrSkrill
		json.NewDecoder(res.Body).Decode(&err)
		return "", err
	}

	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return genRedirectURL(c.url, string(bs)), nil
}

func genRedirectURL(url, sessionID string) string {
	return url + "?sid=" + sessionID
}
