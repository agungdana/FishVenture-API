package authservice

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"

	authconfig "github.com/e-fish/api/auth_http/auth_config"
	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/common/helper/savefile"
	"github.com/e-fish/api/pkg/common/helper/werror"
	"github.com/e-fish/api/pkg/common/infra/firebase"
	"github.com/e-fish/api/pkg/common/infra/token"
	"github.com/e-fish/api/pkg/domain/auth"
	errorauth "github.com/e-fish/api/pkg/domain/auth/error"
	"github.com/e-fish/api/pkg/domain/auth/model"
	"github.com/google/uuid"
)

func NewService(conf authconfig.AuthConfig) Service {
	var (
		ctx = context.Background()
	)

	fb, err := firebase.NewFirebase(conf.FireBaseConfig)
	if err != nil {
		logger.Fatal("###failed create firebase service err: %v", err)
	}

	tokenMaker, err := token.NewTokenMaker(token.SecretKey)
	if err != nil {
		logger.Fatal("###failed create token maker service err: %v", err)
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
	query := s.repo.NewQuery()
	roles, err := query.GetAllRolePermission(ctx)
	if err != nil {
		logger.Fatal("err get all permission: %v", err)
	}
	user, err := query.GetAllUserPermission(ctx)
	if err != nil {
		if !errorauth.ErrGetUserPermissionEmpty.Is(err) {
			logger.Fatal("err get all user permission")
		}
	}
	access := []ctxutil.PermissionAccess{}
	for _, v := range roles {
		access = append(access, ctxutil.PermissionAccess{
			ID:   v.RoleID,
			Path: v.PermissionPath,
		})
	}
	for _, v := range user {
		access = append(access, ctxutil.PermissionAccess{
			ID:   v.RoleID,
			Path: v.PermissionPath,
		})
	}
	ctxutil.AddPermissionAccess(access)
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
	// ctx = ctxutil.NewRequest(ctx)
	newCtx := ctxutil.NewRequestWithOutTimeOut(ctx)

	command := s.repo.NewCommand(newCtx)
	result, err := command.CreateUserRoleByRoleName(newCtx, input)
	if err != nil {
		if err := command.Rollback(newCtx); err != nil {
			logger.ErrorWithContext(newCtx, "can't rollback transaction err: %v", err)
		}
		logger.ErrorWithContext(newCtx, "error create user err: %v", err)
		return nil, err
	}

	if err := command.Commit(newCtx); err != nil {
		logger.ErrorWithContext(newCtx, "error create user err: %v", err)
		return nil, err
	}

	return result, nil
}

func (s *Service) Login(ctx context.Context, input model.UserLoginInput) (*model.UserLoginOutput, error) {
	command := s.repo.NewCommand(ctx)
	result, err := command.Login(ctx, input)
	if err != nil {
		if err := command.Rollback(ctx); err != nil {
			logger.ErrorWithContext(ctx, "can't rollback transaction err: %v", err)
		}
		logger.ErrorWithContext(ctx, "error login user err: %v", err)
		return nil, err
	}

	if err := command.Commit(ctx); err != nil {
		logger.ErrorWithContext(ctx, "error login user err: %v", err)
		return nil, err
	}

	return result, nil
}

func (s *Service) LoginByGoogle(ctx context.Context, input model.UserLoginByGooleInput) (*model.UserLoginOutput, error) {
	command := s.repo.NewCommand(ctx)
	result, err := command.LoginByGoogle(ctx, input)
	if err != nil {
		if err := command.Rollback(ctx); err != nil {
			logger.ErrorWithContext(ctx, "can't rollback transaction err: %v", err)
		}
		logger.ErrorWithContext(ctx, "error login user err: %v", err)
		return nil, err
	}

	if err := command.Commit(ctx); err != nil {
		logger.ErrorWithContext(ctx, "error login user err: %v", err)
		return nil, err
	}

	return result, nil
}

func (s *Service) Profile(ctx context.Context) (*model.Profile, error) {
	query := s.repo.NewQuery()
	result, err := query.GetProfile(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Service) SaveImages(ctx context.Context, file *multipart.FileHeader) (*UploadPhotoResponse, error) {
	ext := filepath.Ext(file.Filename)
	filename := uuid.New().String() + ext

	_, imageExtOk := savefile.ImageExt[ext]

	if !imageExtOk {
		return nil, werror.Error{
			Code:    "FailedSaveFile",
			Message: "extension not suported",
			Details: map[string]any{
				"ext":                  ext,
				"permitted-extensions": fmt.Sprintf("%v", savefile.ImageExt),
			},
		}
	}

	err := savefile.SaveFile(file, s.conf.UserImageConfig.Path+"/"+filename)

	if err != nil {
		return nil, err
	}

	result := UploadPhotoResponse{
		Name: filename,
		Url:  s.conf.UserImageConfig.Url + filename,
	}

	return &result, nil
}

func (s *Service) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*uuid.UUID, error) {
	command := s.repo.NewCommand(ctx)

	result, err := command.UpdateUser(ctx, input)
	if err != nil {
		if err := command.Rollback(ctx); err != nil {
			logger.ErrorWithContext(ctx, "can't rollback transaction err: %v", err)
		}
		logger.ErrorWithContext(ctx, "error update user err: %v", err)
		return nil, err
	}

	if err := command.Commit(ctx); err != nil {
		logger.ErrorWithContext(ctx, "error update user err: %v", err)
		return nil, err
	}

	return result, nil
}
