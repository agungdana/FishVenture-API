package transaction

import (
	"context"

	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/e-fish/api/pkg/common/infra/orm"
	errorproduct "github.com/e-fish/api/pkg/domain/product/error-product"
	errortransaction "github.com/e-fish/api/pkg/domain/transaction/error-transaction"
	"github.com/e-fish/api/pkg/domain/transaction/model"
	"github.com/google/uuid"
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

// CreateOrder implements Command.
func (c *command) CreateOrder(ctx context.Context, input model.CreateOrderInput) (*uuid.UUID, error) {
	var (
		userID, _ = ctxutil.GetUserID(ctx)
	)

	err := input.Validate()
	if err != nil {
		return nil, err
	}

	newOrder := input.ToOrder(userID)

	err = c.dbTxn.Create(&newOrder).Error
	if err != nil {
		return nil, errortransaction.ErrCreateOrder.AttacthDetail(map[string]any{"error": err})
	}

	return &newOrder.ID, nil
}

// Commit implements Command.
func (c *command) Commit(ctx context.Context) error {
	if err := orm.CommitTxn(ctx); err != nil {
		return errorproduct.ErrCommit.AttacthDetail(map[string]any{"errors": err})
	}
	return nil
}

// Rollback implements Command.
func (c *command) Rollback(ctx context.Context) error {
	if err := orm.RollbackTxn(ctx); err != nil {
		return errorproduct.ErrRollback.AttacthDetail(map[string]any{"errors": err})
	}
	return nil
}
