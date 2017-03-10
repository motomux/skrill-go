[![Build Status](https://travis-ci.org/motomux/skrill-go.svg?branch=master)](https://travis-ci.org/motomux/skrill-go)

# skrill-go
Skrill client for Golang

## Example

### Prepare

```
client := skrill.New()
redirectURL, err := client.Prepare(skrill.PrepareParam{
	PayToEmail:    "YOUR_MERCHANT_EMAIL",
	Amount:        10.00,
	Currency:      currency.USD,
	TransactionID: "UNIQUE_ID",
	StatusURL:     "WEBHOOK_URL",
})
```
