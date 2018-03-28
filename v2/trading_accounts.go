package quoinex

import (
	"context"
	"fmt"
	"github.com/jjjjpppp/quoinex-go-client/v2/models"
	"net/url"
	"strconv"
	"strings"
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

func (c *Client) GetATradingAccount(ctx context.Context, tradingAccountID int) (*models.TradingAccount, error) {
	spath := fmt.Sprintf("/trading_accounts/%d", tradingAccountID)

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

	var tradingAccount *models.TradingAccount
	if err := decodeBody(res, &tradingAccount); err != nil {
		return nil, err
	}

	return tradingAccount, nil
}

func (c *Client) UpdateLeverageLevel(ctx context.Context, tradeAccountID, leverageLevel int) (*models.TradingAccount, error) {
	spath := fmt.Sprintf("/trading_accounts/%d", tradeAccountID)
	values := url.Values{}
	values.Add("leverage_level", strconv.Itoa(leverageLevel))
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

	var tradingAccount models.TradingAccount
	if err := decodeBody(res, &tradingAccount); err != nil {
		return nil, err
	}

	return &tradingAccount, nil
}
