package shared

import "github.com/google/uuid"

type CreateProjectCD struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}
