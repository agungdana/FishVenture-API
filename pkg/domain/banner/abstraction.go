package banner

import (
	"context"

	"github.com/e-fish/api/pkg/domain/banner/model"
	"github.com/google/uuid"
)

type Repo interface {
	NewCommand(ctx context.Context) Command
	NewQuery() Query
}

type Command interface {
	CreateBanner(ctx context.Context, input model.BannerInputCreate) (*uuid.UUID, error)
	UpdateBanner(ctx context.Context, input model.BannerInputUpdate) (*uuid.UUID, error)

	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
}

type Query interface {
	ReadAllBanner(ctx context.Context) ([]model.BannerOutput, error)

	lock() Query
}
