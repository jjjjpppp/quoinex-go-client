package quoinex

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClientApiTokenError(t *testing.T) {
	if _, e := NewClient("", "secret", nil); e == nil {
		t.Error("err should be returned")
	}
}

func TestNewClientApiSecretError(t *testing.T) {
	if _, e := NewClient("apiToken", "", nil); e == nil {
		t.Error("err should be returned")
	}
}

func TestNewClientSuccess(t *testing.T) {
	c, e := NewClient("apiTokenID", "secret", nil)
	if e != nil {
		t.Error("err should be nil")
	}
	if c.ApiTokenID != "apiTokenID" {
		t.Error("Worng apiToken")
	}
	if c.ApiSecret != "secret" {
		t.Error("Worng ApiSecret")
	}
	if c.HTTPClient == nil {
		t.Error("HttpClient is nil")
	}
	if c.Logger == nil {
		t.Error("Logger is nil")
	}
}

func TestNewRequestSuccessNoQueryParam(t *testing.T) {
	c, _ := NewClient("apiTokenID", "secret", nil)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	req, err := c.newRequest(ctx, "GET", "product/1", nil, nil)
	if err != nil {
		t.Error("Error is occered")
	}
	if req.Method != "GET" {
		t.Error("Worng method")
	}
	if len(req.Header["X-Quoine-Auth"]) < 1 {
		t.Error("Worng Header")
	}
	if req.URL.String() != "https://api.quoine.com/product/1" {
		t.Errorf("Worng URL : %+v", req.URL.String())
	}
}

func TestNewRequestSuccessWithQueryParam(t *testing.T) {
	c, _ := NewClient("apiTokenID", "secret", nil)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	queryParam := &map[string]string{"product_id": "1", "limit": "1", "page": "1"}
	req, err := c.newRequest(ctx, "GET", "product/1", nil, queryParam)
	if err != nil {
		t.Error("Error is occered")
	}
	if req.Method != "GET" {
		t.Error("Worng method")
	}
	if len(req.Header["X-Quoine-Auth"]) < 1 {
		t.Error("Worng Header")
	}
	aq := req.URL.Query()
	for k, v := range *queryParam {
		if aq[k][0] != v {
			t.Errorf("Worng Query Parameter k:%+v, v:%+v", k, v)
		}
	}
}

func TestGetAnOrder(t *testing.T) {
	// preparing test server
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/orders/1" {
				t.Fatalf("worng URL")
			}
			//if r.URL.Query().Get("greet") != "Hello" {
			//	t.Fatalf("worng Query Params")
			//}
			// set expected json
			w.Header().Set("content-Type", "text")
			fmt.Fprintf(w, getOrderJsonResponse())
			return
		},
	))
	defer ts.Close()

	c, _ := NewClient("apiTokenID", "secret", nil)
	c.testServer = ts
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	order, err := c.GetAnOrder(ctx, 1)
	if err != nil {
		t.Errorf("Error is not nil. Order : %+v, Error : %+v", order, err)
	}
	if order.ID != 2157479 {
		t.Errorf("Worng ID. Order : %+v,", order)
	}
	if order.OrderType != "limit" {
		t.Errorf("Worng OrderType. Order : %+v,", order)
	}
	if order.Quantity != "0.01" {
		t.Errorf("Worng Quantity. Order : %+v,", order)
	}
	if order.DiscQuantity != "0.0" {
		t.Errorf("Worng DiscQuantity. Order : %+v,", order)
	}
	if order.IcebergTotalQuantity != "0.0" {
		t.Errorf("Worng IcebergTotalQuantity. Order : %+v,", order)
	}
	if order.Side != "sell" {
		t.Errorf("Worng Side. Order : %+v,", order)
	}
	if order.FilledQuantity != "0.01" {
		t.Errorf("Worng FilledQuantity. Order : %+v,", order)
	}
	if order.Price != "500.0" {
		t.Errorf("Worng Price. Order : %+v,", order)
	}
	if order.CreatedAt != 1462123639 {
		t.Errorf("Worng CreatedAt. Order : %+v,", order)
	}
	if order.UpdatedAt != 1462123639 {
		t.Errorf("Worng UpdatedAt. Order : %+v,", order)
	}
	if order.Status != "filled" {
		t.Errorf("Worng Status. Order : %+v,", order)
	}
	if order.LeverageLevel != 2 {
		t.Errorf("Worng LeverageLevel. Order : %+v,", order)
	}
	if order.SourceExchange != "QUOINE" {
		t.Errorf("Worng SourceExchange. Order : %+v,", order)
	}
	if order.ProductID != 1 {
		t.Errorf("Worng ProductID. Order : %+v,", order)
	}
	if order.ProductCode != "CASH" {
		t.Errorf("Worng ProductCode. Order : %+v,", order)
	}
	if order.ProductCode != "CASH" {
		t.Errorf("Worng ProductCode. Order : %+v,", order)
	}
	if order.FundingCurrency != "USD" {
		t.Errorf("Worng FundingCurrency. Order : %+v,", order)
	}
	if order.CurrencyPairCode != "BTCUSD" {
		t.Errorf("Worng CurrencyPairCode. Order : %+v,", order)
	}
	if order.OrderFee != "0.0" {
		t.Errorf("Worng OrderFee. Order : %+v,", order)
	}
	if order.Executions[0].ID != 4566133 {
		t.Errorf("Worng Executions.ID. Order : %+v,", order)
	}
	if order.Executions[0].Quantity != "0.01" {
		t.Errorf("Worng Executions.Quantity. Order : %+v,", order)
	}
	if order.Executions[0].Price != "500.0" {
		t.Errorf("Worng Executions.Price. Order : %+v,", order)
	}
	if order.Executions[0].TakerSide != "buy" {
		t.Errorf("Worng Executions.TakerSide. Order : %+v,", order)
	}
	if order.Executions[0].MySide != "sell" {
		t.Errorf("Worng Executions.MySide. Order : %+v,", order)
	}
	if order.Executions[0].CreatedAt != 1465396785 {
		t.Errorf("Worng Executions.CreatedAt. Order : %+v,", order)
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
