package quoinex

import (
	"context"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/jjjjpppp/quoinex-go-client/v2/models"
	"github.com/jjjjpppp/quoinex-go-client/v2/testutil"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetTradingAccounts(t *testing.T) {
	type Param struct {
		jsonResponse string
	}
	type Expect struct {
		path            string
		method          string
		tradingAccounts []*models.TradingAccount
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{jsonResponse: testutil.GetTradingAccountsJsonResponse()},
			expect: Expect{path: "/trading_accounts", method: "GET", tradingAccounts: testutil.GetExpectedTradingAccounts()},
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
		tradingAccounts, _ := client.GetTradingAccounts(ctx)
		if !cmp.Equal(tradingAccounts, c.expect.tradingAccounts) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(tradingAccounts, c.expect.tradingAccounts))
		}

	}
}

func TestGetATradingAccount(t *testing.T) {
	type Param struct {
		tradingAccountID int
		jsonResponse     string
	}
	type Expect struct {
		path           string
		method         string
		tradingAccount *models.TradingAccount
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{tradingAccountID: 1759, jsonResponse: testutil.GetATradingAccountJsonResponse()},
			expect: Expect{path: "/trading_accounts/1759", method: "GET", tradingAccount: testutil.GetExpectedATradingAccount()},
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
		tradingAccount, _ := client.GetATradingAccount(ctx, c.param.tradingAccountID)
		if !cmp.Equal(tradingAccount, c.expect.tradingAccount) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(tradingAccount, c.expect.tradingAccount))
		}

	}
}

func TestUpdateLeverageLevel(t *testing.T) {
	type Param struct {
		tradingAccountID int
		leverageLevel    int
		jsonResponse     string
	}
	type Expect struct {
		path           string
		method         string
		body           string
		tradingAccount *models.TradingAccount
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{tradingAccountID: 1759, leverageLevel: 25, jsonResponse: testutil.GetUpdateLeverageLevelJsonResponse()},
			expect: Expect{path: "/trading_accounts/1759", body: "leverage_level=25", method: "PUT", tradingAccount: testutil.GetExpectedUpdateLeverageLevel()},
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
				// Read body
				b, err := ioutil.ReadAll(r.Body)
				s := string(b)
				defer r.Body.Close()
				if err != nil {
					t.Errorf("Worng body. err:%+v", err)
				}
				if s != c.expect.body {
					t.Errorf("Worng body. actual: %+v, expect:%+v", s, c.expect.body)
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
		tradingAccount, _ := client.UpdateLeverageLevel(ctx, c.param.tradingAccountID, c.param.leverageLevel)
		if !cmp.Equal(tradingAccount, c.expect.tradingAccount) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(tradingAccount, c.expect.tradingAccount))
		}

	}
}
