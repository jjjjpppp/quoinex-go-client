package quoinex

import (
	"context"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/jjjjpppp/quoinex-go-client/v2/models"
	"github.com/jjjjpppp/quoinex-go-client/v2/testutil"
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
			param:  Param{orderID: 1, jsonResponse: testutil.GetOrderJsonResponse()},
			expect: Expect{path: "/orders/1", order: testutil.GetExpectedOrderModel()},
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

func TestGetOrders(t *testing.T) {
	type Param struct {
		productID       int
		withDetails     int
		fundingCurrency string
		status          string
		jsonResponse    string
	}
	type Expect struct {
		path string
		e    *models.Orders
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{productID: 1, withDetails: 1, fundingCurrency: "USD", status: "ok", jsonResponse: testutil.GetOrdersJsonResponse()},
			expect: Expect{path: "/orders", e: testutil.GetExpectedOrdersModel()},
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
		r, _ := client.GetOrders(ctx, c.param.productID, c.param.withDetails, c.param.fundingCurrency, c.param.status)
		if !cmp.Equal(r, c.expect.e) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(r, c.expect.e))
		}
	}
}

func TestGetProducts(t *testing.T) {
	type Param struct {
		jsonResponse string
	}
	type Expect struct {
		path     string
		products []*models.Product
	}

	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{jsonResponse: testutil.GetProductsJsonResponse()},
			expect: Expect{path: "/products", products: testutil.GetExpectedProductsModel()},
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
		products, err := client.GetProducts(ctx)
		if err != nil {
			t.Errorf("Error. %+v", err)
		}
		if !cmp.Equal(products, c.expect.products) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(products, c.expect.products))

		}

	}
}

func TestGetProduct(t *testing.T) {
	type Param struct {
		productID    int
		jsonResponse string
	}
	type Expect struct {
		path    string
		product *models.Product
	}

	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{productID: 1, jsonResponse: testutil.GetProductJsonResponse()},
			expect: Expect{path: "/products/1", product: testutil.GetExpectedProductmodel()},
		},
		// test case 2
	}
	for _, c := range cases {
		// preparing test server
		ts := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != c.expect.path {
					t.Fatalf("worng URL. actual:%+v, expect:%+v", r.URL.Path, c.expect.path)
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
		product, err := client.GetProduct(ctx, c.param.productID)
		if err != nil {
			t.Errorf("Error. %+v", err)
		}
		if !cmp.Equal(product, c.expect.product) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(product, c.expect.product))
		}
	}
}

func TestGetOrderBook(t *testing.T) {
	type Param struct {
		productID    int
		jsonResponse string
	}
	type Expect struct {
		path        string
		priceLevels *models.PriceLevels
	}

	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{productID: 1, jsonResponse: testutil.GetOrderBookJsonResponse()},
			expect: Expect{path: "/products/1/price_levels", priceLevels: testutil.GetExpectedOrderBookModel()},
		},
		// test case 2
	}
	for _, c := range cases {
		// preparing test server
		ts := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != c.expect.path {
					t.Fatalf("worng URL. actual:%+v, expect:%+v", r.URL.Path, c.expect.path)
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
		priceLevels, err := client.GetOrderBook(ctx, c.param.productID)
		if err != nil {
			t.Errorf("Error. %+v", err)
		}
		if !cmp.Equal(priceLevels, c.expect.priceLevels) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(priceLevels, c.expect.priceLevels))
		}
	}
}

func TestGetExecutions(t *testing.T) {
	type Param struct {
		productID    int
		limit        int
		page         int
		jsonResponse string
	}
	type Expect struct {
		path       string
		executions *models.Executions
	}

	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{productID: 1, limit: 1, page: 1, jsonResponse: testutil.GetExecutionsJsonResponse()},
			expect: Expect{path: "/executions?limit=1&page=1&product_id=1", executions: testutil.GetExpectedExecutionsModel()},
		},
		// test case 2
	}
	for _, c := range cases {
		// preparing test server
		ts := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.URL.RequestURI() != c.expect.path {
					t.Fatalf("worng URL. actual:%+v, expect:%+v", r.URL.RequestURI(), c.expect.path)
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
		executions, err := client.GetExecutions(ctx, c.param.productID, c.param.limit, c.param.page)
		if err != nil {
			t.Errorf("Error. %+v", err)
		}
		if !cmp.Equal(executions, c.expect.executions) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(executions, c.expect.executions))
		}
	}
}

