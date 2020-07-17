package quoinex

import (
	"context"
	"fmt"
	"github.com/jjjjpppp/quoinex-go-client/v2/models"
	"strings"
)

func (c *Client) GetFiatAccounts(ctx context.Context) ([]*models.Account, error) {
	spath := "/fiat_accounts"
	res, err := c.sendRequest(ctx, "GET", spath, nil, nil)
	if err != nil {
		return nil, err
	}

	var accounts []*models.Account
	if err := decodeBody(res, &accounts); err != nil {
		return nil, err
	}

	return accounts, nil
}

func (c *Client) CreateAFiatAccount(ctx context.Context, currency string) (*models.Account, error) {
	spath := "/fiat_accounts"
	body := fmt.Sprintf("{\"currency\":\"%s\"}", currency)
	res, err := c.sendRequest(ctx, "POST", spath, strings.NewReader(body), nil)
	if err != nil {
		return nil, err
	}

	var account models.Account
	if err := decodeBody(res, &account); err != nil {
		return nil, err
	}

	return &account, nil
}

func (c *Client) GetCryptoAccounts(ctx context.Context) ([]*models.CryptoAccount, error) {
	spath := "/crypto_accounts"
	res, err := c.sendRequest(ctx, "GET", spath, nil, nil)
	if err != nil {
		return nil, err
	}

	var accounts []*models.CryptoAccount
	if err := decodeBody(res, &accounts); err != nil {
		return nil, err
	}

	return accounts, nil
}

func (c *Client) GetAllAccountBalances(ctx context.Context) ([]*models.AccountBalance, error) {
	spath := "/accounts/balance"
	res, err := c.sendRequest(ctx, "GET", spath, nil, nil)
	if err != nil {
		return nil, err
	}

	var accountBalances []*models.AccountBalance
	if err := decodeBody(res, &accountBalances); err != nil {
		return nil, err
	}

	return accountBalances, nil
}
