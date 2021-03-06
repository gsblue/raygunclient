package stack

import (
	"fmt"

	"github.com/kaeuferportal/stack2struct"
)

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

// Format satisfies the Formatter interface from fmt
func (s *Trace) Format(state fmt.State, verb rune) {
	if s != nil {
		for _, el := range *s {
			fmt.Fprintf(state, "%s:%s in %s line %d", el.PackageName, el.MethodName, el.FileName, el.LineNumber)
		}
	}
}
