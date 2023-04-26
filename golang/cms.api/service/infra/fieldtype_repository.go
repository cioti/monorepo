package infra

import (
	"context"

	"github.com/cioti/monorepo/cms.api/service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ftCollection = "field_types"
)

type fieldTypeRepository struct {
	database mongo.Database
}

func NewFieldTypeRepository(database mongo.Database) domain.FieldTypeRepository {
	return &fieldTypeRepository{
		database: database,
	}
}

func (r fieldTypeRepository) Insert(ctx context.Context, fieldType domain.FieldType) error {
	_, err := r.database.Collection(ftCollection).InsertOne(ctx, fieldType)
	return err
}

func (r fieldTypeRepository) Get(ctx context.Context, id uint) (domain.FieldType, error) {
	var ft domain.FieldType
	result := r.database.Collection(ftCollection).FindOne(ctx, bson.M{"_id": id})
	err := result.Decode(&ft)
	return ft, err
}
