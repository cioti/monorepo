package transport

import (
	"context"
	"errors"
	"net/http"

	"github.com/cioti/monorepo/cms.api/service/app"
	"github.com/cioti/monorepo/cms.api/shared"
	"github.com/cioti/monorepo/pkg/api"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints interface {
	CreateProject() endpoint.Endpoint
	GetProjects() endpoint.Endpoint
}

type endpoints struct {
	projectSvc app.ProjectService
}

func NewEndpoints(projectSvc app.ProjectService) Endpoints {
	return &endpoints{
		projectSvc: projectSvc,
	}
}

func (e endpoints) GetProjects() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		data, err := e.projectSvc.GetProjects()
		if err != nil {
			return nil, err
		}

		return api.NewApiResponse(http.StatusOK, data), nil
	}
}

func (e endpoints) CreateProject() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		cd, ok := request.(shared.CreateProjectCD)
		if !ok {
			return nil, api.NewBadRequestErrorf("unable to cast 'CreateProject' request data to 'CreateProjectCD'")
		}
		err := e.projectSvc.CreateProject(ctx, cd)
		if err != nil {
			return nil, err
		}

		return api.NewApiResponse(http.StatusOK, nil), nil
	}
}

func cast[T any](req interface{}) (result T, err error) {
	result, ok := req.(T)
	if !ok {
		return result, errors.New("test")
	}

	return result, nil
}
