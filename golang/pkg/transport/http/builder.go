package http

import (
	"net/http"
	"os"

	"github.com/cioti/monorepo/pkg/logging"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type HttpRouteBuilder struct {
	router *mux.Router
	logger logging.Logger
}

func NewHttpRouteBuilder() *HttpRouteBuilder {
	return &HttpRouteBuilder{
		router: mux.NewRouter(),
	}
}

func (b *HttpRouteBuilder) AddRoute(
	method string,
	route string,
	endpoint endpoint.Endpoint,
	decoder kithttp.DecodeRequestFunc) {

	errorLogger := kithttp.ServerErrorLogger(b.logger)
	errorEncoder := kithttp.ServerErrorEncoder(JsonErrorEncoder)
	handler := kithttp.NewServer(endpoint, decoder, kithttp.EncodeJSONResponse, errorLogger, errorEncoder)
	b.router.Handle(route, handler).Methods(method)
}

func (b *HttpRouteBuilder) Build() http.Handler {
	return handlers.LoggingHandler(os.Stdout, b.router)
}
