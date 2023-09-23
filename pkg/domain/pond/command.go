package pond

import (
	"context"

	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/e-fish/api/pkg/common/infra/orm"
	errorpond "github.com/e-fish/api/pkg/domain/pond/error-pond"
	"github.com/e-fish/api/pkg/domain/pond/model"
	"github.com/e-fish/api/pkg/domain/verification"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func newCommand(ctx context.Context, db *gorm.DB, verificationRepo verification.Repo) Command {
	var (
		dbTxn = orm.BeginTxn(ctx, db)
	)

	return &command{
		dbTxn:            dbTxn.WithContext(ctx),
		query:            newQuery(dbTxn),
		verificationRepo: verificationRepo,
	}
}

type command struct {
	dbTxn            *gorm.DB
	query            Query
	verificationRepo verification.Repo
}

// UpdatePond implements Command.
func (c *command) UpdatePond(ctx context.Context, input model.UpdatePondInput) (*uuid.UUID, error) {
	var (
		userID, _ = ctxutil.GetUserID(ctx)
		pondID, _ = ctxutil.GetPondID(ctx)
	)

	updatePond := input.ToPond(userID, pondID)

	err := c.dbTxn.Where("deleted_at IS NULL and id = ? and user_id = ?", pondID, userID).Updates(&updatePond).Error
	if err != nil {
		return nil, errorpond.ErrFailedUpdatePond.AttacthDetail(map[string]any{"error": err})
	}

	return &updatePond.ID, nil
}

// UpdatePondStatus implements Command.
func (c *command) UpdatePondStatus(ctx context.Context, input model.UpdatePondStatus) (*uuid.UUID, error) {
	var (
		userID, _ = ctxutil.GetUserID(ctx)
	)

	err := input.Validate()
	if err != nil {
		return nil, err
	}

	updatePond := input.ToPond(userID)

	pond, err := c.query.GetPondByID(ctx, input.PondID)
	if err != nil {
		return nil, err
	}

	val, ok := model.MapStatus[pond.Status][updatePond.Status]
	if !ok || !val {
		return nil, errorpond.ErrCannotUpdateStatusPond.AttacthDetail(map[string]any{"already-status": pond.Status, "targer-status": updatePond.Status})
	}

	err = c.dbTxn.Where("deleted_at IS NULL and id = ?", updatePond.ID).Updates(&updatePond).Error
	if err != nil {
		return nil, errorpond.ErrFailedUpdatePond.AttacthDetail(map[string]any{"error": err})
	}

	return &updatePond.ID, nil

}

// CreatePond implements Command.
func (c *command) CreatePond(ctx context.Context, input model.CreatePondInput) (*uuid.UUID, error) {
	var (
		userID, _ = ctxutil.GetUserID(ctx)
		pondID, _ = ctxutil.GetPondID(ctx)
	)

	err := input.Validate()
	if err != nil {
		return nil, err
	}

	newPond := input.ToPond(userID, pondID)

	err = c.dbTxn.Create(&newPond).Error
	if err != nil {
		return nil, err
	}

	return &newPond.ID, nil
}

// CreatePond implements Command.
func (c *command) ResubmissionPond(ctx context.Context, input model.Resubmission) (*uuid.UUID, error) {
	var (
		userID, _ = ctxutil.GetUserID(ctx)
		pondID, _ = ctxutil.GetPondID(ctx)
	)

	err := input.Validate()
	if err != nil {
		return nil, err
	}

	updatedPond := input.ToPond(userID, pondID)

	err = c.dbTxn.Where("id = ?", pondID).Updates(&updatedPond).Error
	if err != nil {
		return nil, errorpond.ErrFailedUpdatePond.AttacthDetail(map[string]any{"error": err})
	}

	ListPool := model.UpdateListPoolInputToListPool(userID, pondID, input.ListPool)
	ListBerkas := model.UpdateListBerkasInputToListBerkas(userID, pondID, input.ListBerkas)

	err = c.dbTxn.Save(&ListPool).Error
	if err != nil {
		return nil, errorpond.ErrFailedUpdatePond.AttacthDetail(map[string]any{"error": err})
	}
	err = c.dbTxn.Save(&ListBerkas).Error
	if err != nil {
		return nil, errorpond.ErrFailedUpdatePond.AttacthDetail(map[string]any{"error": err})
	}

	return &updatedPond.ID, nil
}

// Commit implements Command.
func (c *command) Commit(ctx context.Context) error {
	if err := orm.CommitTxn(ctx); err != nil {
		return errorpond.ErrCommit.AttacthDetail(map[string]any{"errors": err})
	}
	return nil
}

// Rollback implements Command.
func (c *command) Rollback(ctx context.Context) error {
	if err := orm.RollbackTxn(ctx); err != nil {
		return errorpond.ErrRollback.AttacthDetail(map[string]any{"errors": err})
	}
	return nil
}
