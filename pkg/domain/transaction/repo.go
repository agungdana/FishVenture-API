package transaction

import (
	"context"

	"github.com/e-fish/api/pkg/common/helper/config"
	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/e-fish/api/pkg/domain/budidaya"
	"gorm.io/gorm"
)

func NewRepo(dbConfig config.DbConfig, budidayaRepo budidaya.Repo) (Repo, error) {
	db, err := orm.CreateConnetionDB(dbConfig)
	if err != nil {
		return nil, err
	}

	return &ProductRepo{
		DbConfig:     dbConfig,
		db:           db,
		budidayaRepo: budidayaRepo,
	}, err
}

type ProductRepo struct {
	DbConfig     config.DbConfig
	db           *gorm.DB
	budidayaRepo budidaya.Repo
}

// NewCommand implements Repo.
func (a *ProductRepo) NewCommand(ctx context.Context) Command {
	return newCommand(ctx, a.db, a.budidayaRepo)
}

// NewQuery implements Repo.
func (a *ProductRepo) NewQuery() Query {
	return newQuery(a.db)
}
