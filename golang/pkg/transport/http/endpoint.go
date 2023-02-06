package http

import (
	"github.com/cioti/monorepo/pkg/api"
	"github.com/gin-gonic/gin"
)

type Endpoint interface {
	Execute(ctx *gin.Context)
}

type endpoint struct {
	encoder ResponseEncoder
	decoder RequestDecoder
	handler EndpointHandler
}

func (e *endpoint) Execute(ctx *gin.Context) {
	request, err := e.decoder(ctx)
	if err != nil {
		e.encoder(ctx, err.StatusCode(), err.Err())
		return
	}

	if response, err := e.handler(ctx, request); err != nil {
		e.encoder(ctx, err.StatusCode(), err.Err())
	} else {
		e.encoder(ctx, response.StatusCode(), response.Payload())
	}
}

func CreateEndpoint(e ResponseEncoder, d RequestDecoder, h EndpointHandler) Endpoint {
	return &endpoint{
		encoder: e,
		decoder: d,
		handler: h,
	}
}

type EndpointHandler = func(ctx *gin.Context, request interface{}) (api.ApiResponse, api.ApiError)
