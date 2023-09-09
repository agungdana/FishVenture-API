package chat

import (
	"context"

	"github.com/e-fish/api/pkg/domain/chat/model"
	"github.com/google/uuid"
)

type Repo interface {
	NewCommand(ctx context.Context) Command
	NewQuery() Query
}

type Command interface {
	CreateChat(ctx context.Context, input model.CreateChatInput) (*uuid.UUID, error)
	CreateChatChatItem(ctx context.Context, input model.CreateChatItemInput) (*uuid.UUID, error)

	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
}

type Query interface {
	ReadChatByID(ctx context.Context, id uuid.UUID) (*model.Chat, error)
	ReadChatByCsAndPondID(ctx context.Context, csID, pondID uuid.UUID) (*model.Chat, error)
	ReadListChat(ctx context.Context) ([]*model.ChatOutput, error)
	ReadChatItemsByChatID(ctx context.Context, id uuid.UUID) ([]*model.ChatItemOutput, error)
	ReadChatItemsByID(ctx context.Context, id uuid.UUID) (*model.ChatItemOutput, error)
	lock() Query
}
