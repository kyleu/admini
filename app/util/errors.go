package util

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

type unwrappable interface {
	Unwrap() error
}

// Stack frame definition
type ErrorFrame struct {
	Key string
	Loc string
}

// An error's message, stack, and cause
type ErrorDetail struct {
	Message    string
	StackTrace errors.StackTrace
	Cause      *ErrorDetail
}

// Creates an ErrorDetail for the provided error
func GetErrorDetail(e error) *ErrorDetail {
	var stack errors.StackTrace

	t, ok := e.(stackTracer)
	if ok {
		stack = t.StackTrace()
	}

	var cause *ErrorDetail

	u, ok := e.(unwrappable)
	if ok {
		cause = GetErrorDetail(u.Unwrap())
	}

	return &ErrorDetail{
		Message:    e.Error(),
		StackTrace: stack,
		Cause:      cause,
	}
}

// Converts a stack trace to a set of ErrorFrames
func TraceDetail(trace errors.StackTrace) []ErrorFrame {
	s := fmt.Sprintf("%+v", trace)
	lines := strings.Split(s, "\n")
	validLines := []string{}

	for _, line := range lines {
		l := strings.TrimSpace(line)
		if len(l) > 0 {
			validLines = append(validLines, l)
		}
	}

	var ret []ErrorFrame
	for i := 0; i < len(validLines)-1; i += 2 {
		f := ErrorFrame{Key: validLines[i], Loc: validLines[i+1]}
		ret = append(ret, f)
	}
	return ret
}
