package budidaya

import (
	"context"

	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/common/infra/orm"
	errorbudidaya "github.com/e-fish/api/pkg/domain/budidaya/error-budidaya"
	"github.com/e-fish/api/pkg/domain/budidaya/model"
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

// CreateBudidaya implements Command.
func (c *command) CreateBudidaya(ctx context.Context, input model.CreateBudidayaInput) (*uuid.UUID, error) {
	var (
		userID, _ = ctxutil.GetUserID(ctx)
	)

	err := input.Validate()
	if err != nil {
		return nil, err
	}

	logger.InfoWithContext(ctx, "###find existing budidaya by pool id for validate budidaya not exist")
	exist, err := c.query.ReadBudidayaActiveByPoolID(ctx, input.PoolID)
	if !errorbudidaya.ErrFoundBudidaya.Is(err) {
		return nil, err
	}

	if exist != nil {
		return nil, errorbudidaya.ErrFailedCreateBudidayaExist.AttacthDetail(map[string]any{"pool": exist.PoolID})
	}

	newBudidaya := input.ToBudidaya(userID)

	err = c.dbTxn.Create(&newBudidaya).Error
	if err != nil {
		return nil, errorbudidaya.ErrFailedCreateBudidaya.AttacthDetail(map[string]any{"error": err})
	}

	return &newBudidaya.ID, nil
}

// CreateFishSpecies implements Command.
func (c *command) CreateFishSpecies(ctx context.Context, input model.CreateFishSpeciesInput) (*uuid.UUID, error) {
	var (
		userID, _ = ctxutil.GetUserID(ctx)
	)

	err := input.Validate()
	if err != nil {
		return nil, err
	}

	newFishSpecies := input.ToFishSpecies(userID)

	err = c.dbTxn.Create(&newFishSpecies).Error
	if err != nil {
		return nil, errorbudidaya.ErrFailedCreateBudidaya.AttacthDetail(map[string]any{"error": err})
	}

	return &newFishSpecies.ID, nil
}

// CreateMultiplePricelistBudidaya implements Command.
func (c *command) CreateMultiplePricelistBudidaya(ctx context.Context, input model.CreateMultiplePriceListInput) ([]*uuid.UUID, error) {
	var (
		userID, _ = ctxutil.GetUserID(ctx)
		uid       = []*uuid.UUID{}
	)

	err := input.Validate()
	if err != nil {
		return nil, err
	}

	newPricelist := input.ToMultiplePriceList(userID)

	err = c.dbTxn.Create(&newPricelist).Error
	if err != nil {
		return nil, errorbudidaya.ErrFailedCreateBudidaya.AttacthDetail(map[string]any{"error": err})
	}

	_, err = c.UpdateStatusBudidaya(ctx, model.UpdateBudidayaStatusInput{
		ID:        input.BudidayaID,
		EstTonase: input.EstTonase,
		Status:    model.PANEN,
		EstDate:   input.EstDate,
	})

	if err != nil {
		return nil, err
	}

	for _, v := range newPricelist {
		uid = append(uid, &v.ID)
	}

	return uid, nil
}

// UpdateStatusBudidaya implements Command.
func (c *command) UpdateStatusBudidaya(ctx context.Context, input model.UpdateBudidayaStatusInput) (*uuid.UUID, error) {
	var (
		userID, _ = ctxutil.GetUserID(ctx)
	)

	newStatus := input.ToBudidaya(userID)

	err := c.dbTxn.Updates(&newStatus).Error
	if err != nil {
		return nil, errorbudidaya.ErrFailedUpdateBudidaya.AttacthDetail(map[string]any{"error": err})
	}

	return &newStatus.ID, nil
}

// Commit implements Command.
func (c *command) Commit(ctx context.Context) error {
	if err := orm.CommitTxn(ctx); err != nil {
		return errorbudidaya.ErrCommit.AttacthDetail(map[string]any{"errors": err})
	}
	return nil
}

// Rollback implements Command.
func (c *command) Rollback(ctx context.Context) error {
	if err := orm.RollbackTxn(ctx); err != nil {
		return errorbudidaya.ErrRollback.AttacthDetail(map[string]any{"errors": err})
	}
	return nil
}
