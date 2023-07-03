package budidayaservice

import (
	"context"

	budidayaconfig "github.com/e-fish/api/budidaya_http/budidaya_config"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/domain/budidaya"
	"github.com/e-fish/api/pkg/domain/budidaya/model"
	"github.com/e-fish/api/pkg/domain/verification"
	"github.com/google/uuid"
)

func NewService(conf budidayaconfig.BudidayaConfig) Service {

	verificationRepo, err := verification.NewRepo(conf.DbConfig)
	if err != nil {
		logger.Fatal("###failed create budidaya service err: %v", err)
	}

	buidayaRepo, err := budidaya.NewRepo(conf.DbConfig, verificationRepo)
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

func (s *Service) CreatePond(ctx context.Context, input model.CreatePondInput) (*uuid.UUID, error) {
	command := s.repo.NewCommand(ctx)

	result, err := command.CreatePond(ctx, input)
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

func (s *Service) UpdatePond(ctx context.Context, input model.UpdatePondInput) (*uuid.UUID, error) {
	command := s.repo.NewCommand(ctx)

	result, err := command.UpdatePond(ctx, input)
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

func (s *Service) UpdatePondStatus(ctx context.Context, input model.UpdatePondStatus) (*uuid.UUID, error) {
	command := s.repo.NewCommand(ctx)

	result, err := command.UpdatePondStatus(ctx, input)
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

func (s *Service) GetPondByUserID(ctx context.Context) (*model.PondOutput, error) {
	query := s.repo.NewQuery()
	return query.GetPondByUserPondAdmin(ctx)
}

func (s *Service) GetAllPondSubmission(ctx context.Context) ([]*model.PondOutput, error) {
	query := s.repo.NewQuery()
	return query.GetListPondSubmission(ctx)
}

func (s *Service) GetListPondForUser(ctx context.Context) ([]*model.PondOutput, error) {
	query := s.repo.NewQuery()
	return query.GetListPondForUser(ctx)
}
