package quoinex

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/jjjjpppp/quoinex-go-client/v2/models"
	"github.com/jjjjpppp/quoinex-go-client/v2/testutil"
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
		ts := testutil.GenerateTestServer(t, c.expect.path, c.expect.method, c.expect.body, c.param.jsonResponse)
		defer ts.Close()

		client, _ := NewClient("apiTokenID", "secret", nil)
		client.testServer = ts
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
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
		body     string
		loanBids *models.LoanBids
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{currency: "USD", jsonResponse: testutil.GetLoanBidsJsonResponse()},
			expect: Expect{path: "/loan_bids?currency=USD", method: "GET", body: "", loanBids: testutil.GetExpectedLoanBidsModel()},
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
		body    string
		loanBid *models.LoanBid
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{loanBidID: 3580, jsonResponse: testutil.GetCloseLoanBidJsonResponse()},
			expect: Expect{path: "/loan_bids/3580/close", method: "PUT", body: "", loanBid: testutil.GetExpectedCloseLoanBidModel()},
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
		body   string
		loans  *models.Loans
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{currency: "JPY", jsonResponse: testutil.GetLoansJsonResponse()},
			expect: Expect{path: "/loans?currency=JPY", method: "GET", body: "", loans: testutil.GetExpectedLoansModel()},
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
		loans, _ := client.GetLoans(ctx, c.param.currency)
		if !cmp.Equal(loans, c.expect.loans) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(loans, c.expect.loans))
		}

	}
}

func TestUpdateALoan(t *testing.T) {
	type Param struct {
		loanID       int
		fundReloaned bool
		jsonResponse string
	}
	type Expect struct {
		path   string
		method string
		body   string
		loan   *models.Loan
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{loanID: 144825, fundReloaned: false, jsonResponse: testutil.GetUpdateALoanJsonResponse()},
			expect: Expect{path: "/loans/144825", method: "PUT", body: "fund_reloaned=false", loan: testutil.GetExpectedUpdateALoanModel()},
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
		loan, _ := client.UpdateALoan(ctx, c.param.loanID, c.param.fundReloaned)
		if !cmp.Equal(loan, c.expect.loan) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(loan, c.expect.loan))
		}

	}
}
