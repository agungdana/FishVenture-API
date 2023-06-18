package product

import (
	"context"

	"github.com/e-fish/api/pkg/common/infra/orm"
	errorauth "github.com/e-fish/api/pkg/domain/product/error"
	"gorm.io/gorm"
)

func newCommand(ctx context.Context, db *gorm.DB) Command {
	var (
		dbTxn = orm.BeginTxn(ctx, db)
	)

	return &command{
		dbTxn: dbTxn,
		query: newQuery(dbTxn),
	}
}

type command struct {
	dbTxn *gorm.DB
	query Query
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
