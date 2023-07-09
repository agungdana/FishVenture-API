package budidaya

import (
	"context"

	"github.com/e-fish/api/pkg/common/helper/config"
	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/e-fish/api/pkg/domain/pond"
	"gorm.io/gorm"
)

func NewRepo(dbConfig config.DbConfig, pondRrepo pond.Repo) (Repo, error) {
	db, err := orm.CreateConnetionDB(dbConfig)
	if err != nil {
		return nil, err
	}

	return &Budidaya{
		DbConfig: dbConfig,
		db:       db,
		pondRepo: pondRrepo,
	}, err
}

type Budidaya struct {
	DbConfig config.DbConfig
	db       *gorm.DB
	pondRepo pond.Repo
}

// NewCommand implements Repo.
func (a *Budidaya) NewCommand(ctx context.Context) Command {
	return newCommand(ctx, a.db, a.pondRepo)
}

// NewQuery implements Repo.
func (a *Budidaya) NewQuery() Query {
	return newQuery(a.db)
}
