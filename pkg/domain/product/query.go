package product

import (
	"context"
	"errors"

	errorproduct "github.com/e-fish/api/pkg/domain/product/error-product"
	"github.com/e-fish/api/pkg/domain/product/model"
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

// ReadProductByBudidayaID implements Query.
func (q *query) ReadProductByBudidayaID(ctx context.Context, input uuid.UUID) (*model.ProductOutput, error) {
	var (
		product = model.ProductOutput{}
	)

	err := q.db.Where("deleted_at IS NULL and budidaya_id = ?", input).Take(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorproduct.ErrFoundProduct.AttacthDetail(map[string]any{"budidayaID": input})
		}
		return nil, errorproduct.ErrReadProduct.AttacthDetail(map[string]any{"error": err})
	}
	return &product, nil
}

// ReadProductByBudidayaEstPanenDate implements Query.
func (q *query) ReadProductByBudidayaEstPanenDate(ctx context.Context) ([]*model.ProductOutput, error) {
	panic("unimplemented")
}

// ReadProductByID implements Query.
func (q *query) ReadProductByID(ctx context.Context, input uuid.UUID, withPreload bool) (*model.ProductOutput, error) {
	var (
		data = model.ProductOutput{}
		db   = q.db
	)

	if withPreload {
		db = db.Preload("Budidaya")
	}

	err := db.Where("deleted_at IS NULL and id = ?", input).Take(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorproduct.ErrFoundProduct.AttacthDetail(map[string]any{"id": input})
		}
		return nil, errorproduct.ErrReadProduct.AttacthDetail(map[string]any{"error": err})
	}
	return &data, nil
}

// ReadProductByPondID implements Query.
func (q *query) ReadProductByPondID(ctx context.Context, input uuid.UUID) ([]*model.ProductOutput, error) {
	var (
		data = []*model.ProductOutput{}
		db   = q.db
	)

	err := db.Joins("Budidaya").Find(&data, "Budidaya.pond_id = ?", input).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorproduct.ErrFoundProduct.AttacthDetail(map[string]any{"id": input})
		}
		return nil, errorproduct.ErrReadProduct.AttacthDetail(map[string]any{"error": err})
	}

	return data, nil
}

// lock implements Query.
// lock table row to avoid race condition
func (q *query) lock() Query {
	db := q.db.Clauses(clause.Locking{Strength: "UPDATE"})
	return &query{db: db}
}
