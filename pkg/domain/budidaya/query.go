package budidaya

import (
	"context"
	"errors"

	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	usertype "github.com/e-fish/api/pkg/domain/auth/model"
	errorbudidaya "github.com/e-fish/api/pkg/domain/budidaya/error-budidaya"
	"github.com/e-fish/api/pkg/domain/budidaya/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func newQuery(db *gorm.DB) Query {
	return &query{
		db: db,
	}
}

// lock implements Query.
// lock table row to avoid race condition
func (q *query) lock() Query {
	db := q.db.Clauses(clause.Locking{Strength: "UPDATE"})
	return &query{db: db}
}

type query struct {
	db *gorm.DB
}

// ReadBudidayaByUserLogin implements Query.
func (q *query) ReadBudidayaByUserLogin(ctx context.Context, input model.GetBudidayaInput) ([]*model.BudidayaOutput, error) {
	var (
		appType, _ = ctxutil.GetUserAppType(ctx)
	)
	switch appType {
	case usertype.BUYER:
		return q.ReadBudidayaByUserBuyer(ctx, input)
	case usertype.ADMIN:
		return q.ReadBudidayaByUserAdmin(ctx, input)
	default:
		return nil, errorbudidaya.ErrUnsuportedFindBudidaya.AttacthDetail(map[string]any{"type": appType})
	}
}

// ReadBudidayaByUserAdmin implements Query.
func (q *query) ReadBudidayaByUserAdmin(ctx context.Context, input model.GetBudidayaInput) ([]*model.BudidayaOutput, error) {
	var (
		res = []*model.BudidayaOutput{}
		db  = q.db
	)

	db = db.Preload("Pool")
	db = db.Preload("FishSpecies")
	err := db.Where("deleted_at IS NULL and status <> ? and pond_id = ?", model.END, input.PondID).Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

// ReadBudidayaByUserBuyer implements Query.
func (q *query) ReadBudidayaByUserBuyer(ctx context.Context, input model.GetBudidayaInput) ([]*model.BudidayaOutput, error) {
	var (
		res = []*model.BudidayaOutput{}
		db  = q.db
	)

	db = db.Preload("Pool")
	db = db.Preload("FishSpecies")
	err := db.Where("deleted_at IS NULL and status <> ? and pond_id = ?", model.END, input.PondID).Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

// ReadBudidayaByUserSaller implements Query.
func (q *query) ReadBudidayaByUserSaller(ctx context.Context) ([]*model.BudidayaOutput, error) {
	var (
		pondID, _ = ctxutil.GetPondID(ctx)
		res       = []*model.BudidayaOutput{}
		db        = q.db
	)
	db = db.Preload("Pool")
	db = db.Preload("FishSpecies")
	err := db.Where("deleted_at IS NULL and pond_id = ? and status <> ?", pondID, model.END).Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

// ReadBudidayaActiveByPoolID implements Query.
func (q *query) ReadBudidayaActiveByPoolID(ctx context.Context, input uuid.UUID) (*model.BudidayaOutput, error) {
	var (
		res = model.BudidayaOutput{}
	)

	err := q.db.Where("deleted_at IS NULL and pool_id and status <> ?", input, model.END).Take(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorbudidaya.ErrFoundBudidaya.AttacthDetail(map[string]any{"poolID": input})
		}
		return nil, errorbudidaya.ErrFailedReadBudidaya.AttacthDetail(map[string]any{"error": err})
	}

	return &res, nil
}
