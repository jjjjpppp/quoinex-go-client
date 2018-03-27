package quoinex

import (
	"context"
	"fmt"
	"github.com/jjjjpppp/quoinex-go-client/v2/models"
	"net/url"
	"strconv"
	"strings"
)

func (c *Client) CreateALoanBid(ctx context.Context, quantity, currency, rate string) (*models.LoanBid, error) {
	spath := fmt.Sprintf("/loan_bids/")
	values := url.Values{}
	values.Add("quantity", quantity)
	values.Add("currency", currency)
	values.Add("rate", rate)
	req, err := c.newRequest(ctx, "POST", spath, strings.NewReader(values.Encode()), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("faild to get data. status: %s", res.Status)
	}

	var loanBid models.LoanBid
	if err := decodeBody(res, &loanBid); err != nil {
		return nil, err
	}

	return &loanBid, nil
}

func (c *Client) GetLoanBids(ctx context.Context, currency string) (*models.LoanBids, error) {
	spath := fmt.Sprintf("/loan_bids")
	queryParam := &map[string]string{
		"currency": currency}
	req, err := c.newRequest(ctx, "GET", spath, nil, queryParam)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("faild to get data. status: %s", res.Status)
	}

	var loanBids models.LoanBids
	if err := decodeBody(res, &loanBids); err != nil {
		return nil, err
	}

	return &loanBids, nil
}

func (c *Client) CloseLoanBid(ctx context.Context, loanBidID int) (*models.LoanBid, error) {
	spath := fmt.Sprintf("/loan_bids/%d/close", loanBidID)
	req, err := c.newRequest(ctx, "PUT", spath, nil, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("faild to get data. status: %s", res.Status)
	}

	var loanBid models.LoanBid
	if err := decodeBody(res, &loanBid); err != nil {
		return nil, err
	}

	return &loanBid, nil
}

func (c *Client) GetLoans(ctx context.Context, currency string) (*models.Loans, error) {
	spath := fmt.Sprintf("/loans")
	queryParam := &map[string]string{
		"currency": currency}
	req, err := c.newRequest(ctx, "GET", spath, nil, queryParam)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("faild to get data. status: %s", res.Status)
	}

	var loans models.Loans
	if err := decodeBody(res, &loans); err != nil {
		return nil, err
	}

	return &loans, nil
}

func (c *Client) UpdateALoan(ctx context.Context, loanID int, fundReloaned bool) (*models.Loan, error) {
	spath := fmt.Sprintf("/loans/%d", loanID)
	values := url.Values{}
	values.Add("fund_reloaned", strconv.FormatBool(fundReloaned))
	req, err := c.newRequest(ctx, "PUT", spath, strings.NewReader(values.Encode()), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("faild to get data. status: %s", res.Status)
	}

	var loan models.Loan
	if err := decodeBody(res, &loan); err != nil {
		return nil, err
	}

	return &loan, nil
}
