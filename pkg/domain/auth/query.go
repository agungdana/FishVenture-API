package auth

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type query struct {
	db *gorm.DB
}

func newQuery(db *gorm.DB) Query {
	return &query{db: db}
}

// lock implements Query.
// lock table row to avoid race condition
func (q *query) lock() Query {
	db := q.db.Clauses(clause.Locking{Strength: "UPDATE"})
	return &query{db: db}
}
