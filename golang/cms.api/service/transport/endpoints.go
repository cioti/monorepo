package transport

import (
	"context"
	"net/http"

	"github.com/cioti/monorepo/cms.api/service/app"
	"github.com/cioti/monorepo/cms.api/shared"
	"github.com/cioti/monorepo/pkg/api"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints interface {
	CreateProject() endpoint.Endpoint
	AddModel() endpoint.Endpoint
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
		cd, ok := request.(shared.CreateProjectCommand)
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

func (e endpoints) AddModel() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		cd, err := cast[shared.AddModelCommand](request)
		if err != nil {
			return nil, err
		}
		err = e.projectSvc.AddModel(ctx, cd)
		if err != nil {
			return nil, err
		}

		return api.NewApiResponse(http.StatusOK, nil), nil
	}
}

func cast[T any](req interface{}) (result T, err error) {
	result, ok := req.(T)
	if !ok {
		return result, api.NewBadRequestErrorf("unable to cast endpoint request data")
	}

	return result, nil
}
