package internal

import (
	"os"
	"time"

	"github.com/gsblue/raygunclient/httpdata"
	"github.com/gsblue/raygunclient/stack"
)

var defaultClientInfo = &clientInfo{
	"raygunclient",
	"0.1",
	"https://github.com/gsblue/raygunclient",
}

//PostRequest is the contract for posting errors to raygun
type PostRequest struct {
	OccuredOn string        `json:"occurredOn"` // the time the error occured on, format 2006-01-02T15:04:05Z
	Details   *errorDetails `json:"details"`    // all the details needed by the API
}

type errorDetails struct {
	MachineName    string                `json:"machineName"`    // the machine's hostname
	Version        string                `json:"version"`        // the version from context
	Error          *errorData            `json:"error"`          // everything we know about the error itself
	Tags           []string              `json:"tags"`           // the tags from context
	UserCustomData interface{}           `json:"userCustomData"` // the custom data from the context
	Request        *httpdata.HTTPRequest `json:"request"`        // the request from the context
	User           *user                 `json:"user"`           // the user from the context
	Context        *context              `json:"context"`        // the identifier from the context
	Client         *clientInfo           `json:"client"`         // information on this client
}

// user holds information on the affected user.
type user struct {
	Identifier string `json:"identifier"`
}

// context holds information on the program context.
type context struct {
	Identifier string `json:"identifier"`
}

// clientInfo is the struct holding information on this client.
type clientInfo struct {
	Name      string `json:"identifier"`
	Version   string `json:"version"`
	ClientURL string `json:"clientUrl"`
}

type errorData struct {
	Message    string      `json:"message"`    // the actual message the error produced
	StackTrace stack.Trace `json:"stackTrace"` // the error's stack trace
}

//NewPostRequest creates a post request frome the arguments provided
func NewPostRequest(err error, req *httpdata.HTTPRequest,
	customData interface{}, u string, tags []string,
	version string, id string) *PostRequest {

	return &PostRequest{
		OccuredOn: time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		Details:   newErrorDetails(err, req, customData, u, tags, version, id),
	}
}

func newErrorDetails(err error, req *httpdata.HTTPRequest,
	customData interface{}, u string, tags []string,
	version string, id string) *errorDetails {

	hostname, e := os.Hostname()
	if e != nil {
		hostname = "not available"
	}
	return &errorDetails{
		MachineName:    hostname,
		Version:        version,
		Error:          &errorData{err.Error(), stack.CurrentStack()},
		Tags:           tags,
		User:           &user{u},
		Request:        req,
		UserCustomData: customData,
		Context:        &context{id},
		Client:         defaultClientInfo,
	}
}

//SetStackTrace sets the stack trace information in the request
func (req *PostRequest) SetStackTrace(t stack.Trace) {
	req.Details.Error.StackTrace = t
}
