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
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		trades, _ := client.GetTrades(ctx, c.param.fundingCurrency, c.param.status)
		if !cmp.Equal(trades, c.expect.trades) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(trades, c.expect.trades))
		}

	}
}

func TestCloseTrade(t *testing.T) {
	type Param struct {
		tradeID        int
		closedQuantity float64
		jsonResponse   string
	}
	type Expect struct {
		path   string
		method string
		body   string
		trade  *models.Trade
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{tradeID: 57896, closedQuantity: 0.0001, jsonResponse: testutil.GetCloseTradeJsonResponse()},
			expect: Expect{path: "/trades/57896/close", method: "PUT", body: "closed_quantity=0.0001", trade: testutil.GetExpectedCloseTradeModel()},
		},
		// test case 2
	}
	for _, c := range cases {
		// preparing test server
		ts := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.URL.RequestURI() != c.expect.path {
					t.Errorf("worng URL. actual:%+v, expect:%+v", r.URL.RequestURI(), c.expect.path)
				}
				if r.Method != c.expect.method {
					t.Errorf("worng Method. actual:%+v, expect:%+v", r.Method, c.expect.method)
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
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		trade, _ := client.CloseTrade(ctx, c.param.tradeID, c.param.closedQuantity)
		if !cmp.Equal(trade, c.expect.trade) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(trade, c.expect.trade))
		}

	}
}

func TestCloseAllTrade(t *testing.T) {
	type Param struct {
		side         string
		jsonResponse string
	}
	type Expect struct {
		path   string
		method string
		body   string
		trades []*models.Trade
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{side: "short", jsonResponse: testutil.GetCloseAllTradeJsonResponse()},
			expect: Expect{path: "/trades/close_all", method: "PUT", body: "side=short", trades: testutil.GetExpectedCloseAllTradeModel()},
		},
		// test case 2
	}
	for _, c := range cases {
		// preparing test server
		ts := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.URL.RequestURI() != c.expect.path {
					t.Errorf("worng URL. actual:%+v, expect:%+v", r.URL.RequestURI(), c.expect.path)
				}
				if r.Method != c.expect.method {
					t.Errorf("worng Method. actual:%+v, expect:%+v", r.Method, c.expect.method)
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
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		trades, _ := client.CloseAllTrade(ctx, c.param.side)
		if !cmp.Equal(trades, c.expect.trades) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(trades, c.expect.trades))
		}

	}
}

func TestUpdateTrade(t *testing.T) {
	type Param struct {
		tradeID      int
		stop_loss    int
		take_profit  int
		jsonResponse string
	}
	type Expect struct {
		path   string
		method string
		body   string
		trade  *models.Trade
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{tradeID: 57897, stop_loss: 300, take_profit: 600, jsonResponse: testutil.GetUpdateTradeJsonResponse()},
			expect: Expect{path: "/trades/57897", method: "PUT", body: "stop_loss=300&take_profit=600", trade: testutil.GetExpectedUpdateTradeModel()},
		},
		// test case 2
	}
	for _, c := range cases {
		// preparing test server
		ts := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.URL.RequestURI() != c.expect.path {
					t.Errorf("worng URL. actual:%+v, expect:%+v", r.URL.RequestURI(), c.expect.path)
				}
				if r.Method != c.expect.method {
					t.Errorf("worng Method. actual:%+v, expect:%+v", r.Method, c.expect.method)
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
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		trade, _ := client.UpdateTrade(ctx, c.param.tradeID, c.param.stop_loss, c.param.take_profit)
		if !cmp.Equal(trade, c.expect.trade) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(trade, c.expect.trade))
		}

	}
}

func TestGetTradesLoans(t *testing.T) {
	type Param struct {
		tradeID      int
		jsonResponse string
	}
	type Expect struct {
		path   string
		method string
		loans  []*models.Loan
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{tradeID: 103520, jsonResponse: testutil.GetTradesLoansJsonResponse()},
			expect: Expect{path: "/trades/103520/loans", method: "GET", loans: testutil.GetExpectedTradesLoansModel()},
		},
		// test case 2
	}
	for _, c := range cases {
		// preparing test server
		ts := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.URL.RequestURI() != c.expect.path {
					t.Errorf("worng URL. actual:%+v, expect:%+v", r.URL.RequestURI(), c.expect.path)
				}
				if r.Method != c.expect.method {
					t.Errorf("worng Method. actual:%+v, expect:%+v", r.Method, c.expect.method)
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
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		loans, _ := client.GetTradesLoans(ctx, c.param.tradeID)
		if !cmp.Equal(loans, c.expect.loans) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(loans, c.expect.loans))
		}

	}
}
