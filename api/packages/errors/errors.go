package errors

import (
	"github.com/gin-gonic/gin"
)

type IError interface {
	Error() string
	Response() ErrorResponse
	MarshalJSON() ([]byte, error)
}

type ErrorResponse struct {
	status int
	err    IError
}

func (res ErrorResponse) Do(c *gin.Context, requestID string) {
	c.PureJSON(res.status, struct {
		RequestID string `json:"request_id"`
		Error     IError `json:"error"`
	}{
		RequestID: requestID,
		Error:     res.err,
	})
}
