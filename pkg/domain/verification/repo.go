package verification

import (
	"context"

	"github.com/e-fish/api/pkg/common/helper/config"
	"github.com/e-fish/api/pkg/common/infra/orm"
	"gorm.io/gorm"
)

func NewRepo(dbConfig config.DbConfig) (Repo, error) {
	db, err := orm.CreateConnetionDB(dbConfig)
	if err != nil {
		return nil, err
	}

	return &Verification{
		DbConfig: dbConfig,
		db:       db,
	}, err
}

type Verification struct {
	DbConfig config.DbConfig
	db       *gorm.DB
}

// NewCommand implements Repo.
func (a *Verification) NewCommand(ctx context.Context) Command {
	return newCommand(ctx, a.db)
}

// NewQuery implements Repo.
func (a *Verification) NewQuery() Query {
	return newQuery(a.db)
}
