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

type EndpointInfo struct {
	Method   string
	Route    string
	Endpoint endpoint.Endpoint
	Decoder  kithttp.DecodeRequestFunc
}

type EndpointBuilder struct {
	router *mux.Router
	logger logging.Logger
}

func NewHEndpointBuilder() *EndpointBuilder {
	return &EndpointBuilder{
		router: mux.NewRouter(),
	}
}

func (b *EndpointBuilder) AddEndpoint(info EndpointInfo) {
	errorLogger := kithttp.ServerErrorLogger(b.logger)
	errorEncoder := kithttp.ServerErrorEncoder(JsonErrorEncoder)
	handler := kithttp.NewServer(info.Endpoint, info.Decoder, kithttp.EncodeJSONResponse, errorLogger, errorEncoder)
	b.router.Handle(info.Route, handler).Methods(info.Method)
}

func (b *EndpointBuilder) Build() http.Handler {
	return handlers.LoggingHandler(os.Stdout, b.router)
}
