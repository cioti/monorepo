package app

import (
	"context"

	"github.com/cioti/monorepo/cms.api/service/domain"
	"github.com/cioti/monorepo/cms.api/shared"
)

type ProjectService interface {
	CreateProject(ctx context.Context, cd shared.CreateProjectCD) error
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

func (s projectService) CreateProject(ctx context.Context, cd shared.CreateProjectCD) error {
	project, err := domain.NewProjectAggregate(cd.ID, cd.Name, cd.Description)
	if err != nil {
		return err
	}

	return s.repo.Insert(ctx, project)
}

func (s projectService) GetProjects() ([]shared.ProjectDTO, error) {
	result := make([]shared.ProjectDTO, 0)
	return result, nil
}
