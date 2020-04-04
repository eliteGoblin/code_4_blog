package errors

import (
	"fmt"
	"net/http"
)

type BadRequest struct {
	message string
	code    string
}

func (br BadRequest) Error() string {
	return br.message
}

func (br BadRequest) Code() string {
	return br.code
}

func (br BadRequest) HTTPStatus() int {
	return http.StatusBadRequest
}

func (br BadRequest) IsServiceFailure() bool {
	return false
}

func NewBadRequest(code, msg string, args ...interface{}) *BadRequest {
	return &BadRequest{
		code:    code,
		message: fmt.Sprintf(msg, args...),
	}
}

type NotFound struct {
	message string
}

func NewNotFound(msg string, args ...interface{}) *NotFound {
	return &NotFound{
		message: fmt.Sprintf(msg, args...),
	}
}

func (nf NotFound) Error() string {
	return nf.message
}

func (nf NotFound) HTTPStatus() int {
	return http.StatusNotFound
}

func (nf NotFound) IsServiceFailure() bool {
	return false
}

type Forbidden struct {
	message string
	code    string
}

func NewForbidden(code, msg string, args ...interface{}) *Forbidden {
	return &Forbidden{
		message: fmt.Sprintf(msg, args...),
		code:    code,
	}
}

func (fb Forbidden) Code() string {
	return fb.code
}

func (fb Forbidden) Error() string {
	return fb.message
}

func (fb Forbidden) HTTPStatus() int {
	return http.StatusForbidden
}

func (fb Forbidden) IsServiceFailure() bool {
	return false
}
