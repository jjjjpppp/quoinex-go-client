package quoinex

import (
	"context"
	"fmt"
	"github.com/jjjjpppp/quoinex-go-client/v2/models"
	"strings"
)

func (c *Client) GetFiatAccounts(ctx context.Context) ([]*models.Account, error) {
	spath := fmt.Sprintf("/fiat_accounts")
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

	var accounts []*models.Account
	if err := decodeBody(res, &accounts); err != nil {
		return nil, err
	}

	return accounts, nil
}

func (c *Client) CreateAFiatAccount(ctx context.Context, currency string) (*models.Account, error) {
	spath := fmt.Sprintf("/fiat_accounts")
	body := fmt.Sprintf("{\"currency\":\"%s\"}", currency)
	req, err := c.newRequest(ctx, "POST", spath, strings.NewReader(body), nil)
	if err != nil {
		c.Logger.Printf("req: %#v \nerr: %#v \n", req, err)
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		c.Logger.Printf("res: %#v \nerr: %#v \n", res, err)
		return nil, err
	}

	if res.StatusCode != 200 {
		c.Logger.Printf("res: %#v \nbody: %s \n", res, responseBodyToString(res))
		return nil, fmt.Errorf("faild to get data. status: %s", res.Status)
	}

	var account models.Account
	if err := decodeBody(res, &account); err != nil {
		return nil, err
	}

	return &account, nil
}

func (c *Client) GetCryptoAccounts(ctx context.Context) ([]*models.CryptoAccount, error) {
	spath := fmt.Sprintf("/crypto_accounts")
	req, err := c.newRequest(ctx, "GET", spath, nil, nil)
	if err != nil {
		c.Logger.Printf("req: %#v \nerr: %#v \n", req, err)
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		c.Logger.Printf("res: %#v \nerr: %#v \n", res, err)
		return nil, err
	}

	if res.StatusCode != 200 {
		c.Logger.Printf("res: %#v \nerr: %#v \n", res, err)
		return nil, fmt.Errorf("faild to get data. status: %s", res.Status)
	}

	var accounts []*models.CryptoAccount
	if err := decodeBody(res, &accounts); err != nil {
		return nil, err
	}

	return accounts, nil
}

func (c *Client) GetAllAccountBalances(ctx context.Context) ([]*models.AccountBalance, error) {
	spath := fmt.Sprintf("/accounts/balance")
	req, err := c.newRequest(ctx, "GET", spath, nil, nil)
	if err != nil {
		c.Logger.Printf("req: %#v \nerr: %#v \n", req, err)
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		c.Logger.Printf("res: %#v \nerr: %#v \n", res, err)
		return nil, err
	}

	if res.StatusCode != 200 {
		c.Logger.Printf("res: %#v \nerr: %#v \n", res, err)
		return nil, fmt.Errorf("faild to get data. status: %s", res.Status)
	}

	var accountBalances []*models.AccountBalance
	if err := decodeBody(res, &accountBalances); err != nil {
		return nil, err
	}

	return accountBalances, nil
}
