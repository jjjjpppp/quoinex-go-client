package quoinex

import (
	"context"
	"fmt"
	"github.com/jjjjpppp/quoinex-go-client/v2/models"
	//"net/url"
	//"strconv"
	//"strings"
)

func (c *Client) GetTradingAccounts(ctx context.Context) ([]*models.TradingAccount, error) {
	spath := fmt.Sprintf("/trading_accounts")
	req, err := c.newRequest(ctx, "GET", spath, nil, nil)
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

	var tradingAccounts []*models.TradingAccount
	if err := decodeBody(res, &tradingAccounts); err != nil {
		return nil, err
	}

	return tradingAccounts, nil
}
