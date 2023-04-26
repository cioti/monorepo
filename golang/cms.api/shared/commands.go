package shared

import "github.com/google/uuid"

type CreateProjectCommand struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type AddModelCommand struct {
	ProjectID   uuid.UUID `json:"projectId"`
	ApiID       string    `json:"apiId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type AddModelFieldCommand struct {
	ProjectID   uuid.UUID `json:"projectId"`
	ModelApiID  string    `json:"modelApiId"`
	FieldApiID  string    `json:"fieldApiId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	FieldTypeID uint      `json:"fieldTypeId"`
}

type AddContentCommand struct {
	ProjectID  uuid.UUID   `json:"projectId"`
	ModelApiID string      `json:"modelApiId"`
	FieldApiID string      `json:"fieldApiId"`
	Value      interface{} `json:"value"`
}
