package quoinex

import (
	"context"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/jjjjpppp/quoinex-go-client/v2/models"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	type Param struct {
		apiToken  string
		apiSecret string
		logger    *log.Logger
	}
	type Expect struct {
		client *Client
		err    error
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{apiToken: "", apiSecret: "secret", logger: nil},
			expect: Expect{client: nil, err: fmt.Errorf("apiTokenID is not set")},
		},
		// test case 2
		{
			param:  Param{apiToken: "apiToken", apiSecret: "", logger: nil},
			expect: Expect{client: nil, err: fmt.Errorf("apiSecret is not set")},
		},
		// test case 3
		{
			param: Param{apiToken: "apiToken", apiSecret: "apiSecret", logger: nil},
			expect: Expect{client: &Client{
				ApiTokenID: "apiToken",
				ApiSecret:  "apiSecret",
				HTTPClient: &http.Client{Timeout: time.Duration(10) * time.Second},
				Logger:     log.New(ioutil.Discard, "", log.LstdFlags),
			}, err: nil},
		},
	}
	for _, c := range cases {
		client, e := NewClient(c.param.apiToken, c.param.apiSecret, c.param.logger)
		if client == nil && e.Error() != c.expect.err.Error() {
			t.Errorf("Worng err. test set is %+v", c)
		}
		if client == nil {
			t.Logf("client is nil. skip this case.: test set: %+v", c)
			continue
		}
		if client.ApiTokenID != c.expect.client.ApiTokenID {
			t.Errorf("Worng apiToken. test set: %+v", c)
		}
		if client.ApiSecret != c.expect.client.ApiSecret {
			t.Errorf("Worng ApiSecret. test set: %+v", c)
		}
		if reflect.TypeOf(client.HTTPClient) != reflect.TypeOf(c.expect.client.HTTPClient) {
			t.Errorf("Worng HTTPClient. test set: %+v", c)
		}
		if reflect.TypeOf(client.Logger) != reflect.TypeOf(c.expect.client.Logger) {
			t.Errorf("Worng Logger. test set: %+v", c)
		}
	}
}
func TestNewRequest(t *testing.T) {
	type Param struct {
		method     string
		spath      string
		queryParam *map[string]string
	}
	type Expect struct {
		method string
		url    string
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{method: "GET", spath: "product/1", queryParam: nil},
			expect: Expect{method: "GET", url: "https://api.quoine.com/product/1"},
		},
		// test case 2
		{
			param:  Param{method: "GET", spath: "product/1", queryParam: &map[string]string{"product_id": "1", "limit": "1", "page": "1"}},
			expect: Expect{method: "GET", url: "https://api.quoine.com/product/1?limit=1&page=1&product_id=1"},
		},
	}

	for _, c := range cases {
		client, _ := NewClient("apiTokenID", "secret", nil)
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		req, _ := client.newRequest(ctx, c.param.method, c.param.spath, nil, c.param.queryParam)
		if req.Method != c.expect.method {
			t.Errorf("Worng method. case: %+v", c)
		}
		if len(req.Header["X-Quoine-Auth"]) < 1 {
			t.Errorf("Worng Header. case: %+v", c)
		}
		if req.URL.String() != c.expect.url {
			t.Errorf("Worng URL case: %+v, actual: %+v", c, req.URL.String())
		}
	}
}

