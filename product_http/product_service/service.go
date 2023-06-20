package productservice

import (
	"context"

	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/domain/product"
	"github.com/e-fish/api/pkg/domain/product/model"
	productconfig "github.com/e-fish/api/product_http/product_config"
	"github.com/google/uuid"
)

func NewService(conf productconfig.ProductConfig) Service {

	repo, err := product.NewRepo(conf.DbConfig)
	if err != nil {
		logger.Fatal("###failed create budidaya service err: %v", err)
	}

	service := Service{
		conf: conf,
		repo: repo,
	}

	return service
}

type Service struct {
	conf productconfig.ProductConfig
	repo product.Repo
}

func (s *Service) CreateProduct(ctx context.Context, input model.CreateProductInput) (*uuid.UUID, error) {
	command := s.repo.NewCommand(ctx)

	result, err := command.CreateProduct(ctx, input)
	if err != nil {
		if err := command.Rollback(ctx); err != nil {
			logger.ErrorWithContext(ctx, "failed rollback transaction create pond err: %v", err)
		}
		logger.ErrorWithContext(ctx, "failed create pond err: %v", err)
		return nil, err
	}

	if err := command.Commit(ctx); err != nil {
		logger.ErrorWithContext(ctx, "failed commit transaction create pond err: %v", err)
		return nil, err
	}

	return result, nil
}

func (s *Service) DeleteProduct(ctx context.Context, input uuid.UUID) (*uuid.UUID, error) {
	command := s.repo.NewCommand(ctx)

	result, err := command.DeleteProduct(ctx, input)
	if err != nil {
		if err := command.Rollback(ctx); err != nil {
			logger.ErrorWithContext(ctx, "failed rollback transaction create pond err: %v", err)
		}
		logger.ErrorWithContext(ctx, "failed create pond err: %v", err)
		return nil, err
	}

	if err := command.Commit(ctx); err != nil {
		logger.ErrorWithContext(ctx, "failed commit transaction create pond err: %v", err)
		return nil, err
	}

	return result, nil
}
