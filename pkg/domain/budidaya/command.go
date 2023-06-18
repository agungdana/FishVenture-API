package product

import (
	"context"

	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/e-fish/api/pkg/domain/budidaya/model"
	errorauth "github.com/e-fish/api/pkg/domain/product/error"
	"github.com/e-fish/api/pkg/domain/verification"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func newCommand(ctx context.Context, db *gorm.DB, verificationRepo verification.Repo) Command {
	var (
		dbTxn = orm.BeginTxn(ctx, db)
	)

	return &command{
		dbTxn:            dbTxn,
		query:            newQuery(dbTxn),
		verificationRepo: verificationRepo,
	}
}

type command struct {
	dbTxn            *gorm.DB
	query            Query
	verificationRepo verification.Repo
}

// CreatePond implements Command.
func (c *command) CreatePond(ctx context.Context, input model.CreatePondInput) (*uuid.UUID, error) {
	var (
		userID, _ = ctxutil.GetUserID(ctx)
	)

	err := input.Validate()
	if err != nil {
		return nil, err
	}

	newPond := input.ToPond(userID)

	err = c.dbTxn.Create(&newPond).Error
	if err != nil {
		return nil, err
	}

	return &newPond.ID, nil
}

// Commit implements Command.
func (c *command) Commit(ctx context.Context) error {
	if err := orm.CommitTxn(ctx); err != nil {
		return errorauth.ErrCommit.AttacthDetail(map[string]any{"errors": err})
	}
	return nil
}

// Rollback implements Command.
func (c *command) Rollback(ctx context.Context) error {
	if err := orm.RollbackTxn(ctx); err != nil {
		return errorauth.ErrRollback.AttacthDetail(map[string]any{"errors": err})
	}
	return nil
}
