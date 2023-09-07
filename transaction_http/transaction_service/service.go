package transactionservice

import (
	"context"

	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/domain/budidaya"
	"github.com/e-fish/api/pkg/domain/pond"
	"github.com/e-fish/api/pkg/domain/transaction"
	"github.com/e-fish/api/pkg/domain/transaction/model"
	"github.com/e-fish/api/pkg/domain/verification"
	transactionconfig "github.com/e-fish/api/transaction_http/transaction_config"
	"github.com/google/uuid"
)

func NewService(conf transactionconfig.TransactionConfig) Service {

	verificationRepo, err := verification.NewRepo(conf.DbConfig)
	if err != nil {
		logger.Fatal("###failed create transaction service [causes: %v, err: %v]", "verification.NewRepo", err)
	}

	pondRepo, err := pond.NewRepo(conf.DbConfig, verificationRepo)
	if err != nil {
		logger.Fatal("###failed create transaction service [causes: %v, err: %v]", "pond.NewRepo", err)
	}

	budidayaRepo, err := budidaya.NewRepo(conf.DbConfig, pondRepo)
	if err != nil {
		logger.Fatal("###failed create transaction service [causes: %v, err: %v]", "budidaya.NewRepo", err)
	}

	repo, err := transaction.NewRepo(conf.DbConfig, budidayaRepo)
	if err != nil {
		logger.Fatal("###failed create transaction service [causes: %v, err: %v]", "transaction.NewRepo", err)
	}

	service := Service{
		conf: conf,
		repo: repo,
	}

	return service
}

type Service struct {
	conf transactionconfig.TransactionConfig
	repo transaction.Repo
}

func (s *Service) CreateOrderInput(ctx context.Context, input model.CreateOrderInput) (*uuid.UUID, error) {
	command := s.repo.NewCommand(ctx)
	result, err := command.CreateOrder(ctx, input)
	if err != nil {
		if err := command.Rollback(ctx); err != nil {
			logger.ErrorWithContext(ctx, "failed rollback transaction err: %v", err)
		}
		return nil, err
	}

	if err := command.Commit(ctx); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Service) ReadOrder(ctx context.Context, input model.ReadInput) (*model.OrderOutputPagination, error) {
	command := s.repo.NewQuery()
	result, err := command.ReadOrder(ctx, input)
	return result, err
}

func (s *Service) ReadOrderCancel(ctx context.Context, input model.ReadInput) (*model.OrderOutputPagination, error) {
	command := s.repo.NewQuery()
	result, err := command.ReadOrderByStatus(ctx, input, model.CANCEL)
	return result, err
}

func (s *Service) ReadOrderSuccess(ctx context.Context, input model.ReadInput) (*model.OrderOutputPagination, error) {
	command := s.repo.NewQuery()
	result, err := command.ReadOrderByStatus(ctx, input, model.SUCCESS)
	return result, err
}

func (s *Service) UpdateOrderCancel(ctx context.Context, input uuid.UUID) (*uuid.UUID, error) {
	command := s.repo.NewCommand(ctx)
	result, err := command.UpdateCancelOrder(ctx, input)
	if err != nil {
		if err := command.Rollback(ctx); err != nil {
			logger.ErrorWithContext(ctx, "failed rollback transaction err: %v", err)
		}
		logger.ErrorWithContext(ctx, "failed update order err: %v", err)
		return nil, err
	}

	if err := command.Commit(ctx); err != nil {
		logger.ErrorWithContext(ctx, "failed commit transaction err: %v", err)
		return nil, err
	}

	return result, nil
}

func (s *Service) UpdateOrderSuccess(ctx context.Context, input uuid.UUID) (*uuid.UUID, error) {
	command := s.repo.NewCommand(ctx)
	result, err := command.UpdateSuccesOrder(ctx, input)
	if err != nil {
		if err := command.Rollback(ctx); err != nil {
			logger.ErrorWithContext(ctx, "failed rollback transaction err: %v", err)
		}
		logger.ErrorWithContext(ctx, "failed update order err: %v", err)
		return nil, err
	}

	if err := command.Commit(ctx); err != nil {
		logger.ErrorWithContext(ctx, "failed commit transaction err: %v", err)
		return nil, err
	}

	return result, nil
}
