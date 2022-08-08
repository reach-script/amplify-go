package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type NotFound struct {
	message  string
	original string
	stack    []Frame
}
type NotFoundError struct {
	Kind     string
	NotFound NotFound
}

func NewNotFoundError(err error, options ...notFoundOption) IError {
	callerSkip := 3

	for _, option := range options {
		switch option.kind() {
		case notFoundOptionKindPanic:
			callerSkip = 5
		}
	}

	return &NotFoundError{
		Kind: "not found",
		NotFound: NotFound{
			message:  err.Error(),
			original: fmt.Sprintf("%+v", err),
			stack:    callers(callerSkip),
		},
	}
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("not found error: %s", e.NotFound.message)
}
func (e NotFoundError) Response() ErrorResponse {
	return ErrorResponse{
		status: http.StatusNotFound,
		err:    e,
	}
}
func (e NotFoundError) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind     string  `json:"kind"`
		Message  string  `json:"message"`
		Original string  `json:"original"`
		Stack    []Frame `json:"stack"`
	}{
		Kind:     e.Kind,
		Message:  e.NotFound.message,
		Original: e.NotFound.original,
		Stack:    e.NotFound.stack,
	})
}

type notFoundOption interface {
	kind() notFoundOptionKind
}

type notFoundOptionKind = string

const notFoundOptionKindPanic = "panic"

type withNotFoundPanic struct{}

func (w withNotFoundPanic) Kind() notFoundOptionKind {
	return notFoundOptionKindPanic
}
