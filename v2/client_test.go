package quoinex

import (
	"testing"
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
