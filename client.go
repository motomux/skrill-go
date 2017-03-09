package skrill

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
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

// PaymentSource describes describes the payment source used to make Prepare
type PaymentSource struct {
	PayToEmail  string `json:"pay_to_email"`
	PrepareOnly string `json:"prepare_only"`
	Amount      int    `json:"amount"`
	Currency    string `json:"currency"`
}

// Prepare make a request to prepare payment and returns redirect url
func (c *Client) Prepare(param PaymentSource) (redirectURL string, err error) {
	param.PrepareOnly = "1"
	b := &bytes.Buffer{}
	json.NewEncoder(b).Encode(param)
	res, err := http.Post(c.url, "application/json", b)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return genRedirectURL(c.url, string(bs)), nil
}

func genRedirectURL(url, sessionID string) string {
	return url + "?sid=" + sessionID
}