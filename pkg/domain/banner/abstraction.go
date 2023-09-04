package banner
import (
	"context"

	"github.com/e-fish/api/pkg/domain/banner/model"
)

type Repo interface {
	NewCommand(ctx context.Context) Command
	NewQuery() Query
}

type Command interface {
	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
}

type Query interface {
	ReadAllBanner(ctx context.Context) ([]model.BannerOutput, error)

	lock() Query
}
