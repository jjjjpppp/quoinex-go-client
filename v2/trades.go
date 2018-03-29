package quoinex

import (
	"context"
	"fmt"
	"github.com/jjjjpppp/quoinex-go-client/v2/models"
	"net/url"
	"strconv"
	"strings"
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

func (c *Client) CloseTrade(ctx context.Context, tradeID int, closedQuantity float64) (*models.Trade, error) {
	spath := fmt.Sprintf("/trades/%d/close", tradeID)
	values := url.Values{}
	values.Add("closed_quantity", strconv.FormatFloat(closedQuantity, 'f', 4, 64))
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

	var trade models.Trade
	if err := decodeBody(res, &trade); err != nil {
		return nil, err
	}

	return &trade, nil
}

func (c *Client) CloseAllTrade(ctx context.Context, side string) ([]*models.Trade, error) {
	spath := fmt.Sprintf("/trades/close_all")
	values := url.Values{}
	values.Add("side", side)
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

	var trades []*models.Trade
	if err := decodeBody(res, &trades); err != nil {
		return nil, err
	}

	return trades, nil
}

func (c *Client) UpdateTrade(ctx context.Context, tradeID, stopLoss, takeProfit int) (*models.Trade, error) {
	spath := fmt.Sprintf("/trades/%d", tradeID)
	values := url.Values{}
	values.Add("stop_loss", strconv.Itoa(stopLoss))
	values.Add("take_profit", strconv.Itoa(takeProfit))
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

	var trade models.Trade
	if err := decodeBody(res, &trade); err != nil {
		return nil, err
	}

	return &trade, nil
}
