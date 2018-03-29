package quoinex

import (
	"context"
	"fmt"
	"github.com/jjjjpppp/quoinex-go-client/v2/models"
	//"net/url"
	//"strconv"
	//"strings"
)

func (c *Client) GetTrades(ctx context.Context, fundingCurrency, status string) (*models.Trades, error) {
	spath := fmt.Sprintf("/trades")
	queryParam := &map[string]string{
		"funding_currency": fundingCurrency,
		"status":           status}
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

	var trades models.Trades
	if err := decodeBody(res, &trades); err != nil {
		return nil, err
	}

	return &trades, nil
}
