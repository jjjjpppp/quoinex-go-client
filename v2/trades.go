package quoinex

import (
	"context"
	"fmt"
	"github.com/jjjjpppp/quoinex-go-client/v2/models"
	"strings"
)

func (c *Client) GetTrades(ctx context.Context, fundingCurrency, status string) (*models.Trades, error) {
	spath := "/trades"
	queryParam := &map[string]string{
		"funding_currency": fundingCurrency,
		"status":           status}
	res, err := c.sendRequest(ctx, "GET", spath, nil, queryParam)
	if err != nil {
		return nil, err
	}

	var trades models.Trades
	if err := decodeBody(res, &trades); err != nil {
		return nil, err
	}

	return &trades, nil
}

func (c *Client) CloseTrade(ctx context.Context, tradeID int, closedQuantity float64) (*models.Trade, error) {
	spath := fmt.Sprintf("/trades/%d/close", tradeID)
	bodyTemplate := `{"closed_quantity":%f}`
	body := fmt.Sprintf(bodyTemplate, closedQuantity)
	res, err := c.sendRequest(ctx, "PUT", spath, strings.NewReader(body), nil)
	if err != nil {
		return nil, err
	}

	var trade models.Trade
	if err := decodeBody(res, &trade); err != nil {
		return nil, err
	}

	return &trade, nil
}

func (c *Client) CloseAllTrade(ctx context.Context, side string) ([]*models.Trade, error) {
	spath := "/trades/close_all"
	bodyTemplate := `{"side":"%s"}`
	body := fmt.Sprintf(bodyTemplate, side)
	res, err := c.sendRequest(ctx, "PUT", spath, strings.NewReader(body), nil)
	if err != nil {
		return nil, err
	}

	var trades []*models.Trade
	if err := decodeBody(res, &trades); err != nil {
		return nil, err
	}

	return trades, nil
}

func (c *Client) UpdateTrade(ctx context.Context, tradeID, stopLoss, takeProfit int) (*models.Trade, error) {
	spath := fmt.Sprintf("/trades/%d", tradeID)
	bodyTemplate :=
		`{
			"trade": {
				"stop_loss":"%d",
				"take_profit":"%d"
			}
		}`
	body := fmt.Sprintf(bodyTemplate, stopLoss, takeProfit)
	res, err := c.sendRequest(ctx, "PUT", spath, strings.NewReader(body), nil)
	if err != nil {
		return nil, err
	}

	var trade models.Trade
	if err := decodeBody(res, &trade); err != nil {
		return nil, err
	}

	return &trade, nil
}

func (c *Client) GetTradesLoans(ctx context.Context, tradeID int) ([]*models.Loan, error) {
	spath := fmt.Sprintf("/trades/%d/loans", tradeID)
	res, err := c.sendRequest(ctx, "GET", spath, nil, nil)
	if err != nil {
		return nil, err
	}

	var loans []*models.Loan
	if err := decodeBody(res, &loans); err != nil {
		return nil, err
	}

	return loans, nil
}
