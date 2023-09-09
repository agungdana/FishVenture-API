package chatservice

import (
	"context"
	"mime/multipart"
	"path/filepath"

	chatconfig "github.com/e-fish/api/chat_http/chat_config"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/common/helper/savefile"
	"github.com/e-fish/api/pkg/domain/chat"
	"github.com/e-fish/api/pkg/domain/chat/model"
	"github.com/google/uuid"
)

type Service struct {
	conf chatconfig.ChatConfig
	repo chat.Repo
}

func NewService(conf chatconfig.ChatConfig) Service {
	var (
		service = Service{
			conf: conf,
		}
	)

	chatRepo, err := chat.NewRepo(conf.ChatDBConfig)
	if err != nil {
		logger.Fatal("failed to create a new repo, can't create region service err causes failed create chat repo: %v", err)
	}

	service.repo = chatRepo

	return service
}

func (s *Service) CreateChat(ctx context.Context, input model.CreateChatInput) (*uuid.UUID, error) {
	command := s.repo.NewCommand(ctx)

	result, err := command.CreateChat(ctx, input)
	if err != nil {
		if err := command.Rollback(ctx); err != nil {
			logger.ErrorWithContext(ctx, "failed rollback transaction create chat err: %v", err)
		}
		logger.ErrorWithContext(ctx, "failed create chat err: %v", err)
		return nil, err
	}

	if err := command.Commit(ctx); err != nil {
		logger.ErrorWithContext(ctx, "failed commit transaction create chat err: %v", err)
		return nil, err
	}

	return result, nil
}

func (s *Service) CreateChatItem(ctx context.Context, input model.CreateChatItemInput) (*uuid.UUID, error) {
	command := s.repo.NewCommand(ctx)

	result, err := command.CreateChatChatItem(ctx, input)
	if err != nil {
		if err := command.Rollback(ctx); err != nil {
			logger.ErrorWithContext(ctx, "failed rollback transaction create CreateChatItemInput err: %v", err)
		}
		logger.ErrorWithContext(ctx, "failed create CreateChatItemInput err: %v", err)
		return nil, err
	}

	if err := command.Commit(ctx); err != nil {
		logger.ErrorWithContext(ctx, "failed commit transaction create CreateChatItemInput err: %v", err)
		return nil, err
	}

	return result, nil
}

func (s *Service) ReadListChat(ctx context.Context) ([]*model.ChatOutput, error) {
	query := s.repo.NewQuery()
	result, err := query.ReadListChat(ctx)
	if err != nil {
		logger.ErrorWithContext(ctx, "failed get chat err [%v]", err)
	}
	return result, err
}

func (s *Service) ReadListChatItemByID(ctx context.Context, id uuid.UUID) ([]*model.ChatItemOutput, error) {
	query := s.repo.NewQuery()
	result, err := query.ReadChatItemsByChatID(ctx, id)
	if err != nil {
		logger.ErrorWithContext(ctx, "failed get chat err [%v]", err)
	}
	return result, err
}

func (s *Service) SaveImageChat(ctx context.Context, file *multipart.FileHeader) (*UploadPhotoResponse, error) {
	ext := filepath.Ext(file.Filename)
	filename := uuid.New().String() + ext
	err := savefile.SaveFile(file, s.conf.ChatImageConfig.Path+"/"+filename)

	if err != nil {
		return nil, err
	}

	result := UploadPhotoResponse{
		Name: filename,
		Url:  s.conf.ChatImageConfig.Url + filename,
	}

	return &result, nil
}
