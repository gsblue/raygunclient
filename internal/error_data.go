package internal

import "github.com/gsblue/raygunclient/stack"

type errorData struct {
	Message    string      `json:"message"`    // the actual message the error produced
	StackTrace stack.Trace `json:"stackTrace"` // the error's stack trace
}
