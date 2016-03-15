package raygunclient

import (
	"net/http"

	"github.com/gsblue/raygunclient/httpdata"
)

//ErrorEntry holds the information about the error and meta data
//needed to be recorded in raygun
type ErrorEntry struct {
	Error      error
	Request    *httpdata.HTTPRequest
	CustomData interface{}
	User       string
	Tags       []string
}

//NewErrorEntry creates a new error entry
func NewErrorEntry(err error) *ErrorEntry {
	return &ErrorEntry{
		Error: err,
	}
}

// SetRequest is a chainable option-setting method to add a request to this entry.
func (e *ErrorEntry) SetRequest(r *http.Request) *ErrorEntry {
	e.Request = httpdata.NewHTTPRequest(r)
	return e
}

// SetCustomData is a chainable option-setting method to add arbitrary custom data
// to the entry. Note that the given type (or at least parts of it)
// must implement the Marshaler-interface for this to work.
func (e *ErrorEntry) SetCustomData(data interface{}) *ErrorEntry {
	e.CustomData = data
	return e
}

// SetUser is a chainable option-setting method to add an affected Username to this
// entry.
func (e *ErrorEntry) SetUser(u string) *ErrorEntry {
	e.User = u
	return e
}

// SetTags is a chainable option-setting method to add tags to this entry.
func (e *ErrorEntry) SetTags(tags []string) *ErrorEntry {
	e.Tags = tags
	return e
}
