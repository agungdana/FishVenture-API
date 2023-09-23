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
	UpdateStatusBudidayaWithListPricelist(ctx context.Context, input model.UpdateBudidayaWithPricelist) (*uuid.UUID, error)
	UpdateBudidayaSoldQty(ctx context.Context, input model.UpdateBudidayaSoldQty) (*uuid.UUID, error)
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
	ReadBudidayaByUserSeller(ctx context.Context) ([]*model.BudidayaOutput, error)
	ReadBudidayaCodeActive(ctx context.Context) (*string, error)
	ReadBudidayaNeaerest(ctx context.Context) ([]*model.BudidayaOutput, error)

	ReadBudidayaByID(ctx context.Context, id uuid.UUID) (*model.BudidayaOutput, error)

	ReadAllDataFishSpecies(ctx context.Context) ([]*model.FishSpeciesOutput, error)

	ReadPriceListBudidayaByBiggerThanLimitAndBudidayaID(ctx context.Context, input model.ReadPricelistBudidayaInput) (*model.PriceList, error)
	ReadPriceListBudidayaBySmallerThanLimitAndBudidayaID(ctx context.Context, input model.ReadPricelistBudidayaInput) (*model.PriceList, error)

	lock() Query
}
