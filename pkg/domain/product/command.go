package product

import (
	"context"
	"time"

	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/e-fish/api/pkg/common/infra/orm"
	errorproduct "github.com/e-fish/api/pkg/domain/product/error-product"
	"github.com/e-fish/api/pkg/domain/product/model"
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

// CreateProduct implements Command.
func (c *command) CreateProduct(ctx context.Context, input model.CreateProductInput) (*uuid.UUID, error) {
	var (
		userID, _ = ctxutil.GetUserID(ctx)
		pondID, _ = ctxutil.GetPondID(ctx)
	)

	err := input.Validate()
	if err != nil {
		return nil, err
	}

	exist, err := c.query.ReadProductByBudidayaID(ctx, input.BudidayaID)
	if err != nil {
		if !errorproduct.ErrFoundProduct.Is(err) {
			return nil, err
		}
	}

	if exist != nil {
		return nil, errorproduct.ErrCreateProductExist
	}

	newProduct := input.ToProduct(userID, pondID)

	err = c.dbTxn.Create(&newProduct).Error
	if err != nil {
		return nil, errorproduct.ErrCreateProduct.AttacthDetail(map[string]any{"error": err})
	}

	return &newProduct.ID, nil
}

// DeleteProduct implements Command.
func (c *command) DeleteProduct(ctx context.Context, input uuid.UUID) (*uuid.UUID, error) {
	var (
		userID, _ = ctxutil.GetUserID(ctx)
		today     = time.Now()
	)

	err := c.dbTxn.Where("deleted_at IS NULL and id =?", input).Updates(&model.Product{
		ID: input,
		OrmModel: orm.OrmModel{
			DeletedAt: &today,
			DeletedBy: &userID,
		},
	}).Error
	if err != nil {
		return nil, errorproduct.ErrDeleteProduct.AttacthDetail(map[string]any{"error": err})
	}

	return &input, nil
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
