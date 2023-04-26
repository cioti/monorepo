package app

import (
	"context"

	"github.com/cioti/monorepo/cms.api/service/domain"
	"github.com/cioti/monorepo/cms.api/shared"
)

type ProjectService interface {
	CreateProject(ctx context.Context, cd shared.CreateProjectCommand) error
	AddModel(ctx context.Context, cd shared.AddModelCommand) error
	AddModelField(ctx context.Context, cd shared.AddModelFieldCommand) error
	GetProjects() ([]shared.ProjectDTO, error)
}

type projectService struct {
	repo   domain.ProjectRepository
	ftRepo domain.FieldTypeRepository
	vp     domain.ValidationProcessor
}

func NewProjectService(repo domain.ProjectRepository, ftRepo domain.FieldTypeRepository) ProjectService {
	return &projectService{
		repo:   repo,
		ftRepo: ftRepo,
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

	_, err = project.AddModel(cd)
	if err != nil {
		return err
	}

	return s.repo.Save(ctx, project)
}

func (s projectService) AddModelField(ctx context.Context, cd shared.AddModelFieldCommand) error {
	project, err := s.repo.Get(ctx, cd.ProjectID)
	if err != nil {
		return err
	}

	ft, err := s.ftRepo.Get(ctx, cd.FieldTypeID)
	if err != nil {
		return err
	}

	_, err = project.AddModelField(cd, ft)
	if err != nil {
		return err
	}

	return s.repo.Save(ctx, project)
}

func (s projectService) AddContent(ctx context.Context, cd shared.AddContentCommand) error {
	project, err := s.repo.Get(ctx, cd.ProjectID)
	if err != nil {
		return err
	}

	_, err = project.AddContent(cd, s.vp)
	return err
}

func (s projectService) GetProjects() ([]shared.ProjectDTO, error) {
	result := make([]shared.ProjectDTO, 0)
	return result, nil
}
