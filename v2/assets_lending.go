package quoinex

import (
	"context"
	"fmt"
	"github.com/jjjjpppp/quoinex-go-client/v2/models"
	"strconv"
	"strings"
)

func (c *Client) CreateALoanBid(ctx context.Context, quantity, currency, rate string) (*models.LoanBid, error) {
	spath := "/loan_bids"
	bodyTemplate :=
		`{
			"loan_bid": {
				"quantity":"%s",
				"currency":"%s",
				"rate":"%s"
			}
		}`
	body := fmt.Sprintf(bodyTemplate, quantity, currency, rate)
	res, err := c.sendRequest(ctx, "POST", spath, strings.NewReader(body), nil)
	if err != nil {
		return nil, err
	}

	var loanBid models.LoanBid
	if err := decodeBody(res, &loanBid); err != nil {
		return nil, err
	}

	return &loanBid, nil
}

func (c *Client) GetLoanBids(ctx context.Context, currency string) (*models.LoanBids, error) {
	spath := "/loan_bids"
	queryParam := &map[string]string{
		"currency": currency}
	res, err := c.sendRequest(ctx, "GET", spath, nil, queryParam)
	if err != nil {
		return nil, err
	}

	var loanBids models.LoanBids
	if err := decodeBody(res, &loanBids); err != nil {
		return nil, err
	}

	return &loanBids, nil
}

func (c *Client) CloseLoanBid(ctx context.Context, loanBidID int) (*models.LoanBid, error) {
	spath := fmt.Sprintf("/loan_bids/%d/close", loanBidID)
	res, err := c.sendRequest(ctx, "PUT", spath, nil, nil)
	if err != nil {
		return nil, err
	}

	var loanBid models.LoanBid
	if err := decodeBody(res, &loanBid); err != nil {
		return nil, err
	}

	return &loanBid, nil
}

func (c *Client) GetLoans(ctx context.Context, currency string) (*models.Loans, error) {
	spath := "/loans"
	queryParam := &map[string]string{
		"currency": currency}
	res, err := c.sendRequest(ctx, "GET", spath, nil, queryParam)
	if err != nil {
		return nil, err
	}

	var loans models.Loans
	if err := decodeBody(res, &loans); err != nil {
		return nil, err
	}

	return &loans, nil
}

func (c *Client) UpdateALoan(ctx context.Context, loanID int, fundReloaned bool) (*models.Loan, error) {
	spath := fmt.Sprintf("/loans/%d", loanID)
	bodyTemplate :=
		`{
			"loan": {
				"fund_reloaned":%s
			}
		}`
	body := fmt.Sprintf(bodyTemplate, strconv.FormatBool(fundReloaned))
	res, err := c.sendRequest(ctx, "PUT", spath, strings.NewReader(body), nil)
	if err != nil {
		return nil, err
	}

	var loan models.Loan
	if err := decodeBody(res, &loan); err != nil {
		return nil, err
	}

	return &loan, nil
}
