package quoinex

import (
	"context"
	"fmt"
	"github.com/jjjjpppp/quoinex-go-client/v2/models"
	"strconv"
)

func (c *Client) GetExecutionsByTimestamp(ctx context.Context, productID int, limit int, timestamp int) ([]*models.ExecutionsModels, error) {
	spath := "/executions"
	req, err := c.newRequest(ctx, "GET", spath, nil,
		&map[string]string{
			"product_id": strconv.Itoa(productID),
			"limit":      strconv.Itoa(limit),
			"timestamp":  strconv.Itoa(timestamp)})
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get data. status: %s", res.Status)
	}

	var executions []*models.ExecutionsModels
	if err := decodeBody(res, &executions); err != nil {
		return nil, err
	}

	return executions, nil
}

func (c *Client) GetExecutions(ctx context.Context, productID int, limit int, page int) (*models.Executions, error) {
	spath := "/executions"
	req, err := c.newRequest(ctx, "GET", spath, nil,
		&map[string]string{
			"product_id": strconv.Itoa(productID),
			"limit":      strconv.Itoa(limit),
			"page":       strconv.Itoa(page)})
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get data. status: %s", res.Status)
	}

	var executions models.Executions
	if err := decodeBody(res, &executions); err != nil {
		return nil, err
	}

	return &executions, nil
}

func (c *Client) GetOwnExecutions(ctx context.Context, productID int) (*models.Executions, error) {
	spath := "/executions/me"
	req, err := c.newRequest(ctx, "GET", spath, nil,
		&map[string]string{
			"product_id": strconv.Itoa(productID),
		})
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get data. status: %s", res.Status)
	}

	var executions models.Executions
	if err := decodeBody(res, &executions); err != nil {
		return nil, err
	}

	return &executions, nil
}
