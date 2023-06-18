package product

import (
	"context"

	"github.com/e-fish/api/pkg/common/helper/config"
	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/e-fish/api/pkg/domain/verification"
	"gorm.io/gorm"
)

func NewRepo(dbConfig config.DbConfig, verificationRepo verification.Repo) (Repo, error) {
	db, err := orm.CreateConnetionDB(dbConfig)
	if err != nil {
		return nil, err
	}

	return &Budidaya{
		DbConfig:         dbConfig,
		db:               db,
		verificationRepo: verificationRepo,
	}, err
}

type Budidaya struct {
	DbConfig         config.DbConfig
	db               *gorm.DB
	verificationRepo verification.Repo
}

// NewCommand implements Repo.
func (a *Budidaya) NewCommand(ctx context.Context) Command {
	return newCommand(ctx, a.db, a.verificationRepo)
}

// NewQuery implements Repo.
func (a *Budidaya) NewQuery() Query {
	return newQuery(a.db)
}
