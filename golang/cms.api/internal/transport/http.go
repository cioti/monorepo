package transport

import (
	"context"
	"net/http"

	"github.com/cioti/monorepo/cms.api/shared"
	pkghttp "github.com/cioti/monorepo/pkg/transport/http"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
)

func CreateHandler() http.Handler {
	builder := pkghttp.NewHttpRouteBuilder()
	builder.AddRoute(http.MethodGet, shared.GetProjectsRoute, createGetProjectsEndpoint(), createDecodeGetProjects())

	return builder.Build()
}

func createGetProjectsEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return "lalallala", nil
	}
}

func createDecodeGetProjects() kithttp.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (request interface{}, err error) {
		return "test", nil
	}
}
