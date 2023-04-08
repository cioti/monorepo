package domain

import (
	"github.com/cioti/monorepo/cms.api/shared"
	"github.com/cioti/monorepo/pkg/api"
	"github.com/cioti/monorepo/pkg/datetime"
	"github.com/cioti/monorepo/pkg/storage/mongo"
	"github.com/google/uuid"
)

type ProjectAggregate struct {
	ID          mongo.MUUID       `bson:"_id"`
	Name        string            `bson:"name"`
	Description string            `bson:"description"`
	Models      []Model           `bson:"models"`
	DateCreated datetime.DateTime `bson:"date_created"`
}

func NewProjectAggregate(cmd shared.CreateProjectCommand) (ProjectAggregate, error) {
	if cmd.ID == uuid.Nil {
		return ProjectAggregate{}, api.NewBadRequestErrorf("Project ID cannot be nil")
	}
	if len(cmd.Name) == 0 {
		return ProjectAggregate{}, api.NewBadRequestErrorf("Project name cannot be empty")
	}

	return ProjectAggregate{
		ID:          mongo.NewMUUID(cmd.ID),
		Name:        cmd.Name,
		Description: cmd.Description,
		Models:      make([]Model, 0),
		DateCreated: datetime.Now(),
	}, nil
}

func (p *ProjectAggregate) AddModel(cd shared.AddModelCommand) error {
	model, err := NewModel(cd.ProjectID, cd.ApiID, cd.Name, cd.Description)
	if err != nil {
		return err
	}
	for _, m := range p.Models {
		if m.Name == cd.Name {
			return api.NewBadRequestErrorf("Model with name '%s' already exists", cd.Name)
		}
		if m.ApiID == cd.ApiID {
			return api.NewBadRequestErrorf("Model with apiID '%s' already exists", cd.ApiID)
		}
	}

	p.Models = append(p.Models, model)

	return nil
}
