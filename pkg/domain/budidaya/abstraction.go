package product

import (
	"context"

	"github.com/e-fish/api/pkg/domain/budidaya/model"
	"github.com/google/uuid"
)

type Repo interface {
	NewCommand(ctx context.Context) Command
	NewQuery() Query
}

type Command interface {
	CreatePond(ctx context.Context, input model.CreatePondInput) (*uuid.UUID, error)

	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
}

type Query interface {
	lock() Query
}
