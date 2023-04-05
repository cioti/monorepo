package domain

import (
	"github.com/cioti/monorepo/pkg/datetime"
	"github.com/google/uuid"
)

type ContentAggregate struct {
	ID           uint
	ProjectID    uuid.UUID
	ModelApiID   ApiID
	Fields       []ContentField
	DateCreated  datetime.DateTime
	DateModified datetime.DateTime
}
