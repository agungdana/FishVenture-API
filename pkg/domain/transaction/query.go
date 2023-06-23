package transaction

import (
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

// lock implements Query.
// lock table row to avoid race condition
func (q *query) lock() Query {
	db := q.db.Clauses(clause.Locking{Strength: "UPDATE"})
	return &query{db: db}
}
