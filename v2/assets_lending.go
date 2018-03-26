package quoinex

import (
	"context"
	"fmt"
	"github.com/jjjjpppp/quoinex-go-client/v2/models"
	"net/url"
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
