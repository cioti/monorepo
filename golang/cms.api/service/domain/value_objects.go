package domain

import (
	"regexp"

	"github.com/cioti/monorepo/pkg/api"
	"github.com/samber/lo"
)

var (
	apiIDRegex     = regexp.MustCompile(`^[a-zA-Z0-9_-]*$`)
	apiIDMaxLength = 50
)

type ApiID string

func NewApiID(id string) (ApiID, error) {
	if len(id) == 0 {
		return "", api.NewBadRequestErrorf("ApiID cannot be empty")
	}

	if len(id) > apiIDMaxLength {
		return "", api.NewBadRequestErrorf("ApiID length cannot exceed %d characters", apiIDMaxLength)
	}

	if !apiIDRegex.MatchString(id) {
		return "", api.NewBadRequestErrorf("ApiID must be an alpha numeric string with '-' and '_' allowed only")
	}

	return ApiID(id), nil
}

func (a ApiID) String() string {
	return string(a)
}

type Field struct {
	Name        string    `bson:"name"`
	ApiID       ApiID     `bson:"api_id"`
	Description string    `bson:"description"`
	Type        FieldType `bson:"type"`
}

func NewField(name string, apiID string, description string, fieldType FieldType) (Field, error) {
	apiId, err := NewApiID(apiID)
	if err != nil {
		return Field{}, err
	}

	return Field{
		Name:        name,
		ApiID:       apiId,
		Description: description,
		Type:        fieldType,
	}, nil
}

type Content struct {
	ApiID ApiID       `bson:"api_id"`
	Value interface{} `bson:"value"`
}

func NewContent(apiID string, value interface{}) (Content, error) {
	apiId, err := NewApiID(apiID)
	if err != nil {
		return Content{}, err
	}
	if value == nil {
		return Content{}, api.NewBadRequestErrorf("Content value cannot be nil")
	}

	return Content{
		ApiID: apiId,
		Value: value,
	}, nil
}

type ValidationRule struct {
	Logic  ValidationRuleLogic `bson:"logic"`
	Target interface{}         `bson:"target"`
}

func NewValidationRule(logic ValidationRuleLogic, target interface{}) (ValidationRule, error) {
	if !lo.Contains(ValidationRuleLogics, logic) {
		return ValidationRule{}, api.NewBadRequestErrorf("Validation rule logic '%s' is not valid", logic)
	}

	if logic == Regex {
		expr, ok := target.(string)
		if !ok {
			return ValidationRule{}, api.NewBadRequestErrorf("Validation rule target must be a string for regex type")
		}
		_, err := regexp.Compile(expr)
		if err != nil {
			return ValidationRule{}, api.NewBadRequestError(err, "Validation rule target is not a valid regex")
		}
	}

	return ValidationRule{
		Logic:  logic,
		Target: target,
	}, nil
}
