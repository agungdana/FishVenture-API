package region

import (
	"context"

	"github.com/e-fish/api/pkg/common/helper/config"
	"github.com/e-fish/api/pkg/common/infra/orm"
	"gorm.io/gorm"
)

func NewRepo(appConfig config.DbConfig) (Repo, error) {
	db, err := orm.CreateConnetionDB(appConfig)
	if err != nil {
		return nil, err
	}

	return &RegionRepo{
		db: db,
	}, err
}

type RegionRepo struct {
	db *gorm.DB
}

// NewCommand implements Repo.
func (a *RegionRepo) NewCommand(ctx context.Context) Command {
	return newCommand(ctx, a.db)
}

// NewQuery implements Repo.
func (a *RegionRepo) NewQuery() Query {
	return newQuery(a.db)
}
