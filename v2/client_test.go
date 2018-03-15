package quoinex

import (
	"context"
	"testing"
	"time"
)

func Test_NewClient_ApiTokenError(t *testing.T) {
	if _, e := NewClient("", "secret", nil); e == nil {
		t.Error("err should be returned")
	}
}

func Test_NewClient_ApiSecretError(t *testing.T) {
	if _, e := NewClient("apiToken", "", nil); e == nil {
		t.Error("err should be returned")
	}
}

func Test_NewClient_Success(t *testing.T) {
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

func Test_NewRequest_Success_No_QueryParam(t *testing.T) {
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

func Test_NewRequest_Success_With_QueryParam(t *testing.T) {
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
