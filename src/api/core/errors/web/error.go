package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Error struct {
	Status  int
	Code    string
	Message string
}

func (e *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{
		Code:    e.Code,
		Message: e.Message,
	})
}

func (e *Error) StatusCode() int {
	return e.Status
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func NewError(statusCode int, message string) error {
	return NewError(statusCode, message)
}

func NewErrorf(status int, format string, args ...interface{}) error {
	return &Error{
		Code:    strings.ReplaceAll(strings.ToLower(http.StatusText(status)), " ", " - "),
		Message: fmt.Sprintf(format, args...),
		Status:  status,
	}
}
