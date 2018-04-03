package quoinex

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/jjjjpppp/quoinex-go-client/v2/models"
	"github.com/jjjjpppp/quoinex-go-client/v2/testutil"
	"testing"
	"time"
)

func TestGetAnOrder(t *testing.T) {
	type Param struct {
		orderID      int
		jsonResponse string
	}
	type Expect struct {
		path   string
		method string
		body   string
		order  *models.Order
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{orderID: 1, jsonResponse: testutil.GetOrderJsonResponse()},
			expect: Expect{path: "/orders/1", method: "GET", body: "", order: testutil.GetExpectedOrderModel()},
		},
		// test case 2
	}
	for _, c := range cases {
		// preparing test server
		ts := testutil.GenerateTestServer(t, c.expect.path, c.expect.method, c.expect.body, c.param.jsonResponse)
		defer ts.Close()

		client, _ := NewClient("apiTokenID", "secret", nil)
		client.testServer = ts
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
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
		path   string
		method string
		body   string
		e      *models.Orders
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{productID: 1, withDetails: 1, fundingCurrency: "USD", status: "ok", jsonResponse: testutil.GetOrdersJsonResponse()},
			expect: Expect{path: "/orders?funding_currency=USD&product_id=1&status=ok&with_details=1", method: "GET", body: "", e: testutil.GetExpectedOrdersModel()},
		},
		// test case 2
	}
	for _, c := range cases {
		// preparing test server
		ts := testutil.GenerateTestServer(t, c.expect.path, c.expect.method, c.expect.body, c.param.jsonResponse)
		defer ts.Close()

		client, _ := NewClient("apiTokenID", "secret", nil)
		client.testServer = ts
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
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
		path   string
		method string
		body   string
		a      *models.Order
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{orderType: "limit", productID: 1, side: "sell", quantity: "0.01", price: "500.0", priceRange: "", jsonResponse: testutil.GetCreateAnOrderJsonResponse()},
			expect: Expect{path: "/orders", body: testutil.GetExpectedCreateAnOrderRequestBody(), method: "POST", a: testutil.GetExpectedCreateAnOrderModel()},
		},
		// test case 2
	}
	for _, c := range cases {
		// preparing test server
		ts := testutil.GenerateTestServer(t, c.expect.path, c.expect.method, c.expect.body, c.param.jsonResponse)
		defer ts.Close()

		client, _ := NewClient("apiTokenID", "secret", nil)
		client.testServer = ts
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
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
		body   string
		order  *models.Order
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{orderID: 2157474, jsonResponse: testutil.GetCancelAnOrderJsonResponse()},
			expect: Expect{path: "/orders/2157474/cancel", method: "PUT", body: "", order: testutil.GetExpectedCancelAnOrderModel()},
		},
		// test case 2
	}
	for _, c := range cases {
		// preparing test server
		ts := testutil.GenerateTestServer(t, c.expect.path, c.expect.method, c.expect.body, c.param.jsonResponse)
		defer ts.Close()

		client, _ := NewClient("apiTokenID", "secret", nil)
		client.testServer = ts
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		order, _ := client.CancelAnOrder(ctx, c.param.orderID)
		if !cmp.Equal(order, c.expect.order) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(order, c.expect.order))
		}
	}
}

func TestEditALiveOrder(t *testing.T) {
	type Param struct {
		orderID      int
		quantity     string
		price        string
		jsonResponse string
	}
	type Expect struct {
		path   string
		method string
		body   string
		order  *models.Order
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{orderID: 2157474, quantity: "0.02", price: "520.0", jsonResponse: testutil.GetEditALiveOrderJsonResponse()},
			expect: Expect{path: "/orders/2157474", method: "PUT", body: "price=520.0&quantity=0.02", order: testutil.GetExpectedEditALiveOrderModel()},
		},
		// test case 2
	}
	for _, c := range cases {
		// preparing test server
		ts := testutil.GenerateTestServer(t, c.expect.path, c.expect.method, c.expect.body, c.param.jsonResponse)
		defer ts.Close()

		client, _ := NewClient("apiTokenID", "secret", nil)
		client.testServer = ts
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		order, _ := client.EditALiveOrder(ctx, c.param.orderID, c.param.quantity, c.param.price)
		if !cmp.Equal(order, c.expect.order) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(order, c.expect.order))
		}
	}
}

func TestGetAnOrderTrades(t *testing.T) {
	type Param struct {
		orderID      int
		jsonResponse string
	}
	type Expect struct {
		path   string
		method string
		body   string
		trades []*models.Trade
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{orderID: 1, jsonResponse: testutil.GetOrderTradesJsonResponse()},
			expect: Expect{path: "/orders/1/trades", method: "GET", body: "", trades: testutil.GetExpectedOrderTradesModel()},
		},
		// test case 2
	}
	for _, c := range cases {
		// preparing test server
		ts := testutil.GenerateTestServer(t, c.expect.path, c.expect.method, c.expect.body, c.param.jsonResponse)
		defer ts.Close()

		client, _ := NewClient("apiTokenID", "secret", nil)
		client.testServer = ts
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		trades, _ := client.GetAnOrderTrades(ctx, 1)
		if !cmp.Equal(trades, c.expect.trades) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(trades, c.expect.trades))
		}
	}
}
