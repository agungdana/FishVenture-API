package budidaya

import (
	"context"

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
func (q *query) ReadBudidayaByUserLogin(ctx context.Context) ([]*model.BudidayaOutput, error) {
	var (
		appType, _ = ctxutil.GetUserAppType(ctx)
	)
	switch appType {
	case usertype.BUYER:
		return q.ReadBudidayaByUserBuyer(ctx)
	case usertype.ADMIN:
		return q.ReadBudidayaByUserBuyer(ctx)
	case usertype.SALLER:
		return q.ReadBudidayaByUserBuyer(ctx)
	default:
		return nil, errorbudidaya.ErrUnsuportedFindBudidaya.AttacthDetail(map[string]any{"type": appType})
	}
}

// ReadBudidayaByUserAdmin implements Query.
func (q *query) ReadBudidayaByUserAdmin(ctx context.Context) ([]*model.BudidayaOutput, error) {
	panic("unimplemented")
}

// ReadBudidayaByUserBuyer implements Query.
func (q *query) ReadBudidayaByUserBuyer(ctx context.Context) ([]*model.BudidayaOutput, error) {
	panic("unimplemented")
}

// ReadBudidayaByUserSaller implements Query.
func (q *query) ReadBudidayaByUserSaller(ctx context.Context) ([]*model.BudidayaOutput, error) {
	panic("unimplemented")
}

// ReadBudidayaActiveByPoolID implements Query.
func (q *query) ReadBudidayaActiveByPoolID(ctx context.Context, input uuid.UUID) (*model.BudidayaOutput, error) {
	panic("unimplemented")
}
