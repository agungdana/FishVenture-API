package transaction

import (
	"context"

	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/e-fish/api/pkg/domain/budidaya"
	modelBudidaya "github.com/e-fish/api/pkg/domain/budidaya/model"
	errortransaction "github.com/e-fish/api/pkg/domain/transaction/error-transaction"
	"github.com/e-fish/api/pkg/domain/transaction/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func newCommand(ctx context.Context, db *gorm.DB, budidayaRepo budidaya.Repo) Command {
	var (
		dbTxn = orm.BeginTxn(ctx, db)
	)

	return &command{
		dbTxn:         dbTxn.WithContext(ctx),
		query:         newQuery(dbTxn),
		budidayaQuery: budidayaRepo.NewQuery(),
	}
}

type command struct {
	dbTxn         *gorm.DB
	query         Query
	budidayaQuery budidaya.Query
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

	price, err := c.budidayaQuery.ReadPriceListBudidayaByBiggerThanLimitAndBudidayaID(ctx, modelBudidaya.ReadPricelistBudidayaInput{
		BudidayaID: input.BudidayaID,
		Qty:        input.Qty,
	})

	if err != nil {
		return nil, err
	}

	if price == nil {
		return nil, errortransaction.ErrCreateOrder.AttacthDetail(map[string]any{"price": "empty"})
	}

	newOrder := input.ToOrder(userID, *price)

	err = c.dbTxn.Create(&newOrder).Error
	if err != nil {
		return nil, errortransaction.ErrCreateOrder.AttacthDetail(map[string]any{"error": err})
	}

	return &newOrder.ID, nil
}

// Commit implements Command.
func (c *command) Commit(ctx context.Context) error {
	if err := orm.CommitTxn(ctx); err != nil {
		return errortransaction.ErrCommit.AttacthDetail(map[string]any{"errors": err})
	}
	return nil
}

// Rollback implements Command.
func (c *command) Rollback(ctx context.Context) error {
	if err := orm.RollbackTxn(ctx); err != nil {
		return errortransaction.ErrRollback.AttacthDetail(map[string]any{"errors": err})
	}
	return nil
}
