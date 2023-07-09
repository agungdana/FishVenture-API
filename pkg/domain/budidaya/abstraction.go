package budidaya

import (
	"context"

	"github.com/e-fish/api/pkg/domain/budidaya/model"
	"github.com/google/uuid"
)

type Repo interface {
	NewCommand(ctx context.Context) Command
	NewQuery() Query
}

type Command interface {
	CreateBudidaya(ctx context.Context, input model.CreateBudidayaInput) (*uuid.UUID, error)
	UpdateStatusBudidaya(ctx context.Context, input model.UpdateBudidayaStatusInput) (*uuid.UUID, error)

	CreateMultiplePricelistBudidaya(ctx context.Context, input model.CreateMultiplePriceListInput) ([]*uuid.UUID, error)

	CreateFishSpecies(ctx context.Context, input model.CreateFishSpeciesInput) (*uuid.UUID, error)

	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
}

type Query interface {
	ReadBudidayaActiveByPoolID(ctx context.Context, input uuid.UUID) (*model.BudidayaOutput, error)
	ReadBudidayaByUserLogin(ctx context.Context, input model.GetBudidayaInput) ([]*model.BudidayaOutput, error)
	ReadBudidayaByUserBuyer(ctx context.Context, input model.GetBudidayaInput) ([]*model.BudidayaOutput, error)
	ReadBudidayaByUserAdmin(ctx context.Context, input model.GetBudidayaInput) ([]*model.BudidayaOutput, error)
	ReadBudidayaByUserSaller(ctx context.Context) ([]*model.BudidayaOutput, error)

	lock() Query
}
