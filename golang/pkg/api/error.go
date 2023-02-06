package api

import "net/http"

type ApiError interface {
	StatusCode() int
	Err() string
}

type apiError struct {
	statusCode int
	err        error
}

func (e *apiError) StatusCode() int {
	return e.statusCode
}

func (e *apiError) Err() string {
	return e.err.Error()
}

func NewBadRequestError(err error) ApiError {
	return &apiError{
		statusCode: http.StatusBadRequest,
		err:        err,
	}
}

func NewNotFoundError(err error) ApiError {
	return &apiError{
		statusCode: http.StatusNotFound,
		err:        err,
	}
}
