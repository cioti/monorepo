package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/cioti/monorepo/cms.api/service/transport"
	"github.com/cioti/monorepo/cms.api/shared"
	"github.com/cioti/monorepo/pkg/api"
	pkghttp "github.com/cioti/monorepo/pkg/transport/http"
)

func CreateHandler(endpoints transport.Endpoints) http.Handler {
	builder := pkghttp.NewHEndpointBuilder()
	builder.AddEndpoint(pkghttp.EndpointInfo{
		Method:   http.MethodGet,
		Route:    shared.GetProjectsRoute,
		Decoder:  decodeGetProjectsRequest,
		Endpoint: endpoints.GetProjects(),
	})

	builder.AddEndpoint(pkghttp.EndpointInfo{
		Method:   http.MethodPost,
		Route:    shared.CreateProjectRoute,
		Decoder:  decodeCreateProjectRequest,
		Endpoint: endpoints.CreateProject(),
	})

	return builder.Build()
}

func decodeGetProjectsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return "test", nil
}

func decodeCreateProjectRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var cd shared.CreateProjectCD
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&cd)
	if err != nil {
		return nil, api.NewBadRequestError(err, "unable to decode 'createProject' request data")
	}

	return cd, nil
}
