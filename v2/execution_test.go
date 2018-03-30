package quoinex

import (
	"context"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/jjjjpppp/quoinex-go-client/v2/models"
	"github.com/jjjjpppp/quoinex-go-client/v2/testutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

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
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
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
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		executions, err := client.GetExecutionsByTimestamp(ctx, c.param.productID, c.param.limit, c.param.timestamp)
		if err != nil {
			t.Errorf("Error. %+v", err)
		}
		if !cmp.Equal(executions, c.expect.executions) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(executions, c.expect.executions))
		}
	}
}

func TestGetOwnExecutions(t *testing.T) {
	type Param struct {
		productID    int
		jsonResponse string
	}
	type Expect struct {
		path       string
		method     string
		executions *models.Executions
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{productID: 1001232, jsonResponse: testutil.GetOwnExecutionsJsonResponse()},
			expect: Expect{path: "/executions/me?product_id=1001232", method: "GET", executions: testutil.GetExpectedOwnExecutionsModel()},
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
				if r.Method != c.expect.method {
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
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		executions, _ := client.GetOwnExecutions(ctx, c.param.productID)
		if !cmp.Equal(executions, c.expect.executions) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(executions, c.expect.executions))

		}

	}
}
