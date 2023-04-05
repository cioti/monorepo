package infra

import (
	"context"

	"github.com/cioti/monorepo/cms.api/service/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collection = "projects"
)

type projectRepository struct {
	database mongo.Database
}

func NewProjectRepository(database mongo.Database) domain.ProjectRepository {
	return &projectRepository{
		database: database,
	}
}

func (r projectRepository) Insert(ctx context.Context, project domain.ProjectAggregate) error {
	_, err := r.database.Collection(collection).InsertOne(ctx, project)
	return err
}
