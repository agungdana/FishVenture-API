package transaction

import (
	"context"
	"fmt"
	"time"

	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/e-fish/api/pkg/common/helper/werror"
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

// UpdateCancelOrder implements Command.
func (c *command) UpdateCancelOrder(ctx context.Context, input uuid.UUID) (*uuid.UUID, error) {
	var (
		userID, _ = ctxutil.GetUserID(ctx)
		today     = time.Now()
	)

	exist, err := c.query.lock().ReadOrderByID(ctx, input)
	if err != nil {
		return nil, err
	}

	isTrue, ok := model.ValidateStatus[exist.Status][model.CANCEL]
	if !ok || !isTrue {
		return nil, werror.Error{
			Code:    errortransaction.ErrUpdateOrderStatus.Code,
			Message: fmt.Sprintf("the order status has [%s], failed to update the order [%s]", exist.Status, model.CANCEL),
		}
	}

	err = c.dbTxn.Where("deleted_at IS NULL and id = ?", input).Updates(
		&model.Order{
			ID:     input,
			Status: model.CANCEL,
			OrmModel: orm.OrmModel{
				UpdatedBy: &userID,
				UpdatedAt: &today,
			},
		},
	).Error

	if err != nil {
		return nil, err
	}

	return &input, nil
}

// UpdateSuccesOrder implements Command.
func (c *command) UpdateSuccesOrder(ctx context.Context, input uuid.UUID) (*uuid.UUID, error) {
	var (
		userID, _ = ctxutil.GetUserID(ctx)
		today     = time.Now()
	)

	exist, err := c.query.lock().ReadOrderByID(ctx, input)
	if err != nil {
		return nil, err
	}

	isTrue, ok := model.ValidateStatus[exist.Status][model.SUCCESS]
	if !ok || !isTrue {
		return nil, werror.Error{
			Code:    errortransaction.ErrUpdateOrderStatus.Code,
			Message: fmt.Sprintf("the order status has [%s], failed to update the order [%s]", exist.Status, model.SUCCESS),
		}
	}

	err = c.dbTxn.Where("deleted_at IS NULL and id = ?", input).Updates(
		&model.Order{
			ID:     input,
			Status: model.SUCCESS,
			OrmModel: orm.OrmModel{
				UpdatedBy: &userID,
				UpdatedAt: &today,
			},
		},
	).Error

	if err != nil {
		return nil, err
	}

	return &input, nil
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
