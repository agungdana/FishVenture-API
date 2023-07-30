package verification

import (
	"context"
	"time"

	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/e-fish/api/pkg/common/infra/orm"
	errorverification "github.com/e-fish/api/pkg/domain/verification/error_verification"
	"github.com/e-fish/api/pkg/domain/verification/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func newCommand(ctx context.Context, db *gorm.DB) Command {
	var (
		dbTxn = orm.BeginTxn(ctx, db)
	)

	return &command{
		dbTxn: dbTxn.WithContext(ctx),
		query: newQuery(dbTxn),
	}
}

type command struct {
	dbTxn *gorm.DB
	query Query
}

// DeleteOTP implements Command.
func (c *command) DeleteOTP(ctx context.Context, input uuid.UUID) (*uuid.UUID, error) {
	var (
		userID, _ = ctxutil.GetUserID(ctx)
		today     = time.Now()
	)

	err := c.dbTxn.Unscoped().Where("deleted_at IS NULL and id = ?", userID).Delete(&model.OTP{
		ID:     input,
		UserID: userID,
		OrmModel: orm.OrmModel{
			DeletedAt: &today,
			DeletedBy: &userID,
		},
	}).Error

	if err != nil {
		return nil, errorverification.ErrFailedDeleteCodeOTP.AttacthDetail(map[string]any{"error": err})
	}

	return &input, nil
}

// CreateRandCode implements Command.
func (c *command) CreateOTP(ctx context.Context, input model.CreateCodeOTPInput) (*uuid.UUID, error) {
	userID, _ := ctxutil.GetUserID(ctx)

	if err := input.Validate(); err != nil {
		return nil, err
	}

	newOTP := input.ToOTP(userID)

	err := c.dbTxn.Create(&newOTP).Error
	if err != nil {
		return nil, errorverification.ErrFailedCreateCodeOTP.AttacthDetail(map[string]any{"error": err})
	}

	return &newOTP.ID, nil
}

// Commit implements Command.
func (c *command) Commit(ctx context.Context) error {
	if err := orm.CommitTxn(ctx); err != nil {
		return errorverification.ErrCommit.AttacthDetail(map[string]any{"errors": err})
	}
	return nil
}

// Rollback implements Command.
func (c *command) Rollback(ctx context.Context) error {
	if err := orm.RollbackTxn(ctx); err != nil {
		return errorverification.ErrRollback.AttacthDetail(map[string]any{"errors": err})
	}
	return nil
}
