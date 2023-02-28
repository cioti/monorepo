package http

import (
	"context"
	"encoding/json"
	"net/http"
)

func JsonErrorEncoder(_ context.Context, err error, responseWriter http.ResponseWriter) {
	// if err == nil {
	// 	err = tleerrors.NewInternalServerErrorf("Unexpected")
	// }
	WriteErrorResponseHeaders(responseWriter, err)
	responseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(responseWriter).Encode(map[string]interface{}{"error": err.Error()})
}

// WriteErrorResponseHeaders writes relevant headers to json response based on a given error.
func WriteErrorResponseHeaders(responseWriter http.ResponseWriter, err error) {

}
