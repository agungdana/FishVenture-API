package product

import (
	"context"

	"github.com/e-fish/api/pkg/domain/product/model"
	"github.com/google/uuid"
)

type Repo interface {
	NewCommand(ctx context.Context) Command
	NewQuery() Query
}

type Command interface {
	CreateProduct(ctx context.Context, input model.CreateProductInput) (*uuid.UUID, error)
	DeleteProduct(ctx context.Context, input uuid.UUID) (*uuid.UUID, error)

	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
}

type Query interface {
	ReadProductByBudidayaID(ctx context.Context, input uuid.UUID) (*model.ProductOutput, error)
	ReadProductByPondID(ctx context.Context, input uuid.UUID) ([]*model.ProductOutput, error)
	ReadProductByID(ctx context.Context, input uuid.UUID) (*model.ProductOutput, error)
	ReadProductByBudidayaEstPanenDate(ctx context.Context) ([]*model.ProductOutput, error)
	lock() Query
}
