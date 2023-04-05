package api

import (
	"fmt"
	"net/http"

	"github.com/cioti/monorepo/pkg/errors"
)

type ApiError interface {
	StatusCode() int
	ErrorCode() string
	Error() string
}

type apiError struct {
	statusCode int
	errorCode  string
	err        error
}

func (e *apiError) StatusCode() int {
	if e.statusCode < 1 || e.statusCode > 500 {
		return http.StatusInternalServerError
	}

	return e.statusCode
}

func (e *apiError) Error() string {
	return e.err.Error()
}

func (e *apiError) ErrorCode() string {
	return e.errorCode
}

func NewBadRequestErrorf(format string, args ...interface{}) ApiError {
	return &apiError{
		statusCode: http.StatusBadRequest,
		err:        errors.New(fmt.Sprintf(format, args...)),
	}
}

func NewBadRequestError(err error, format string, args ...interface{}) ApiError {
	return &apiError{
		statusCode: http.StatusBadRequest,
		err:        errors.Wrapf(err, format),
	}
}

func NewNotFoundErrorf(format string, args ...interface{}) ApiError {
	return &apiError{
		statusCode: http.StatusNotFound,
		err:        errors.New(fmt.Sprintf(format, args...)),
	}
}

func NewNotFoundError(err error) ApiError {
	return &apiError{
		statusCode: http.StatusNotFound,
		err:        err,
	}
}
