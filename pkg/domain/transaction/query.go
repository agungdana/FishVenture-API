package transaction

import (
	"context"
	"errors"

	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/e-fish/api/pkg/common/infra/orm"
	userModel "github.com/e-fish/api/pkg/domain/auth/model"
	errortransaction "github.com/e-fish/api/pkg/domain/transaction/error-transaction"
	"github.com/e-fish/api/pkg/domain/transaction/model"
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

// ReadAllOrderActive implements Query.
func (q *query) ReadAllOrderActive(ctx context.Context) ([]*model.Order, error) {
	data := []*model.Order{}
	err := q.db.Where("deleted_at IS NULL and status = ?", model.ACTIVE).Find(&data).Error
	if err != nil {
		return nil, errortransaction.ErrFoundOrder.AttacthDetail(map[string]any{"error": err})
	}
	return data, nil
}

// ReadOrder implements Query.
func (q *query) ReadOrder(ctx context.Context, input model.ReadInput) (*model.OrderOutputPagination, error) {
	var (
		order      []*model.OrderOutput
		pagination model.OrderOutputPagination
		db         = q.db
		userID, _  = ctxutil.GetUserID(ctx)
		pondID, _  = ctxutil.GetPondID(ctx)
		appType, _ = ctxutil.GetUserAppType(ctx)
	)

	input.ObjectTable = model.Order{}

	db = db.Where("deleted_at is NULL")

	switch appType {
	case userModel.BUYER:
		db = db.Where("user_id = ?", userID).Preload("Budidaya.Pond").Preload("Budidaya.Pool").Preload("Budidaya.FishSpecies")
	case userModel.SELLER:
		db = db.Where("pond_id = ?", pondID).Preload("Budidaya.Pool").Preload("User")
	}

	err := db.Scopes(orm.Paginate(db, &input.Paginantion)).Find(&order).Error
	if err != nil {
		return nil, errortransaction.ErrReadOrderData.AttacthDetail(map[string]any{"error": err})
	}

	if len(order) < 1 {
		return nil, errortransaction.ErrFoundOrder
	}

	pagination = model.OrderOutputPagination{
		FindBy:    input.FindBy,
		Keyword:   input.Keyword,
		Limit:     input.Limit,
		Page:      input.Page,
		Sort:      input.Sort,
		Direction: input.Direction,
		TotalRows: input.TotalRows,
		TotalPage: input.TotalPage,
		Rows:      order,
	}

	return &pagination, nil
}

// ReadOrder implements Query.
func (q *query) ReadOrderByStatus(ctx context.Context, input model.ReadInput, status string) (*model.OrderOutputPagination, error) {
	var (
		order      []*model.OrderOutput
		pagination model.OrderOutputPagination
		db         = q.db
		userID, _  = ctxutil.GetUserID(ctx)
		pondID, _  = ctxutil.GetPondID(ctx)
		appType, _ = ctxutil.GetUserAppType(ctx)
	)

	input.ObjectTable = model.Order{}

	db = db.Where("deleted_at is NULL and status = ?", status)

	switch appType {
	case userModel.BUYER:
		db = db.Where("user_id = ?", userID).Preload("Budidaya.Pond").Preload("Budidaya.Pool").Preload("Budidaya.FishSpecies")
	case userModel.SELLER:
		db = db.Where("pond_id = ?", pondID).Preload("Budidaya.Pool").Preload("User")
	}

	err := db.Scopes(orm.Paginate(db, &input.Paginantion)).Find(&order).Error
	if err != nil {
		return nil, errortransaction.ErrReadOrderData.AttacthDetail(map[string]any{"error": err})
	}

	if len(order) < 1 {
		return nil, errortransaction.ErrFoundOrder
	}

	pagination = model.OrderOutputPagination{
		FindBy:    input.FindBy,
		Keyword:   input.Keyword,
		Limit:     input.Limit,
		Page:      input.Page,
		Sort:      input.Sort,
		Direction: input.Direction,
		TotalRows: input.TotalRows,
		TotalPage: input.TotalPage,
		Rows:      order,
	}

	return &pagination, nil
}

// ReadOrderByID implements Query.
func (q *query) ReadOrderByID(ctx context.Context, id uuid.UUID) (*model.OrderOutput, error) {
	var order model.OrderOutput

	err := q.db.Where("deleted_at is NULL and id = ?", id).Take(&order).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errortransaction.ErrFoundOrder.AttacthDetail(map[string]any{"error": err, "id": id})
		}
		return nil, errortransaction.ErrReadOrderData.AttacthDetail(map[string]any{"error": err, "id": id})
	}

	return &order, nil
}
