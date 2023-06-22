package transaction

import (
	"context"

	"github.com/e-fish/api/pkg/domain/transaction/model"
	"github.com/google/uuid"
)

type Repo interface {
	NewCommand(ctx context.Context) Command
	NewQuery() Query
}

type Command interface {
	CreateOrder(ctx context.Context, input model.CreateOrderInput) (*uuid.UUID, error)

	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
}

type Query interface {
	lock() Query
}
