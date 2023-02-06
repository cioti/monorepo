package transport

import (
	"github.com/cioti/monorepo/cms.api/shared"
	"github.com/cioti/monorepo/pkg/api"
	"github.com/cioti/monorepo/pkg/transport/http"
	"github.com/gin-gonic/gin"
)

type HttpTransportBuilder interface {
	Build() *gin.Engine
}

type httpTransportBuilder struct {
}

func NewHttpTransportBuilder() HttpTransportBuilder {
	return &httpTransportBuilder{}
}

func (b *httpTransportBuilder) Build() *gin.Engine {
	router := gin.Default()

	router.GET(shared.GetProjectsRoute, createGetProjectsEndpoint())
	return router
}

func createGetProjectsEndpoint() gin.HandlerFunc {
	endpoint := http.CreateEndpoint(
		http.EncodeJSONResponse,
		http.NopDecoder,
		func(ctx *gin.Context, request interface{}) (api.ApiResponse, api.ApiError) {
			return api.NewApiResponse(200, "success"), nil
		},
	)

	return func(ctx *gin.Context) {
		endpoint.Execute(ctx)
	}
}
