package stack

import "github.com/kaeuferportal/stack2struct"

// CurrentStack returns the current stack.
func CurrentStack() Trace {
	s := make(Trace, 0, 0)
	stack2struct.Current(&s)
	return s
}

// TraceElement is one element of the error's stack trace. It is filled by
// stack2struct.
type TraceElement struct {
	LineNumber  int    `json:"lineNumber"`
	PackageName string `json:"className"`
	FileName    string `json:"fileName"`
	MethodName  string `json:"methodName"`
}

// Trace is the stack the trace will be parsed into.
type Trace []*TraceElement

// AddEntry is the method used by stack2struct to dump parsed elements.
func (s *Trace) AddEntry(lineNumber int, packageName, fileName, methodName string) {
	*s = append(*s, &TraceElement{lineNumber, packageName, fileName, methodName})
}
