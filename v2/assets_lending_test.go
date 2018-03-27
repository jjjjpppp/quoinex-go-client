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

func TestCreateALoanBid(t *testing.T) {
	type Param struct {
		quantity     string
		currency     string
		rate         string
		jsonResponse string
	}
	type Expect struct {
		path    string
		method  string
		body    string
		loanBid *models.LoanBid
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{currency: "USD", quantity: "50", rate: "0.0002", jsonResponse: testutil.GetCreateLoanBidJsonResponse()},
			expect: Expect{path: "/loan_bids", method: "POST", body: "currency=USD&quantity=50&rate=0.0002", loanBid: testutil.GetExpectedCreateLoanBidModel()},
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
		loanBid, _ := client.CreateALoanBid(ctx, c.param.quantity, c.param.currency, c.param.rate)
		if !cmp.Equal(loanBid, c.expect.loanBid) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(loanBid, c.expect.loanBid))
		}

	}
}

func TestGetLoanBids(t *testing.T) {
	type Param struct {
		currency     string
		jsonResponse string
	}
	type Expect struct {
		path     string
		method   string
		loanBids *models.LoanBids
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{currency: "USD", jsonResponse: testutil.GetLoanBidsJsonResponse()},
			expect: Expect{path: "/loan_bids?currency=USD", method: "GET", loanBids: testutil.GetExpectedLoanBidsModel()},
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
		loanBids, _ := client.GetLoanBids(ctx, c.param.currency)
		if !cmp.Equal(loanBids, c.expect.loanBids) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(loanBids, c.expect.loanBids))
		}

	}
}

func TestCloseLoanBid(t *testing.T) {
	type Param struct {
		loanBidID    int
		jsonResponse string
	}
	type Expect struct {
		path    string
		method  string
		loanBid *models.LoanBid
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{loanBidID: 3580, jsonResponse: testutil.GetCloseLoanBidJsonResponse()},
			expect: Expect{path: "/loan_bids/3580/close", method: "PUT", loanBid: testutil.GetExpectedCloseLoanBidModel()},
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
		loanBid, _ := client.CloseLoanBid(ctx, c.param.loanBidID)
		if !cmp.Equal(loanBid, c.expect.loanBid) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(loanBid, c.expect.loanBid))
		}

	}
}

func TestGetLoans(t *testing.T) {
	type Param struct {
		currency     string
		jsonResponse string
	}
	type Expect struct {
		path   string
		method string
		loans  *models.Loans
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{currency: "JPY", jsonResponse: testutil.GetLoansJsonResponse()},
			expect: Expect{path: "/loans?currency=JPY", method: "GET", loans: testutil.GetExpectedLoansModel()},
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
		loans, _ := client.GetLoans(ctx, c.param.currency)
		if !cmp.Equal(loans, c.expect.loans) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(loans, c.expect.loans))
		}

	}
}
