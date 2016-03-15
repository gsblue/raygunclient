# Raygun Client [![Build Status](https://travis-ci.org/gsblue/raygunclient.svg)](https://travis-ci.org/gsblue/raygunclient) [![Coverage Status](https://coveralls.io/repos/github/gsblue/raygunclient/badge.svg?branch=master)](https://coveralls.io/github/gsblue/raygunclient?branch=master) [![GoDoc](https://godoc.org/github.com/gsblue/raygunclient?status.svg)](https://godoc.org/github.com/gsblue/raygunclient)

A Raygun.io golang client for handling errors.

##Usage

```go
import (
	"github.com/gsblue/raygunclient"
    "errors"
)
var n raygunclient.Notifier

func init() {
	n := raygunclient.NewClient("your api key", "application version no", nil)
}

func SomeFunctionWhichNeedsToHandleError() {

	if someErr := doSomeWork(); someErr != nil {
        entry := NewErrorEntry(someErr)
        entry.SetUser("user identifier").
            SetTags([]string{"tag 1", "tag 2"}).
            SetCustomData(&struct{ OrderNo int }{340})
        
        if err := n.Notify(entry); err != nil {
            panic(err)
        }
    }
}

func doSomeWork() error {
    return errors.New("some error")
}

```
If you are capturing an error in context of a http request, you can send the request data too

```go
import (
	"github.com/gsblue/raygunclient"
    "errors"
)
var n raygunclient.Notifier

func init() {
	n := raygunclient.NewClient("your api key", "application version no", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
    
    if someErr := doSomeWork(); someErr != nil {
        entry := NewErrorEntry(someErr)
        entry.SetRequest(r)
        
        if err := n.Notify(entry); err != nil {
            panic(err)
        }
    }
}

func doSomeWork() error {
    return errors.New("some error")
}

```
Note, by default stack trace is captured and sent to raygun. 
But if you want to use send custom stack trace, you can use the ```NotifyWithStackTrace``` method.


Pull requests are welcome.