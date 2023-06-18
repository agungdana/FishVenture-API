package product

import (
	"context"
)

type Repo interface {
	NewCommand(ctx context.Context) Command
	NewQuery() Query
}

type Command interface {
	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
}

type Query interface {
	lock() Query
}
