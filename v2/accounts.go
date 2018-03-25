package quoinex

import (
	"context"
	"fmt"
	"github.com/jjjjpppp/quoinex-go-client/v2/models"
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
