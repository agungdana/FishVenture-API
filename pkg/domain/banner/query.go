package banner

import (
	"context"

	"github.com/e-fish/api/pkg/domain/banner/model"
	errorBanner "github.com/e-fish/api/pkg/domain/banner/model/error-banner"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func newQuery(db *gorm.DB) Query {
	return &query{db: db}
}

// lock implements Query.
func (q *query) lock() Query {
	db := q.db.Clauses(clause.Locking{Strength: "UPDATE"})
	return &query{db: db}
}

type query struct {
	db *gorm.DB
}

// ReadAllBanner implements Query
func (q *query) ReadAllBanner(ctx context.Context) ([]model.BannerOutput, error) {
	data := []model.BannerOutput{}

	err := q.db.Where("deleted_at IS NULL").Find(&data).Error
	if err != nil {

		return nil, errorBanner.ErrFindBanner.AttacthDetail(map[string]any{"error find banner": err})
	}
	return data, nil
}
