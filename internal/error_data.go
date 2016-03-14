package internal

import "github.com/kaeuferportal/stack2struct"

type errorData struct {
	Message    string     `json:"message"`    // the actual message the error produced
	StackTrace stackTrace `json:"stackTrace"` // the error's stack trace
}

// currentStac returns the current stack.
func currentStack() stackTrace {
	s := make(stackTrace, 0, 0)
	stack2struct.Current(&s)
	return s
}

// stackTraceElement is one element of the error's stack trace. It is filled by
// stack2struct.
type stackTraceElement struct {
	LineNumber  int    `json:"lineNumber"`
	PackageName string `json:"className"`
	FileName    string `json:"fileName"`
	MethodName  string `json:"methodName"`
}

// stackTrace is the stack the trace will be parsed into.
type stackTrace []*stackTraceElement

// AddEntry is the method used by stack2struct to dump parsed elements.
func (s *stackTrace) AddEntry(lineNumber int, packageName, fileName, methodName string) {
	*s = append(*s, &stackTraceElement{lineNumber, packageName, fileName, methodName})
}
