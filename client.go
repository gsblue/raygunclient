package raygunclient

import (
	"github.com/gsblue/raygunclient/internal"
	"github.com/gsblue/raygunclient/stack"
	"github.com/pborman/uuid"
)

const raygunEndpoint = "https://api.raygun.io"

var defaultOptions = &ClientOptions{false, false}

//Notifier notifies raygun about the error
type Notifier interface {
	Notify(entry *ErrorEntry) error
	NotifyWithStackTrace(entry *ErrorEntry, t stack.Trace) error
}

type client struct {
	endpoint      string
	apiKey        string
	opts          *ClientOptions
	ctxIdentifier string
	version       string
}

//ClientOptions provides the options to configure the behavior of the client
type ClientOptions struct {
	Silent bool
	Debug  bool
}

//NewClient creates a new client to send errors to raygun
func NewClient(apiKey, version string, opts *ClientOptions) Notifier {
	if opts == nil {
		opts = defaultOptions
	}

	return &client{
		endpoint:      raygunEndpoint,
		apiKey:        apiKey,
		version:       version,
		opts:          opts,
		ctxIdentifier: uuid.NewUUID().String(),
	}
}

func (c *client) Notify(entry *ErrorEntry) error {
	req := internal.NewPostRequest(entry.Error, entry.Request, entry.CustomData,
		entry.User, entry.Tags, c.version, c.ctxIdentifier)

	return internal.Post(c.endpoint, req, c.apiKey, c.opts.Silent, c.opts.Debug)
}

func (c *client) NotifyWithStackTrace(entry *ErrorEntry, t stack.Trace) error {
	req := internal.NewPostRequest(entry.Error, entry.Request, entry.CustomData,
		entry.User, entry.Tags, c.version, c.ctxIdentifier)
	req.SetStackTrace(t)

	return internal.Post(c.endpoint, req, c.apiKey, c.opts.Silent, c.opts.Debug)
}
