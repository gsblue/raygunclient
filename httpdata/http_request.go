package httpdata

import (
	"fmt"
	"net/http"
	"strings"
)

// HTTPRequest holds all information on the request from the context
type HTTPRequest struct {
	HostName    string            `json:"hostName"`
	URL         string            `json:"url"`
	HTTPMethod  string            `json:"httpMethod"`
	IPAddress   string            `json:"ipAddress"`
	QueryString map[string]string `json:"queryString"`
	Form        map[string]string `json:"form"`
	Headers     map[string]string `json:"headers"`
}

// NewHTTPRequest parses all information from the request in the context to a
// struct. The struct is empty if no request was set.
func NewHTTPRequest(r *http.Request) *HTTPRequest {
	if r == nil {
		return &HTTPRequest{}
	}

	r.ParseForm()

	return &HTTPRequest{
		HostName:    r.Host,
		URL:         r.URL.String(),
		HTTPMethod:  r.Method,
		IPAddress:   r.RemoteAddr,
		QueryString: arrayMapToStringMap(r.URL.Query()),
		Form:        arrayMapToStringMap(r.PostForm),
		Headers:     arrayMapToStringMap(r.Header),
	}
}

func arrayMapToStringMap(arrayMap map[string][]string) map[string]string {
	entries := make(map[string]string)
	for k, v := range arrayMap {
		if len(v) > 1 {
			entries[k] = fmt.Sprintf("[%s]", strings.Join(v, "; "))
		} else {
			entries[k] = v[0]
		}
	}
	return entries
}
