package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cioti/monorepo/cms.api/service/transport"
	"github.com/cioti/monorepo/cms.api/shared"
	"github.com/cioti/monorepo/pkg/api"
	"github.com/cioti/monorepo/pkg/logging"
	pkghttp "github.com/cioti/monorepo/pkg/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func CreateHandler(endpoints transport.Endpoints, logger logging.Logger) http.Handler {
	builder := pkghttp.NewHEndpointBuilder(logger)
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

	builder.AddEndpoint(pkghttp.EndpointInfo{
		Method:   http.MethodPost,
		Route:    shared.AddModelRoute,
		Decoder:  decodeAddModelRequest,
		Endpoint: endpoints.AddModel(),
	})

	return builder.Build()
}

func decodeGetProjectsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return "test", nil
}

func decodeCreateProjectRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var cmd shared.CreateProjectCommand
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&cmd)
	if err != nil {
		return nil, api.NewBadRequestError(err, "unable to decode 'CreateProject' request data")
	}

	return cmd, nil
}

func decodeAddModelRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var cmd shared.AddModelCommand
	vars := mux.Vars(r)
	id := vars[shared.ProjectIdRouteParam]
	if len(id) == 0 {
		return nil, api.NewBadRequestErrorf("project id route parameter is required")
	}

	projectID, err := uuid.Parse(id)
	fmt.Println(projectID.String())
	if err != nil {
		return nil, api.NewBadRequestError(err, "project id is not a valid UUID")
	}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&cmd)
	if err != nil {
		return nil, api.NewBadRequestError(err, "unable to decode 'AddModel' request data")
	}

	cmd.ProjectID = projectID

	return cmd, nil
}
