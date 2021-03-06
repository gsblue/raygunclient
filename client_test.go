package raygunclient

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"reflect"

	"github.com/gsblue/raygunclient/internal"
	"github.com/gsblue/raygunclient/stack"
)

func TestClientCanNotifyRaygunAboutError(t *testing.T) {
	err := errors.New("something went wrong")
	entry := NewErrorEntry(err)
	entry.SetUser("user").SetTags([]string{"1", "2"}).SetCustomData("extra info")
	hr, _ := http.NewRequest("GET", "/test", nil)

	entry.SetRequest(hr)
	c := NewClient("apiKey", "version", defaultOptions).(*client)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		d := json.NewDecoder(r.Body)
		var req *internal.PostRequest
		if err := d.Decode(&req); err != nil {
			t.Error(err)
		}

		if req.Details.Error.Message != err.Error() {
			t.Errorf("expected error %v, but got %s", err, req.Details.Error.Message)
		} else if req.Details.Version != c.version {
			t.Errorf("expected version %s, but got %s", c.version, req.Details.Version)
		} else if req.Details.Context.Identifier != c.ctxIdentifier {
			t.Errorf("expected identifier %s, but got %s", c.ctxIdentifier, req.Details.Context.Identifier)
		} else if req.Details.User.Identifier != entry.User {
			t.Errorf("expected user %s, but got %s", entry.User, req.Details.User.Identifier)
		} else if !reflect.DeepEqual(req.Details.Tags, entry.Tags) {
			t.Errorf("expected tags %s, but got %s", entry.Tags, req.Details.Tags)
		} else if !reflect.DeepEqual(req.Details.UserCustomData, entry.CustomData) {
			t.Errorf("expected custom data %s, but got %s", entry.CustomData, req.Details.UserCustomData)
		} else if req.Details.Request.URL != hr.URL.String() {
			t.Errorf("expected request url %s, but got %s", hr.URL.String(), req.Details.Request.URL)
		} else if req.Details.Request.HTTPMethod != hr.Method {
			t.Errorf("expected request method %s, but got %s", hr.Method, req.Details.Request.HTTPMethod)
		} else if len(req.Details.Error.StackTrace) == 0 {
			t.Error("expected stack trace to be sent")
		}

		w.WriteHeader(202)
	}))
	defer server.Close()

	c.endpoint = server.URL

	if err := c.Notify(entry); err != nil {
		t.Error(err)
	}
}

func TestClientCanNotifyRaygunAboutErrorWithCustomStackTrace(t *testing.T) {
	err := errors.New("something went wrong")
	entry := NewErrorEntry(err)
	trace := make(stack.Trace, 0, 0)
	trace.AddEntry(23, "github.com/gsblue/raygunclient", "main.go", "Notify")

	c := NewClient("apiKey", "version", defaultOptions).(*client)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		d := json.NewDecoder(r.Body)
		var req *internal.PostRequest
		if err := d.Decode(&req); err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(req.Details.Error.StackTrace, trace) {
			t.Errorf("expected stack trace %v, got %v", trace, req.Details.Error.StackTrace)
		}

		w.WriteHeader(202)
	}))
	defer server.Close()

	c.endpoint = server.URL

	if err := c.NotifyWithStackTrace(entry, trace); err != nil {
		t.Error(err)
	}
}
