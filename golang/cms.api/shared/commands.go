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
