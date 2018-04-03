package quoinex

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/jjjjpppp/quoinex-go-client/v2/models"
	"github.com/jjjjpppp/quoinex-go-client/v2/testutil"
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
		body            string
		tradingAccounts []*models.TradingAccount
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{jsonResponse: testutil.GetTradingAccountsJsonResponse()},
			expect: Expect{path: "/trading_accounts", method: "GET", body: "", tradingAccounts: testutil.GetExpectedTradingAccounts()},
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
		body           string
		tradingAccount *models.TradingAccount
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{tradingAccountID: 1759, jsonResponse: testutil.GetATradingAccountJsonResponse()},
			expect: Expect{path: "/trading_accounts/1759", method: "GET", body: "", tradingAccount: testutil.GetExpectedATradingAccount()},
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
			expect: Expect{path: "/trading_accounts/1759", body: testutil.GetExpectedUpdateLeverageLevelRequestBody(), method: "PUT", tradingAccount: testutil.GetExpectedUpdateLeverageLevel()},
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
		tradingAccount, _ := client.UpdateLeverageLevel(ctx, c.param.tradingAccountID, c.param.leverageLevel)
		if !cmp.Equal(tradingAccount, c.expect.tradingAccount) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(tradingAccount, c.expect.tradingAccount))
		}

	}
}
