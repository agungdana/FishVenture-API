package transactionservice

import (
	"context"

	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/domain/transaction"
	"github.com/e-fish/api/pkg/domain/transaction/model"
	transactionconfig "github.com/e-fish/api/transaction_http/transaction_config"
	"github.com/google/uuid"
)

func NewService(conf transactionconfig.TransactionConfig) Service {

	repo, err := transaction.NewRepo(conf.DbConfig)
	if err != nil {
		logger.Fatal("###failed create transaction service err: %v", err)
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
