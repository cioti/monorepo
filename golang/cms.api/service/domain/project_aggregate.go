package domain

import (
	"github.com/cioti/monorepo/cms.api/shared"
	"github.com/cioti/monorepo/pkg/api"
	"github.com/cioti/monorepo/pkg/datetime"
	"github.com/cioti/monorepo/pkg/storage/mongo"
	"github.com/google/uuid"
	"github.com/samber/lo"
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

func (p *ProjectAggregate) AddModel(cd shared.AddModelCommand) (Model, error) {
	model, err := NewModel(cd.ApiID, cd.Name, cd.Description)
	if err != nil {
		return Model{}, err
	}
	for _, m := range p.Models {
		if m.Name == cd.Name {
			return Model{}, api.NewBadRequestErrorf("Model with name '%s' already exists", cd.Name)
		}
		if m.ApiID.String() == cd.ApiID {
			return Model{}, api.NewBadRequestErrorf("Model with apiID '%s' already exists", cd.ApiID)
		}
	}

	p.Models = append(p.Models, model)

	return model, nil
}

func (p *ProjectAggregate) AddModelField(cd shared.AddModelFieldCommand, ft FieldType) (Field, error) {
	model, found := lo.Find(p.Models, func(item Model) bool {
		return item.ApiID.String() == cd.ModelApiID
	})
	if !found {
		return Field{}, api.NewBadRequestErrorf("Model with apiID '%s' not found", cd.FieldApiID)
	}

	return model.AddField(cd.Name, cd.FieldApiID, cd.Description, ft)
}

func (p *ProjectAggregate) AddContent(cd shared.AddContentCommand, vp ValidationProcessor) (Content, error) {
	model, found := lo.Find(p.Models, func(item Model) bool {
		return item.ApiID.String() == cd.ModelApiID
	})
	if !found {
		return Content{}, api.NewBadRequestErrorf("Model with apiID '%s' not found", cd.ModelApiID)
	}
	field, found := lo.Find(model.Fields, func(item Field) bool {
		return item.ApiID.String() == cd.FieldApiID
	})
	if !found {
		return Content{}, api.NewBadRequestErrorf("Field with apiID '%s' not found", cd.FieldApiID)
	}

	isValid, err := vp.Validate(cd.Value, field)
	if err != nil {
		return Content{}, err
	}

	if !isValid {
		return Content{}, api.NewBadRequestErrorf("Field validation failed")
	}

	return model.AddContent(cd.FieldApiID, cd.Value)
}
