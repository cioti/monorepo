package http

import (
	"context"
	"encoding/json"
	"net/http"

	pkgapi "github.com/cioti/monorepo/pkg/api"
)

const (
	jsonContentType = "application/json; charset=utf-8"
	unknownErrMsg   = "An unknown error occured on the server"
	unknownErrCode  = "500-Internal"
)

func JsonErrorEncoder(_ context.Context, err error, responseWriter http.ResponseWriter) {
	responseWriter.Header().Set("Content-Type", jsonContentType)
	apiErr, ok := err.(pkgapi.ApiError)
	if !ok {
		errorResponse := pkgapi.NewApiErrorResponse(http.StatusInternalServerError, unknownErrMsg, unknownErrCode)
		responseWriter.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(responseWriter).Encode(errorResponse)

		return
	}

	errorResponse := pkgapi.NewApiErrorResponseFromError(apiErr)
	responseWriter.WriteHeader(apiErr.StatusCode())
	json.NewEncoder(responseWriter).Encode(errorResponse)
}
