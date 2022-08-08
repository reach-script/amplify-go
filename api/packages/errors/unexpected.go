package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Unexpected struct {
	message  string
	original string
	stack    []Frame
}
type UnexpectedError struct {
	Kind       string
	Unexpected Unexpected
}

func NewUnexpectedError(err error, options ...unexpectedOption) IError {
	callerSkip := 3

	for _, option := range options {
		switch option.kind() {
		case unexpectedOptionKindPanic:
			callerSkip = 5
		}
	}

	return &UnexpectedError{
		Kind: "unexpected",
		Unexpected: Unexpected{
			message:  err.Error(),
			original: fmt.Sprintf("%+v", err),
			stack:    callers(callerSkip),
		},
	}
}

func (e UnexpectedError) Error() string {
	return fmt.Sprintf("unexpected error: %s", e.Unexpected.message)
}
func (e UnexpectedError) Response() ErrorResponse {
	return ErrorResponse{
		status: http.StatusInternalServerError,
		err:    e,
	}
}
func (e UnexpectedError) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind     string  `json:"kind"`
		Message  string  `json:"message"`
		Original string  `json:"original"`
		Stack    []Frame `json:"stack"`
	}{
		Kind:     e.Kind,
		Message:  e.Unexpected.message,
		Original: e.Unexpected.original,
		Stack:    e.Unexpected.stack,
	})
}

type unexpectedOption interface {
	kind() unexpectedOptionKind
}

type unexpectedOptionKind = string

const unexpectedOptionKindPanic = "panic"

type withUnexpectedPanic struct{}

func (w withUnexpectedPanic) Kind() unexpectedOptionKind {
	return unexpectedOptionKindPanic
}
