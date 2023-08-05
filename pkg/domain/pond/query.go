package pond

import (
	"context"
	"errors"

	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	userModel "github.com/e-fish/api/pkg/domain/auth/model"
	errorpond "github.com/e-fish/api/pkg/domain/pond/error-pond"
	"github.com/e-fish/api/pkg/domain/pond/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func newQuery(db *gorm.DB) Query {
	return &query{
		db: db,
	}
}

type query struct {
	db *gorm.DB
}

// GetPondByID implements Query.
func (q *query) GetPondByID(ctx context.Context, input uuid.UUID) (*model.PondOutput, error) {
	var (
		data = model.PondOutput{}
	)

	err := q.db.Where("deleted_at IS NULL and id = ?", input).Take(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorpond.ErrFoundPond
		}
		return nil, errorpond.ErrFailedFindPond.AttacthDetail(map[string]any{"error": err})
	}

	return &data, nil
}

// GetListPondSubmission implements Query.
func (q *query) GetListPondSubmission(ctx context.Context) ([]*model.PondOutput, error) {
	var (
		data = []*model.PondOutput{}
	)

	err := q.db.Where("deleted_at IS NULL").Preload("Team").Preload("ListPool").Preload("ListBerkas").Find(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorpond.ErrFoundPond
		}
		return nil, errorpond.ErrFailedFindPond.AttacthDetail(map[string]any{"error": err})
	}

	if len(data) < 1 {
		return nil, errorpond.ErrFoundPond
	}

	return data, nil
}

// GetPondAdmin implements Query.
func (q *query) GetPondAdmin(ctx context.Context) (*model.PondOutput, error) {
	var (
		userID, _ = ctxutil.GetUserID(ctx)
		pondID, _ = ctxutil.GetPondID(ctx)
		data      = model.PondOutput{}
	)

	err := q.db.Where("deleted_at IS NULL and id = ? and user_id = ?", pondID, userID).Take(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorpond.ErrFoundPond
		}
		return nil, errorpond.ErrFailedFindPond.AttacthDetail(map[string]any{"error": err})
	}

	return &data, nil
}

// GetListPond implements Query.
func (q *query) GetListPond(ctx context.Context) ([]*model.PondOutput, error) {
	var (
		data       = []*model.PondOutput{}
		appType, _ = ctxutil.GetUserAppType(ctx)
		db         = q.db
	)

	if userModel.BUYER == appType {
		db = db.Where("status = ?", model.ACTIVED)
	}

	if userModel.ADMIN == appType {
		db = db.Preload("Team").Preload("ListBerkas").Preload("ListPool").Order("CASE WHEN status = " + model.SUBMISION + " THEN 1 ELSE 2 END, status")
	}

	err := db.Where("deleted_at IS NULL").Find(&data).Error
	if err != nil {
		return nil, errorpond.ErrFailedFindPond.AttacthDetail(map[string]any{"error": err})
	}

	return data, nil
}

// lock implements Query.
// lock table row to avoid race condition
func (q *query) lock() Query {
	db := q.db.Clauses(clause.Locking{Strength: "UPDATE"})
	return &query{db: db}
}
