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

func TestGetFiatAccounts(t *testing.T) {
	type Param struct {
		jsonResponse string
	}
	type Expect struct {
		path     string
		method   string
		accounts []*models.Account
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{jsonResponse: testutil.GetFiatAccountsJsonResponse()},
			expect: Expect{path: "/fiat_accounts", method: "GET", accounts: testutil.GetExpectedFiatAccountsModel()},
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
		accounts, _ := client.GetFiatAccounts(ctx)
		if !cmp.Equal(accounts, c.expect.accounts) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(accounts, c.expect.accounts))
		}

	}
}

func TestCreateAFiatAccount(t *testing.T) {
	type Param struct {
		currency     string
		jsonResponse string
	}
	type Expect struct {
		path    string
		method  string
		body    string
		account *models.Account
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{currency: "USD", jsonResponse: testutil.GetCreateFiatAccountJsonResponse()},
			expect: Expect{path: "/fiat_accounts", method: "POST", body: "currency=USD", account: testutil.GetExpectedCreateFiatAccountModel()},
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
		account, _ := client.CreateAFiatAccount(ctx, c.param.currency)
		if !cmp.Equal(account, c.expect.account) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(account, c.expect.account))
		}

	}
}

func TestGetCryptoAccounts(t *testing.T) {
	type Param struct {
		jsonResponse string
	}
	type Expect struct {
		path     string
		method   string
		accounts []*models.CryptoAccount
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{jsonResponse: testutil.GetCryptoAccountsJsonResponse()},
			expect: Expect{path: "/crypto_accounts", method: "GET", accounts: testutil.GetExpectedCryptoAccountsModel()},
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
		accounts, _ := client.GetCryptoAccounts(ctx)
		if !cmp.Equal(accounts, c.expect.accounts) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(accounts, c.expect.accounts))
		}

	}
}

func TestGetAllAccountBalances(t *testing.T) {
	type Param struct {
		jsonResponse string
	}
	type Expect struct {
		path            string
		method          string
		accountBalances []*models.AccountBalance
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{jsonResponse: testutil.GetAllAccountBalancesJsonResponse()},
			expect: Expect{path: "/accounts/balance", method: "GET", accountBalances: testutil.GetExpectedAllAccountBalancesModel()},
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
		accountBalances, _ := client.GetAllAccountBalances(ctx)
		if !cmp.Equal(accountBalances, c.expect.accountBalances) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(accountBalances, c.expect.accountBalances))
		}

	}
}