func TestGetExecutionsByTimestamp(t *testing.T) {
	type Param struct {
		productID    int
		limit        int
		timestamp    int
		jsonResponse string
	}
	type Expect struct {
		path       string
		executions []*models.ExecutionsModels
	}

	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{productID: 1, limit: 2, timestamp: 1430630863, jsonResponse: testutil.GetExecutionsByTimestampJsonResponse()},
			expect: Expect{path: "/executions?limit=2&product_id=1&timestamp=1430630863", executions: testutil.GetExpectedExecutionsByTimestampModel()},
		},
		// test case 2
	}
	for _, c := range cases {
		// preparing test server
		ts := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.URL.RequestURI() != c.expect.path {
					t.Fatalf("worng URL. actual:%+v, expect:%+v", r.URL.RequestURI(), c.expect.path)
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
		executions, err := client.GetExecutionsByTimestamp(ctx, c.param.productID, c.param.limit, c.param.timestamp)
		if err != nil {
			t.Errorf("Error. %+v", err)
		}
		if !cmp.Equal(executions, c.expect.executions) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(executions, c.expect.executions))
		}
	}
}

func TestGetInterestRates(t *testing.T) {
	type Param struct {
		currency     string
		jsonResponse string
	}
	type Expect struct {
		path string
		a    *models.InterestRates
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{currency: "USD", jsonResponse: testutil.GetInterestRatesJsonResponse()},
			expect: Expect{path: "/ir_ladders/USD", a: testutil.GetExpectedInterestRatesModel()},
		},
		// test case 2
	}
	for _, c := range cases {
		// preparing test server
		ts := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.URL.RequestURI() != c.expect.path {
					t.Fatalf("worng URL. actual:%+v, expect:%+v", r.URL.RequestURI(), c.expect.path)
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
		r, err := client.GetInterestRates(ctx, c.param.currency)
		if err != nil {
			t.Errorf("Error. %+v", err)
		}
		if !cmp.Equal(r, c.expect.a) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(r, c.expect.a))
		}
	}
}

func TestCreateAnOrder(t *testing.T) {
	type Param struct {
		orderType    string
		productID    int
		side         string
		quantity     string
		price        string
		priceRange   string
		jsonResponse string
	}
	type Expect struct {
		path string
		body string
		a    *models.Order
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{orderType: "limit", productID: 1, side: "sell", quantity: "0.01", price: "500.0", priceRange: "", jsonResponse: testutil.GetCreateAnOrderJsonResponse()},
			expect: Expect{path: "/orders", body: "order_type=limit&price=500.0&price_range=&product_id=1&quantity=0.01&side=sell", a: testutil.GetExpectedCreateAnOrderModel()},
		},
		// test case 2
	}
	for _, c := range cases {
		// preparing test server
		ts := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.URL.RequestURI() != c.expect.path {
					t.Fatalf("worng URL. actual:%+v, expect:%+v", r.URL.RequestURI(), c.expect.path)
				}
				// Read body
				b, err := ioutil.ReadAll(r.Body)
				s := string(b)
				defer r.Body.Close()
				if err != nil {
					t.Errorf("Worng body. err:%+v", err)
				}
				if s != c.expect.body {
					t.Errorf("Worng body. actual: %+v, expect:%+v", s, c.expect.body)
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
		r, err := client.CreateAnOrder(ctx, c.param.orderType, c.param.side, c.param.quantity, c.param.price, c.param.priceRange, c.param.productID)
		if err != nil {
			t.Errorf("Error. %+v", err)
		}
		if !cmp.Equal(r, c.expect.a) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(r, c.expect.a))
		}
	}
}
