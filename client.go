package skrill

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/motomux/skrill-go/currency"
	"github.com/motomux/skrill-go/language"
	"github.com/motomux/skrill-go/target"
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
	// Merchant details
	PayToEmail           string            `json:"pay_to_email"`
	RecipientDescription string            `json:"recipient_description,omitempty"`
	TransactionID        string            `json:"transaction_id,omitempty"`
	ReturnURL            string            `json:"return_url,omitempty"`
	ReturnURLText        string            `json:"return_url_text,omitempty"`
	ReturnURLTarget      target.Target     `json:"return_url_target,omitempty"`
	CancelURL            string            `json:"cancel_url,omitempty"`
	CancelURLTarget      string            `json:"cancel_url_target,omitempty"`
	StatusURL            string            `json:"status_url,omitempty"`
	StatusURL2           string            `json:"status_url2,omitempty"`
	Language             language.Language `json:"language,omitempty"`
	LogoURL              string            `json:"logo_url,omitempty"`
	PrepareOnly          string            `json:"prepare_only"`
	SID                  string            `json:"sid,omitempty"`
	RID                  string            `json:"rid,omitempty"`
	ExtRefID             string            `json:"ext_ref_id,omitempty"`
	DynamicDescriptor    string            `json:"dynamic_descriptor,omitempty"`
	PayFromEmail         string            `json:"pay_from_email,omitempty"`
	FirstName            string            `json:"firstname,omitempty"`
	LastName             string            `json:"lastname,omitempty"`
	DateOfBirth          string            `json:"date_of_birth,omitempty"`
	Address              string            `json:"address,omitempty"`
	Address2             string            `json:"address2,omitempty"`
	PhoneNumber          string            `json:"phone_number,omitempty"`
	PostalCode           string            `json:"postal_code,omitempty"`
	City                 string            `json:"city,omitempty"`
	Country              string            `json:"country,omitempty"`
	NetellerAccount      string            `json:"neteller_account,omitempty"`
	NetellerSecureID     string            `json:"neteller_secure_id,omitempty"`
	// Payment details
	Amount             float64           `json:"amount"`
	Currency           currency.Currency `json:"currency"`
	Amount2Description string            `json:"amount2_description,omitempty"`
	Amount2            float64           `json:"amount2,omitempty"`
	Amount3Description string            `json:"amount3_description,omitempty"`
	Amount3            float64           `json:"amount3,omitempty"`
	Amount4Description string            `json:"amount4_description,omitempty"`
	Amount4            float64           `json:"amount4,omitempty"`
	Detail1Description string            `json:"detail1_description,omitempty"`
	Detail1Text        string            `json:"detail1_text,omitempty"`
	Detail2Description string            `json:"detail2_description,omitempty"`
	Detail2Text        string            `json:"detail2_text,omitempty"`
	Detail3Description string            `json:"detail3_description,omitempty"`
	Detail3Text        string            `json:"detail3_text,omitempty"`
	Detail4Description string            `json:"detail4_description,omitempty"`
	Detail4Text        string            `json:"detail4_text,omitempty"`
	Detail5Description string            `json:"detail5_description,omitempty"`
	Detail5Text        string            `json:"detail5_text,omitempty"`
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
		if e := json.NewDecoder(res.Body).Decode(&err); e != nil {
			return "", e
		}
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
