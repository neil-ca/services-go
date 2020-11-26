package handler

import (
	"github.com/go-chi/render"
	"net/http"
)

type ErrorResponse struct {
	Err        error  `json:"-"`
	StatusCode int    `json:"-"`
	StatusText string `json:"status_text"`
	Message    string `json:"message"`
}

var (
	ErrMethodNotAllowed = &ErrorResponse{StatusCode: 403, Message: "method not allowed"}
	ErrNotFound         = &ErrorResponse{StatusCode: 404, Message: "Resource not found"}
	ErrBadRequest       = &ErrorResponse{StatusCode: 400, Message: "Bad Request"}
)

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

func ErrorRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:        err,
		StatusCode: 400,
		StatusText: "Bad Request",
		Message:    err.Error(),
	}
}

func ServerErrorRendered(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:        err,
		StatusCode: 500,
		StatusText: "Internal Server Error",
		Message:    err.Error(),
	}
}
