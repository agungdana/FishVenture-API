package budidaya

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
	UpdatePond(ctx context.Context, input model.UpdatePondInput) (*uuid.UUID, error)
	UpdatePondStatus(ctx context.Context, input model.UpdatePondStatus) (*uuid.UUID, error)

	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
}

type Query interface {
	GetListPondSubmission(ctx context.Context) ([]*model.PondOutput, error)
	GetPondByUserID(ctx context.Context) (*model.PondOutput, error)
	GetPondByID(ctx context.Context, input uuid.UUID) (*model.PondOutput, error)

	lock() Query
}
