package domain

import (
	"github.com/cioti/monorepo/pkg/api"
	"github.com/cioti/monorepo/pkg/datetime"
	"github.com/google/uuid"
)

const (
	NilIDErr     = "Project ID cannot be nil"
	EmptyNameErr = "Project name cannot be empty"
)

type ProjectAggregate struct {
	ID          uuid.UUID         `bson:"id"`
	Name        string            `bson:"name"`
	Description string            `bson:"description"`
	Models      []Model           `bson:"models"`
	DateCreated datetime.DateTime `bson:"date_created"`
}

func NewProjectAggregate(id uuid.UUID, name string, description string) (ProjectAggregate, error) {
	if id == uuid.Nil {
		return ProjectAggregate{}, api.NewBadRequestErrorf(NilIDErr)
	}
	if len(name) == 0 {
		return ProjectAggregate{}, api.NewBadRequestErrorf(EmptyNameErr)
	}

	return ProjectAggregate{
		ID:          id,
		Name:        name,
		Description: description,
		Models:      make([]Model, 0),
		DateCreated: datetime.Now(),
	}, nil
}
