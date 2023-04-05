package shared

import (
	"github.com/google/uuid"
)

type ProjectDTO struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}
