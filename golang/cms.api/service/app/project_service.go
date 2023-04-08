package app

import (
	"context"

	"github.com/cioti/monorepo/cms.api/service/domain"
	"github.com/cioti/monorepo/cms.api/shared"
)

type ProjectService interface {
	CreateProject(ctx context.Context, cd shared.CreateProjectCommand) error
	AddModel(ctx context.Context, cd shared.AddModelCommand) error
	GetProjects() ([]shared.ProjectDTO, error)
}

type projectService struct {
	repo domain.ProjectRepository
}

func NewProjectService(repo domain.ProjectRepository) ProjectService {
	return &projectService{
		repo: repo,
	}
}

func (s projectService) CreateProject(ctx context.Context, cd shared.CreateProjectCommand) error {
	project, err := domain.NewProjectAggregate(cd)
	if err != nil {
		return err
	}

	return s.repo.Insert(ctx, project)
}

func (s projectService) AddModel(ctx context.Context, cd shared.AddModelCommand) error {
	project, err := s.repo.Get(ctx, cd.ProjectID)
	if err != nil {
		return err
	}

	err = project.AddModel(cd)
	if err != nil {
		return err
	}

	return s.repo.Save(ctx, project)
}

func (s projectService) GetProjects() ([]shared.ProjectDTO, error) {
	result := make([]shared.ProjectDTO, 0)
	return result, nil
}
