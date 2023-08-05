package pondservice

import (
	"context"
	"mime/multipart"
	"path/filepath"

	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/common/helper/savefile"
	"github.com/e-fish/api/pkg/domain/pond"
	"github.com/e-fish/api/pkg/domain/pond/model"
	"github.com/e-fish/api/pkg/domain/verification"
	pondconfig "github.com/e-fish/api/pond_http/pond_config"
	"github.com/google/uuid"
)

func NewService(conf pondconfig.PondConfig) Service {

	verificationRepo, err := verification.NewRepo(conf.DbConfig)
	if err != nil {
		logger.Fatal("###failed create pond service err: %v", err)
	}

	pondRepo, err := pond.NewRepo(conf.DbConfig, verificationRepo)
	if err != nil {
		logger.Fatal("###failed create pond service err: %v", err)
	}

	service := Service{
		conf: conf,
		repo: pondRepo,
	}

	return service
}

type Service struct {
	conf pondconfig.PondConfig
	repo pond.Repo
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
	return query.GetPondAdmin(ctx)
}

func (s *Service) GetAllPondSubmission(ctx context.Context) ([]*model.PondOutput, error) {
	query := s.repo.NewQuery()
	return query.GetListPondSubmission(ctx)
}

func (s *Service) GetListPond(ctx context.Context) ([]*model.PondOutput, error) {
	query := s.repo.NewQuery()
	return query.GetListPond(ctx)
}

func (s *Service) SaveImagesPond(ctx context.Context, file *multipart.FileHeader) (*UploadPhotoResponse, error) {
	ext := filepath.Ext(file.Filename)
	filename := uuid.New().String() + ext
	err := savefile.SaveFile(file, s.conf.PondImageConfig.Path+"/"+filename)

	if err != nil {
		return nil, err
	}

	result := UploadPhotoResponse{
		Name: filename,
		Url:  s.conf.PondImageConfig.Url + filename,
	}

	return &result, nil
}
func (s *Service) SaveFilePond(ctx context.Context, file *multipart.FileHeader) (*UploadFileResponse, error) {
	ext := filepath.Ext(file.Filename)
	filename := uuid.New().String() + ext
	err := savefile.SaveFile(file, s.conf.PondImageConfig.Path+"/"+filename)

	if err != nil {
		return nil, err
	}

	result := UploadFileResponse{
		Name: filename,
		Url:  s.conf.PondImageConfig.Url + filename,
	}

	return &result, nil
}

func (s *Service) SaveImagesPool(ctx context.Context, file *multipart.FileHeader) (*UploadPhotoResponse, error) {
	ext := filepath.Ext(file.Filename)
	filename := uuid.New().String() + ext
	err := savefile.SaveFile(file, s.conf.PoolImageConfig.Path+"/"+filename)

	if err != nil {
		return nil, err
	}

	result := UploadPhotoResponse{
		Name: filename,
		Url:  s.conf.PoolImageConfig.Url + filename,
	}

	return &result, nil
}
