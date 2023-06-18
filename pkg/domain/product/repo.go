package product

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

	return &ProductRepo{
		DbConfig: dbConfig,
		db:       db,
	}, err
}

type ProductRepo struct {
	DbConfig config.DbConfig
	db       *gorm.DB
}

// NewCommand implements Repo.
func (a *ProductRepo) NewCommand(ctx context.Context) Command {
	return newCommand(ctx, a.db)
}

// NewQuery implements Repo.
func (a *ProductRepo) NewQuery() Query {
	return newQuery(a.db)
}
