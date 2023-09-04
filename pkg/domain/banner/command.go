package banner

import (
	"context"

	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/e-fish/api/pkg/domain/banner/model"
	errorBanner "github.com/e-fish/api/pkg/domain/banner/model/error-banner"
	errorregion "github.com/e-fish/api/pkg/domain/region/model/error-region"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func newCommand(ctx context.Context, db *gorm.DB) Command {
	var (
		dbTxn = orm.BeginTxn(ctx, db.WithContext(ctx))
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

// UpdateBanner implements Command.
func (c *command) UpdateBanner(ctx context.Context, input model.BannerInputUpdate) (*uuid.UUID, error) {
	var (
		userID, _ = ctxutil.GetUserID(ctx)
		err       = input.Validate()
	)

	if err != nil {
		return nil, err
	}

	updatedBanner := input.NewBanner(userID)

	err = c.dbTxn.Where("deleted_at IS NULL and id = ?", updatedBanner.ID).Updates(&updatedBanner).Error
	if err != nil {
		return nil, errorBanner.ErrUpdateBanner.AttacthDetail(map[string]any{"id": input.ID, "error": err})
	}

	return &updatedBanner.ID, nil
}

// CreateBanner implements Command.
func (c *command) CreateBanner(ctx context.Context, input model.BannerInputCreate) (*uuid.UUID, error) {
	var (
		userID, _ = ctxutil.GetUserID(ctx)
		err       = input.Validate()
	)

	if err != nil {
		return nil, err
	}

	newBanner := input.NewBanner(userID)

	err = c.dbTxn.Create(&newBanner).Error
	if err != nil {
		return nil, errorBanner.ErrCreateBanner.AttacthDetail(map[string]any{"error": err})
	}
	return &newBanner.ID, nil
}

// Commit implements Command.
func (c *command) Commit(ctx context.Context) error {
	if err := orm.CommitTxn(ctx); err != nil {
		return errorregion.ErrCommit.AttacthDetail(map[string]any{"error": err})
	}
	return nil
}

// Rollback implements Command.
func (c *command) Rollback(ctx context.Context) error {
	if err := orm.RollbackTxn(ctx); err != nil {
		return errorregion.ErrRollback.AttacthDetail(map[string]any{"error": err})
	}
	return nil
}
