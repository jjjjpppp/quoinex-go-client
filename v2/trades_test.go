package quoinex

import (
	"context"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/jjjjpppp/quoinex-go-client/v2/models"
	"github.com/jjjjpppp/quoinex-go-client/v2/testutil"
	//"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetTrades(t *testing.T) {
	type Param struct {
		fundingCurrency string
		status          string
		jsonResponse    string
	}
	type Expect struct {
		path   string
		method string
		trades *models.Trades
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{fundingCurrency: "USD", status: "open", jsonResponse: testutil.GetTradesJsonResponse()},
			expect: Expect{path: "/trades?funding_currency=USD&status=open", method: "GET", trades: testutil.GetExpectedTradesModel()},
		},
		// test case 2
	}
	for _, c := range cases {
		// preparing test server
		ts := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.URL.RequestURI() != c.expect.path {
					t.Fatalf("worng URL. actual:%+v, expect:%+v", r.URL.RequestURI(), c.expect.path)
				}
				if r.Method != c.expect.method {
					t.Fatalf("worng Method. actual:%+v, expect:%+v", r.Method, c.expect.method)
				}
				// set expected json
				w.Header().Set("content-Type", "text")
				fmt.Fprintf(w, c.param.jsonResponse)
				return
			},
		))
		defer ts.Close()

		client, _ := NewClient("apiTokenID", "secret", nil)
		client.testServer = ts
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		trades, _ := client.GetTrades(ctx, c.param.fundingCurrency, c.param.status)
		if !cmp.Equal(trades, c.expect.trades) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(trades, c.expect.trades))
		}

	}
}
