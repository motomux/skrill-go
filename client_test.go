package skrill

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := map[string]struct {
		configs []Config
		out     *Client
	}{

		"should return client with default url when configs is empty": {
			configs: nil,
			out: &Client{
				url: defaultURL,
			},
		},

		"should return client with given url when configs is not empty": {
			configs: []Config{Config{URL: "http://test.com"}},
			out: &Client{
				url: "http://test.com",
			},
		},

		"should return client with url defined on the first config when configs more than one": {
			configs: []Config{
				Config{URL: "http://test.com"},
				Config{URL: "http://test2.com"},
			},
			out: &Client{
				url: "http://test.com",
			},
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			out := New(test.configs...)
			if !reflect.DeepEqual(out, test.out) {
				t.Errorf("actual=%#v expected=%#v", out, test.out)
			}
		})
	}
}

func TestPrepare(t *testing.T) {
	type (
		in struct {
			param PaymentSource
		}
		out struct {
			redirectURL string
			err         error
		}

		ts struct {
			method, path string
			reqBody      PaymentSource
			resBody      string
		}
	)

	tests := map[string]struct {
		in
		out
		ts
	}{

		"should make a request to skrill server": {
			in{
				param: PaymentSource{
					PayToEmail: "test@test.com",
					Amount:     1,
					Currency:   "USD",
				},
			},
			out{
				redirectURL: "?sid=SESSION_ID",
				err:         nil,
			},
			ts{
				path:   "/",
				method: "POST",
				reqBody: PaymentSource{
					PayToEmail:  "test@test.com",
					PrepareOnly: "1",
					Amount:      1,
					Currency:    "USD",
				},
				resBody: "SESSION_ID",
			},
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				defer r.Body.Close()

				if r.URL.Path != test.ts.path {
					t.Errorf("TestServer path doens't match: actual=%s expected=%s", r.URL.Path, test.ts.path)
				}

				if r.Method != test.ts.method {
					t.Errorf("TestServer method doens't match: actual=%s expected=%s", r.Method, test.ts.method)
				}

				var reqBody PaymentSource
				json.NewDecoder(r.Body).Decode(&reqBody)
				if !reflect.DeepEqual(reqBody, test.ts.reqBody) {
					t.Errorf("TestServer request body doesn't match: actual=%#v expected=%#v", reqBody, test.ts.reqBody)
				}

				w.Write([]byte(test.ts.resBody))
			}))
			defer testServer.Close()

			client := New(Config{URL: testServer.URL})

			redirectURL, err := client.Prepare(test.in.param)
			if err != test.out.err {
				t.Errorf("Output err doens't match: actual=%v expected=%v", err, test.out, err)
			}
			if redirectURL != testServer.URL+test.out.redirectURL {
				t.Errorf("Output redirectURL doesn't match: actual=%v expected=%v", redirectURL, testServer.URL+test.out.redirectURL)
			}
		})
	}
}
