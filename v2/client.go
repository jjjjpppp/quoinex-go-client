package quoinex

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jjjjpppp/quoinex_go_client/v2/models"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"runtime"
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

func (c *Client) GetOrder(ctx context.Context, orderID int) (*models.Order, error) {
	spath := fmt.Sprintf("/orders/%d", orderID)
	req, err := c.newRequest(ctx, "GET", spath, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.Status != "200" {
		return nil, fmt.Errorf("faild get data. status: %s", res.Status)
	}

	var order models.Order
	if err := decodeBody(res, &order); err != nil {
		return nil, err
	}

	return &order, nil
}

func (c *Client) newRequest(ctx context.Context, method, spath string, body io.Reader) (*http.Request, error) {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)
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
