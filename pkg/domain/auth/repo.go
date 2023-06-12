package auth

import (
	"context"

	"github.com/e-fish/api/pkg/common/helper/config"
	"github.com/e-fish/api/pkg/common/infra/firebase"
	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/e-fish/api/pkg/common/infra/token"
	"gorm.io/gorm"
)

func NewRepo(dbConfig config.DbConfig, maker token.Token, fireBase firebase.Firebase) (Repo, error) {
	ctx := context.Background()
	db, err := orm.CreateConnetionDB(dbConfig)
	if err != nil {
		return nil, err
	}

	gauth, err := fireBase.NewGoogleAuth(ctx)
	if err != nil {
		return nil, err
	}

	return &AuthRepo{
		DbConfig:   dbConfig,
		tokenMaker: maker,
		gauth:      gauth,
		db:         db,
	}, err
}

type AuthRepo struct {
	DbConfig   config.DbConfig
	tokenMaker token.Token
	gauth      firebase.GoogleAuth
	db         *gorm.DB
}

// NewCommand implements Repo.
func (a *AuthRepo) NewCommand(ctx context.Context) Command {
	return newCommand(ctx, a.db, a.tokenMaker, a.gauth)
}

// NewQuery implements Repo.
func (a *AuthRepo) NewQuery() Query {
	return newQuery(a.db)
}
