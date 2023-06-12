package authservice

import (
	"context"

	authconfig "github.com/e-fish/api/auth_http/auth_config"
	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/common/infra/firebase"
	"github.com/e-fish/api/pkg/common/infra/token"
	"github.com/e-fish/api/pkg/domain/auth"
	"github.com/e-fish/api/pkg/domain/auth/model"
	"github.com/google/uuid"
)

func NewService(conf authconfig.AuthConfig) Service {
	var (
		ctx = context.Background()
	)

	fb, err := firebase.NewFirebase(conf.FireBaseConfig)
	if err != nil {
		logger.Fatal("###failed create auth service err: %v", err)
	}

	tokenMaker, err := token.NewTokenMaker(token.SecretKey)
	if err != nil {
		logger.Fatal("###failed create auth service err: %v", err)
	}

	authRepo, err := auth.NewRepo(conf.DbConfig, tokenMaker, fb)
	if err != nil {
		logger.Fatal("###failed create auth service err: %v", err)
	}

	service := Service{
		conf: conf,
		repo: authRepo,
	}

	service.RegisterPermissionAccess(ctx)

	return service
}

type Service struct {
	conf authconfig.AuthConfig
	repo auth.Repo
}

func (s *Service) RegisterPermissionAccess(ctx context.Context) error {

	return nil
}

func (s *Service) CreateUser(ctx context.Context, input model.CreateUserInput) (*uuid.UUID, error) {
	command := s.repo.NewCommand(ctx)

	result, err := command.CreateUser(ctx, input)
	if err != nil {
		if err := command.Rollback(ctx); err != nil {
			logger.ErrorWithContext(ctx, "can't rollback transaction err: %v", err)
		}
		logger.ErrorWithContext(ctx, "error create user err: %v", err)
		return nil, err
	}

	if err := command.Commit(ctx); err != nil {
		logger.ErrorWithContext(ctx, "error create user err: %v", err)
		return nil, err
	}

	go s.CreateUserRole(ctx, model.AddUserRoleInput{
		UserID:   *result,
		RoleName: input.ApplicationType,
	})

	return result, nil
}

func (s *Service) CreateUserRole(ctx context.Context, input model.AddUserRoleInput) (*uuid.UUID, error) {
	ctx = ctxutil.NewRequest(ctx)

	command := s.repo.NewCommand(ctx)
	result, err := command.CreateUserRoleByRoleName(ctx, input)
	if err != nil {
		if err := command.Rollback(ctx); err != nil {
			logger.ErrorWithContext(ctx, "can't rollback transaction err: %v", err)
		}
		logger.ErrorWithContext(ctx, "error create user err: %v", err)
		return nil, err
	}

	if err := command.Commit(ctx); err != nil {
		logger.ErrorWithContext(ctx, "error create user err: %v", err)
		return nil, err
	}

	return result, nil
}
