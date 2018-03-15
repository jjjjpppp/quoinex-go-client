package quoinex

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jjjjpppp/quoinex-go-client/v2/models"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"runtime"
	"strconv"
	"time"
)

const (
	baseUrl string = "https://api.quoine.com"
	version string = "0.0.1"
)

type Client struct {
	URL        *url.URL
	ApiTokenID string
	ApiSecret  string
	HTTPClient *http.Client
	Logger     *log.Logger
}

func NewClient(apiTokenID string, apiSecret string, logger *log.Logger) (*Client, error) {
	if len(apiTokenID) == 0 {
		return nil, fmt.Errorf("apiTokenID is not set")
	}

	if len(apiSecret) == 0 {
		return nil, fmt.Errorf("apiSecret is not set")
	}

	url, err := url.ParseRequestURI(baseUrl)
	if err != nil {
		return nil, err
	}

	var discardLogger = log.New(ioutil.Discard, "", log.LstdFlags)
	if logger == nil {
		logger = discardLogger
	}

	client := &http.Client{Timeout: time.Duration(10) * time.Second}
	return &Client{URL: url, ApiTokenID: apiTokenID, ApiSecret: apiSecret, HTTPClient: client, Logger: logger}, nil

}

func (c *Client) GetExecutions(ctx context.Context, productID int, limit int, page int) (*models.Executions, error) {
	spath := fmt.Sprintf("/executions")
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

func (c *Client) GetOrderBook(ctx context.Context, id int) (*models.PriceLevels, error) {
	spath := fmt.Sprintf("/products/%d/price_levels", id)
	req, err := c.newRequest(ctx, "GET", spath, nil, nil)
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

	var priceLevels models.PriceLevels
	if err := decodeBody(res, &priceLevels); err != nil {
		return nil, err
	}

	return &priceLevels, nil
}

func (c *Client) GetProducts(ctx context.Context) (*[]models.Product, error) {
	spath := fmt.Sprintf("/products")
	req, err := c.newRequest(ctx, "GET", spath, nil, nil)
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

	var products []models.Product
	if err := decodeBody(res, &products); err != nil {
		return nil, err
	}

	return &products, nil
}

func (c *Client) GetProduct(ctx context.Context, productID int) (*models.Product, error) {
	spath := fmt.Sprintf("/products/%d", productID)
	req, err := c.newRequest(ctx, "GET", spath, nil, nil)
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

	var product models.Product
	if err := decodeBody(res, &product); err != nil {
		return nil, err
	}

	return &product, nil
}

func (c *Client) GetOrder(ctx context.Context, orderID int) (*models.Order, error) {
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

func (c *Client) newRequest(ctx context.Context, method, spath string, body io.Reader, queryParam *map[string]string) (*http.Request, error) {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)
	// build QueryParameter
	if queryParam != nil {
		q := u.Query()
		for k, v := range *queryParam {
			q.Set(k, v)
		}
		u.RawQuery = q.Encode()
	}

	userAgent := fmt.Sprintf("GoClient/%s (%s)", version, runtime.Version())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"path":     spath,
		"nonce":    time.Now().Unix(),
		"token_id": c.ApiTokenID,
	})

	tokenString, err := token.SignedString([]byte(c.ApiSecret))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Quoine-API-Version", "2")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("X-Quoine-Auth", tokenString)

	return req, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}
