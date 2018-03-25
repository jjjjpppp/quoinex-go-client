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

func TestGetFiatAccounts(t *testing.T) {
	type Param struct {
		jsonResponse string
	}
	type Expect struct {
		path     string
		method   string
		accounts []*models.Account
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{jsonResponse: testutil.GetFiatAccountsJsonResponse()},
			expect: Expect{path: "/fiat_accounts", accounts: testutil.GetExpectedFiatAccountsModel()},
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
				if r.Method != "GET" {
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
		accounts, _ := client.GetFiatAccounts(ctx)
		if !cmp.Equal(accounts, c.expect.accounts) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(accounts, c.expect.accounts))
		}

	}
}
