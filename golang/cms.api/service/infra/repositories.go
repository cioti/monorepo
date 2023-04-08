package infra

import (
	"context"

	"github.com/cioti/monorepo/cms.api/service/domain"
	mongopkg "github.com/cioti/monorepo/pkg/storage/mongo"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
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

func (r projectRepository) Get(ctx context.Context, id uuid.UUID) (domain.ProjectAggregate, error) {
	var agg domain.ProjectAggregate
	result := r.database.Collection(collection).FindOne(ctx, bson.M{"_id": mongopkg.NewMUUID(id)})
	err := result.Decode(&agg)
	return agg, err
}

func (r projectRepository) Save(ctx context.Context, project domain.ProjectAggregate) error {
	filter := bson.D{{Key: "_id", Value: project.ID}}
	result, err := r.database.Collection(collection).ReplaceOne(ctx, filter, project)
	if err != nil {
		return err
	}
	if result.UpsertedCount != 1 {
		return nil
	}

	return nil
}
