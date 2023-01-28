package errors

import (
	"net/http"
	"strings"

	"github.com/engmanuelricardoperez/fast-truck-api/src/api/core/errors/web"
)

const (
	BadRequestMessage          = "Invalid request parameters."
	ResourceNotFountMessage    = "Resource not found"
	InternalServerErrorMessage = "Internal Server Error"
)

func GetCmmonAPIErrors(err error) error {
	switch err.(type) {
	case BadRequestError:
		return NewBadRequest(err.Error())
	default:
		return NewInternalServer(err.Error())
	}
}

func NewBadRequest(messages ...string) error {
	message := BadRequestMessage
	if len(message) > 0 {
		message = strings.Join(messages, " - ")
	}
	return web.NewError(http.StatusBadRequest, message)
}

func NewInternalServer(messages ...string) error {
	message := InternalServerErrorMessage
	if len(message) > 0 {
		message = strings.Join(messages, " - ")
	}
	return web.NewError(http.StatusInternalServerError, message)
}
