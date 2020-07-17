package quoinex

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jjjjpppp/quoinex-go-client/v2/models"
)

func (c *Client) GetAnOrder(ctx context.Context, orderID uint64) (*models.Order, error) {
	spath := fmt.Sprintf("/orders/%d", orderID)
	res, err := c.sendRequest(ctx, "GET", spath, nil, nil)
	if err != nil {
		return nil, err
	}

	var order models.Order
	if err := decodeBody(res, &order); err != nil {
		return nil, err
	}

	return &order, nil
}

func (c *Client) GetOrders(ctx context.Context, productID, withDetails int, fundingCurrency, status string) (*models.Orders, error) {
	spath := "/orders"
	queryParam := &map[string]string{
		"product_id":       strconv.Itoa(productID),
		"with_details":     strconv.Itoa(withDetails),
		"status":           status,
		"funding_currency": fundingCurrency}
	res, err := c.sendRequest(ctx, "GET", spath, nil, queryParam)
	if err != nil {
		return nil, err
	}

	var orders models.Orders
	if err := decodeBody(res, &orders); err != nil {
		return nil, err
	}

	return &orders, nil
}

func (c *Client) CreateAnOrder(ctx context.Context, orderType, side, quantity, price, priceRange string, productID int) (*models.Order, error) {
	spath := "/orders/"
	bodyTemplate :=
		`{
			"order": {
				"order_type":"%s",
				"product_id":%d,
				"side":"%s",
				"quantity":"%s",
				"price":"%s",
				"price_range":"%s"
			}
		}`
	body := fmt.Sprintf(bodyTemplate, orderType, productID, side, quantity, price, priceRange)
	res, err := c.sendRequest(ctx, "POST", spath, strings.NewReader(body), nil)
	if err != nil {
		return nil, err
	}

	var order models.Order
	if err := decodeBody(res, &order); err != nil {
		return nil, err
	}

	return &order, nil
}

func (c *Client) CancelAnOrder(ctx context.Context, orderID uint64) (*models.Order, error) {
	spath := fmt.Sprintf("/orders/%d/cancel", orderID)
	res, err := c.sendRequest(ctx, "PUT", spath, nil, nil)
	if err != nil {
		return nil, err
	}

	var order models.Order
	if err := decodeBody(res, &order); err != nil {
		return nil, err
	}

	return &order, nil
}

func (c *Client) EditALiveOrder(ctx context.Context, orderID uint64, quantity, price string) (*models.Order, error) {
	spath := fmt.Sprintf("/orders/%d", orderID)
	bodyTemplate :=
		`{
			"order": {
				"quantity":"%s",
				"price":"%s",
			}
		}`
	body := fmt.Sprintf(bodyTemplate, quantity, price)
	res, err := c.sendRequest(ctx, "PUT", spath, strings.NewReader(body), nil)
	if err != nil {
		return nil, err
	}

	var order models.Order
	if err := decodeBody(res, &order); err != nil {
		return nil, err
	}

	return &order, nil
}

func (c *Client) GetAnOrderTrades(ctx context.Context, orderID uint64) ([]*models.Trade, error) {
	spath := fmt.Sprintf("/orders/%d/trades", orderID)
	res, err := c.sendRequest(ctx, "GET", spath, nil, nil)
	if err != nil {
		return nil, err
	}

	var trades []*models.Trade
	if err := decodeBody(res, &trades); err != nil {
		return nil, err
	}

	return trades, nil
}
