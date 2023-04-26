package domain

import (
	"github.com/cioti/monorepo/pkg/api"
	"github.com/cioti/monorepo/pkg/datetime"
	"github.com/samber/lo"
)

type Model struct {
	ApiID        ApiID             `bson:"api_id"`
	Name         string            `bson:"name"`
	Description  string            `bson:"description"`
	Fields       []Field           `bson:"fields"`
	Contents     []Content         `bson:"content"`
	DateCreated  datetime.DateTime `bson:"date_created"`
	DateModified datetime.DateTime `bson:"date_modified"`
}

func NewModel(apiID string, name string, description string) (Model, error) {
	if len(name) == 0 {
		return Model{}, api.NewBadRequestErrorf("Model name cannot be empty")
	}
	apiId, err := NewApiID(apiID)
	if err != nil {
		return Model{}, err
	}

	return Model{
		ApiID:       apiId,
		Name:        name,
		Description: description,
		DateCreated: datetime.Now(),
		Fields:      make([]Field, 0),
		Contents:    make([]Content, 0),
	}, nil
}

func (m *Model) AddField(name string, apiID string, desc string, fieldType FieldType) (Field, error) {
	field, err := NewField(name, apiID, desc, fieldType)
	if err != nil {
		return Field{}, err
	}
	alreadyExists := lo.ContainsBy(m.Fields, func(item Field) bool {
		return item.ApiID.String() == apiID || item.Name == name
	})
	if alreadyExists {
		return Field{}, api.NewBadRequestErrorf("Field with name '%s' and apiID '%s' already exists", name, apiID)
	}
	m.Fields = append(m.Fields, field)
	m.DateModified = datetime.Now()

	return field, nil
}

func (m *Model) AddContent(apiID string, value interface{}) (Content, error) {
	content, err := NewContent(apiID, value)
	if err != nil {
		return Content{}, err
	}
	m.DateModified = datetime.Now()

	return content, nil
}

type FieldType struct {
	ID              uint             `bson:"_id"`
	Name            string           `bson:"name"`
	DataType        DataType         `bson:"data_type"`
	ValidationRules []ValidationRule `bson:"validation_rules"`
}

func NewFieldType(name string, dt DataType) (FieldType, error) {
	if len(name) == 0 {
		return FieldType{}, api.NewBadRequestErrorf("FieldType name cannot be empty")
	}
	if len(dt) == 0 {
		return FieldType{}, api.NewBadRequestErrorf("FieldType dataType cannot be empty")
	}

	return FieldType{
		Name:            name,
		DataType:        dt,
		ValidationRules: make([]ValidationRule, 0),
	}, nil
}

func (ft *FieldType) AddValidation(logic ValidationRuleLogic, target interface{}) (ValidationRule, error) {
	rule, err := NewValidationRule(logic, target)
	if err != nil {
		return ValidationRule{}, err
	}
	alreadyExists := lo.ContainsBy(ft.ValidationRules, func(item ValidationRule) bool {
		return item.Logic == logic
	})
	if alreadyExists {
		return ValidationRule{}, api.NewBadRequestErrorf("Unable to add field type validation, validation rule logic already exists")
	}
	ft.ValidationRules = append(ft.ValidationRules, rule)

	return rule, nil
}
