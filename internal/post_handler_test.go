package internal

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostHandlerPostsRequest(t *testing.T) {
	apiKey := "my key"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.EqualFold(r.Method, "post") {
			t.Error("expected request to be post method")
		}
		if r.Header.Get(apiKeyHeader) != apiKey {
			t.Errorf("expected apikey %s, got %s", apiKey, r.Header.Get(apiKeyHeader))
		}
		d := json.NewDecoder(r.Body)
		var actualReq *PostRequest
		if err := d.Decode(&actualReq); err != nil {
			t.Error(err)
		}

		w.WriteHeader(202)
	}))
	defer server.Close()

	req := &PostRequest{}
	if err := Post(server.URL, req, apiKey, false, true); err != nil {
		t.Error(err)
	}
}

func TestPostHandlerReturnsErrorOnFailure(t *testing.T) {
	apiKey := "my key"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(400)
	}))
	defer server.Close()

	req := &PostRequest{}
	if err := Post(server.URL, req, apiKey, false, false); err == nil {
		t.Error("expected error")
	}
}

func TestPostHandlerShouldNotSendRequestInSilentMode(t *testing.T) {
	apiKey := "my key"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		t.Error("this should not be called")
	}))
	defer server.Close()

	req := &PostRequest{}
	if err := Post(server.URL, req, apiKey, true, false); err != nil {
		t.Error(err)
	}
}
