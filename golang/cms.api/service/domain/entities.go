package domain

import (
	"github.com/cioti/monorepo/pkg/api"
	"github.com/cioti/monorepo/pkg/datetime"
	"github.com/cioti/monorepo/pkg/storage/mongo"
	"github.com/google/uuid"
)

type Model struct {
	ProjectID    mongo.MUUID       `bson:"project_id"`
	ApiID        string            `bson:"api_id"`
	Name         string            `bson:"name"`
	Description  string            `bson:"description"`
	Fields       []Field           `bson:"fields"`
	Contents     []Content         `bson:"content"`
	DateCreated  datetime.DateTime `bson:"date_created"`
	DateModified datetime.DateTime `bson:"date_modified"`
}

func NewModel(projectID uuid.UUID, apiID string, name string, description string) (Model, error) {
	if projectID == uuid.Nil {
		return Model{}, api.NewBadRequestErrorf("Project ID cannot be nil")
	}
	if len(apiID) == 0 {
		return Model{}, api.NewBadRequestErrorf("Model apiID cannot be empty")
	}
	if len(name) == 0 {
		return Model{}, api.NewBadRequestErrorf("Model name cannot be empty")
	}

	return Model{
		ProjectID:   mongo.NewMUUID(projectID),
		ApiID:       apiID,
		Name:        name,
		Description: description,
		DateCreated: datetime.Now(),
	}, nil
}
