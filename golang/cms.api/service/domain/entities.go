package domain

import (
	"github.com/google/uuid"
)

type Model struct {
	ID          uint         `bson:"id"`
	ProjectID   uuid.UUID    `bson:"project_id"`
	ApiID       ApiID        `bson:"api_id"`
	Name        string       `bson:"name"`
	Description string       `bson:"description"`
	Fields      []ModelField `bson:"fields"`
}