func TestGetAnOrder(t *testing.T) {
	type Param struct {
		orderID      int
		jsonResponse string
	}
	type Expect struct {
		path  string
		order *models.Order
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param: Param{orderID: 1, jsonResponse: getOrderJsonResponse()},
			expect: Expect{path: "/orders/1", order: &models.Order{
				ID:                   2157479,
				OrderType:            "limit",
				Quantity:             "0.01",
				DiscQuantity:         "0.0",
				IcebergTotalQuantity: "0.0",
				Side:                 "sell",
				FilledQuantity:       "0.01",
				Price:                "500.0",
				CreatedAt:            1462123639,
				UpdatedAt:            1462123639,
				Status:               "filled",
				LeverageLevel:        2,
				SourceExchange:       "QUOINE",
				ProductID:            1,
				ProductCode:          "CASH",
				FundingCurrency:      "USD",
				CurrencyPairCode:     "BTCUSD",
				OrderFee:             "0.0",
				Executions: models.OrderExecutions{
					{
						ID:        4566133,
						Quantity:  "0.01",
						Price:     "500.0",
						TakerSide: "buy",
						MySide:    "sell",
						CreatedAt: 1465396785,
					},
				},
			}},
		},
		// test case 2
	}
	for _, c := range cases {
		// preparing test server
		ts := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != c.expect.path {
					t.Fatalf("worng URL")
				}
				// set expected json
				w.Header().Set("content-Type", "text")
				fmt.Fprintf(w, c.param.jsonResponse)
				return
			},
		))
		defer ts.Close()

		client, _ := NewClient("apiTokenID", "secret", nil)
		client.testServer = ts
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		order, _ := client.GetAnOrder(ctx, 1)
		if !cmp.Equal(order, c.expect.order) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(order, c.expect.order))

		}

	}
}

func getOrderJsonResponse() string {
	return `{
  	"id": 2157479,
  	"order_type": "limit",
  	"quantity": "0.01",
  	"disc_quantity": "0.0",
  	"iceberg_total_quantity": "0.0",
  	"side": "sell",
  	"filled_quantity": "0.01",
  	"price": "500.0",
  	"created_at": 1462123639,
  	"updated_at": 1462123639,
  	"status": "filled",
  	"leverage_level": 2,
  	"source_exchange": "QUOINE",
  	"product_id": 1,
  	"product_code": "CASH",
  	"funding_currency": "USD",
  	"currency_pair_code": "BTCUSD",
  	"order_fee": "0.0",
  	"executions": [
  	  {
  	    "id": 4566133,
  	    "quantity": "0.01",
  	    "price": "500.0",
  	    "taker_side": "buy",
  	    "my_side": "sell",
  	    "created_at": 1465396785
  	  }
  	]
	}`
}

//func TestGetProducts(t *testing.T) {
//	// preparing test server
//	ts := httptest.NewServer(http.HandlerFunc(
//		func(w http.ResponseWriter, r *http.Request) {
//			if r.URL.Path != "/products" {
//				t.Fatalf("worng URL")
//			}
//			w.Header().Set("content-Type", "text")
//			fmt.Fprintf(w, getProductsJsonResponse())
//			return
//		},
//	))
//	defer ts.Close()
//
//	c, _ := NewClient("apiTokenID", "secret", nil)
//	c.testServer = ts
//	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
//	products, err := c.GetProducts(ctx)
//	t.Logf("products: %v", products)
//	if err != nil {
//		t.Errorf("Worng err is not nil. Error: %v,", err)
//	}
//}

//func getProductsJsonResponse() string {
//	return `[
//    {
//        "id": 5,
//        "product_type": "CurrencyPair",
//        "code": "CASH",
//        "name": "CASH Trading",
//        "market_ask": "48203.05",
//        "market_bid": "48188.15",
//        "indicator": -1,
//        "currency": "JPY",
//        "currency_pair_code": "BTCJPY",
//        "symbol": "¥",
//        "fiat_minimum_withdraw": "1500.0",
//        "pusher_channel": "product_cash_btcjpy_5",
//        "taker_fee": "0.0",
//        "maker_fee": "0.0",
//        "low_market_bid": "47630.99",
//        "high_market_ask": "48396.71",
//        "volume_24h": "2915.627366519999999998",
//        "last_price_24h": "48217.2",
//        "last_traded_price": "48203.05",
//        "last_traded_quantity": "1.0",
//        "quoted_currency": "JPY",
//        "base_currency": "BTC",
//        "exchange_rate": "0.009398151671149725"
//    },
//  ]`
//}
