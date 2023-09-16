package budidayaservice

import (
	"context"

	budidayaconfig "github.com/e-fish/api/budidaya_http/budidaya_config"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/domain/budidaya"
	"github.com/e-fish/api/pkg/domain/budidaya/model"
	"github.com/e-fish/api/pkg/domain/pond"
	"github.com/e-fish/api/pkg/domain/verification"
	"github.com/google/uuid"
)

func NewService(conf budidayaconfig.BudidayaConfig) Service {

	verificationRepo, err := verification.NewRepo(conf.DbConfig)
	if err != nil {
		logger.Fatal("###failed create budidaya service err: %v", err)
	}

	pondRepo, err := pond.NewRepo(conf.DbConfig, verificationRepo)
	if err != nil {
		logger.Fatal("###failed create budidaya service err: %v", err)
	}

	buidayaRepo, err := budidaya.NewRepo(conf.DbConfig, pondRepo)
	if err != nil {
		logger.Fatal("###failed create budidaya service err: %v", err)
	}

	service := Service{
		conf: conf,
		repo: buidayaRepo,
	}

	return service
}

type Service struct {
	conf budidayaconfig.BudidayaConfig
	repo budidaya.Repo
}

func (s *Service) CreateBudidaya(ctx context.Context, input model.CreateBudidayaInput) (*uuid.UUID, error) {
	command := s.repo.NewCommand(ctx)

	result, err := command.CreateBudidaya(ctx, input)
	if err != nil {
		if err := command.Rollback(ctx); err != nil {
			logger.ErrorWithContext(ctx, "failed rollback transaction create budidaya err: %v", err)
		}
		logger.ErrorWithContext(ctx, "failed create budidaya err: %v", err)
		return nil, err
	}

	if err := command.Commit(ctx); err != nil {
		logger.ErrorWithContext(ctx, "failed commit transaction create budidaya err: %v", err)
		return nil, err
	}

	return result, nil
}

func (s *Service) UpdateBudidayaWithPricelist(ctx context.Context, input model.UpdateBudidayaWithPricelist) (*uuid.UUID, error) {
	command := s.repo.NewCommand(ctx)

	result, err := command.UpdateStatusBudidayaWithListPricelist(ctx, input)
	if err != nil {
		if err := command.Rollback(ctx); err != nil {
			logger.ErrorWithContext(ctx, "failed rollback transaction update budidaya with pricelist err: %v", err)
		}
		logger.ErrorWithContext(ctx, "failed update budidaya with pricelist err: %v", err)
		return nil, err
	}

	if err := command.Commit(ctx); err != nil {
		logger.ErrorWithContext(ctx, "failed commit transaction update budidaya with pricelist err: %v", err)
		return nil, err
	}

	return result, nil
}

func (s *Service) CreateFishSpecies(ctx context.Context, input model.CreateFishSpeciesInput) (*uuid.UUID, error) {
	command := s.repo.NewCommand(ctx)

	result, err := command.CreateFishSpecies(ctx, input)
	if err != nil {
		if err := command.Rollback(ctx); err != nil {
			logger.ErrorWithContext(ctx, "failed rollback transaction create fish species err: %v", err)
		}
		logger.ErrorWithContext(ctx, "failed create fish species err: %v", err)
		return nil, err
	}

	if err := command.Commit(ctx); err != nil {
		logger.ErrorWithContext(ctx, "failed commit transaction create fish species err: %v", err)
		return nil, err
	}

	return result, nil
}

func (s *Service) CreateMultiplePricelist(ctx context.Context, input model.CreateMultiplePriceListInput) ([]*uuid.UUID, error) {
	command := s.repo.NewCommand(ctx)

	result, err := command.CreateMultiplePricelistBudidaya(ctx, input)
	if err != nil {
		if err := command.Rollback(ctx); err != nil {
			logger.ErrorWithContext(ctx, "failed rollback transaction create fish species err: %v", err)
		}
		logger.ErrorWithContext(ctx, "failed create fish species err: %v", err)
		return nil, err
	}

	if err := command.Commit(ctx); err != nil {
		logger.ErrorWithContext(ctx, "failed commit transaction create fish species err: %v", err)
		return nil, err
	}

	return result, nil
}

func (s *Service) GetBudidayaByUserLoginAdminOrCustomer(ctx context.Context, input model.GetBudidayaInput) ([]*model.BudidayaOutput, error) {
	query := s.repo.NewQuery()
	return query.ReadBudidayaByUserLogin(ctx, input)
}

func (s *Service) GetBudidayaByUserSeller(ctx context.Context) ([]*model.BudidayaOutput, error) {
	query := s.repo.NewQuery()
	return query.ReadBudidayaByUserSeller(ctx)
}

func (s *Service) GetAllFishSpecies(ctx context.Context) ([]*model.FishSpeciesOutput, error) {
	query := s.repo.NewQuery()
	return query.ReadAllDataFishSpecies(ctx)
}

func (s *Service) ReadBudidayaNeaerest(ctx context.Context) ([]*model.BudidayaOutput, error) {
	query := s.repo.NewQuery()
	return query.ReadBudidayaNeaerest(ctx)
}
