package domain

import "context"

type ProjectRepository interface {
	Insert(ctx context.Context, project ProjectAggregate) error
}
