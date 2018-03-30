package quoinex

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/jjjjpppp/quoinex-go-client/v2/models"
	"github.com/jjjjpppp/quoinex-go-client/v2/testutil"
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
		body     string
		accounts []*models.Account
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{jsonResponse: testutil.GetFiatAccountsJsonResponse()},
			expect: Expect{path: "/fiat_accounts", method: "GET", body: "", accounts: testutil.GetExpectedFiatAccountsModel()},
		},
		// test case 2
	}
	for _, c := range cases {
		// preparing test server
		ts := testutil.GenerateTestServer(t, c.expect.path, c.expect.method, c.expect.body, c.param.jsonResponse)
		defer ts.Close()

		client, _ := NewClient("apiTokenID", "secret", nil)
		client.testServer = ts
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
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
		ts := testutil.GenerateTestServer(t, c.expect.path, c.expect.method, c.expect.body, c.param.jsonResponse)
		defer ts.Close()

		client, _ := NewClient("apiTokenID", "secret", nil)
		client.testServer = ts
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
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
		body     string
		accounts []*models.CryptoAccount
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{jsonResponse: testutil.GetCryptoAccountsJsonResponse()},
			expect: Expect{path: "/crypto_accounts", method: "GET", body: "", accounts: testutil.GetExpectedCryptoAccountsModel()},
		},
		// test case 2
	}
	for _, c := range cases {
		// preparing test server
		ts := testutil.GenerateTestServer(t, c.expect.path, c.expect.method, c.expect.body, c.param.jsonResponse)
		defer ts.Close()

		client, _ := NewClient("apiTokenID", "secret", nil)
		client.testServer = ts
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
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
		body            string
		accountBalances []*models.AccountBalance
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{jsonResponse: testutil.GetAllAccountBalancesJsonResponse()},
			expect: Expect{path: "/accounts/balance", method: "GET", body: "", accountBalances: testutil.GetExpectedAllAccountBalancesModel()},
		},
		// test case 2
	}
	for _, c := range cases {
		// preparing test server
		ts := testutil.GenerateTestServer(t, c.expect.path, c.expect.method, c.expect.body, c.param.jsonResponse)
		defer ts.Close()

		client, _ := NewClient("apiTokenID", "secret", nil)
		client.testServer = ts
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		accountBalances, _ := client.GetAllAccountBalances(ctx)
		if !cmp.Equal(accountBalances, c.expect.accountBalances) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(accountBalances, c.expect.accountBalances))
		}

	}
}
