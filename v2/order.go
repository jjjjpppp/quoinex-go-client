package quoinex

import (
	"context"
	"fmt"
	"github.com/jjjjpppp/quoinex-go-client/v2/models"
	"net/url"
	"strconv"
	"strings"
)

func (c *Client) GetAnOrder(ctx context.Context, orderID int) (*models.Order, error) {
	spath := fmt.Sprintf("/orders/%d", orderID)
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

	var order models.Order
	if err := decodeBody(res, &order); err != nil {
		return nil, err
	}

	return &order, nil
}

func (c *Client) GetOrders(ctx context.Context, productID, withDetails int, fundingCurrency, status string) (*models.Orders, error) {
	spath := fmt.Sprintf("/orders")
	queryParam := &map[string]string{
		"product_id":       strconv.Itoa(productID),
		"with_details":     strconv.Itoa(withDetails),
		"status":           status,
		"funding_currency": fundingCurrency}
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

	var orders models.Orders
	if err := decodeBody(res, &orders); err != nil {
		return nil, err
	}

	return &orders, nil
}

func (c *Client) CreateAnOrder(ctx context.Context, orderType, side, quantity, price, priceRange string, productID int) (*models.Order, error) {
	spath := fmt.Sprintf("/orders/")
	values := url.Values{}
	values.Add("order_type", orderType)
	values.Add("product_id", strconv.Itoa(productID))
	values.Add("side", side)
	values.Add("quantity", quantity)
	values.Add("price", price)
	values.Add("price_range", priceRange)
	req, err := c.newRequest(ctx, "POST", spath, strings.NewReader(values.Encode()), nil)
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

	var order models.Order
	if err := decodeBody(res, &order); err != nil {
		return nil, err
	}

	return &order, nil
}

func (c *Client) CancelAnOrder(ctx context.Context, orderID int) (*models.Order, error) {
	spath := fmt.Sprintf("/orders/%d/cancel", orderID)
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

	var order models.Order
	if err := decodeBody(res, &order); err != nil {
		return nil, err
	}

	return &order, nil
}
