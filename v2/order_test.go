package quoinex

import (
	"context"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/jjjjpppp/quoinex-go-client/v2/models"
	"github.com/jjjjpppp/quoinex-go-client/v2/testutil"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

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

func TestCancelAnOrder(t *testing.T) {
	type Param struct {
		orderID      int
		jsonResponse string
	}
	type Expect struct {
		path   string
		method string
		order  *models.Order
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{orderID: 2157474, jsonResponse: testutil.GetCancelAnOrderJsonResponse()},
			expect: Expect{path: "/orders/2157474/cancel", method: "PUT", order: testutil.GetExpectedCancelAnOrderModel()},
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
				if r.Method != "PUT" {
					t.Fatalf("worng Method. actual:%+v, expect:%+v", r.Method, c.expect.method)
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
		order, _ := client.CancelAnOrder(ctx, c.param.orderID)
		if !cmp.Equal(order, c.expect.order) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(order, c.expect.order))

		}

	}
}
