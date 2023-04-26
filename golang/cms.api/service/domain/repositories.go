package domain

import (
	"context"

	"github.com/google/uuid"
)

type ProjectRepository interface {
	Insert(ctx context.Context, project ProjectAggregate) error
	Get(ctx context.Context, id uuid.UUID) (ProjectAggregate, error)
	Save(ctx context.Context, project ProjectAggregate) error
}

type FieldTypeRepository interface {
	Insert(ctx context.Context, fieldType FieldType) error
	Get(ctx context.Context, id uint) (FieldType, error)
}
