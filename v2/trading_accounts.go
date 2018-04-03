package quoinex

import (
	"context"
	"fmt"
	"github.com/jjjjpppp/quoinex-go-client/v2/models"
	"strings"
)

func (c *Client) GetTradingAccounts(ctx context.Context) ([]*models.TradingAccount, error) {
	spath := fmt.Sprintf("/trading_accounts")
	res, err := c.sendRequest(ctx, "GET", spath, nil, nil)
	if err != nil {
		return nil, err
	}

	var tradingAccounts []*models.TradingAccount
	if err := decodeBody(res, &tradingAccounts); err != nil {
		return nil, err
	}

	return tradingAccounts, nil
}

func (c *Client) GetATradingAccount(ctx context.Context, tradingAccountID int) (*models.TradingAccount, error) {
	spath := fmt.Sprintf("/trading_accounts/%d", tradingAccountID)

	res, err := c.sendRequest(ctx, "GET", spath, nil, nil)
	if err != nil {
		return nil, err
	}

	var tradingAccount *models.TradingAccount
	if err := decodeBody(res, &tradingAccount); err != nil {
		return nil, err
	}

	return tradingAccount, nil
}

func (c *Client) UpdateLeverageLevel(ctx context.Context, tradeAccountID, leverageLevel int) (*models.TradingAccount, error) {
	spath := fmt.Sprintf("/trading_accounts/%d", tradeAccountID)
	bodyTemplate :=
		`{
			"trading_account": {
				"leverage_level":%d
			}
		}`
	body := fmt.Sprintf(bodyTemplate, leverageLevel)
	res, err := c.sendRequest(ctx, "PUT", spath, strings.NewReader(body), nil)
	if err != nil {
		return nil, err
	}

	var tradingAccount models.TradingAccount
	if err := decodeBody(res, &tradingAccount); err != nil {
		return nil, err
	}

	return &tradingAccount, nil
}
