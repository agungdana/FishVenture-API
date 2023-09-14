package chat

import (
	"context"
	"errors"

	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	userModel "github.com/e-fish/api/pkg/domain/auth/model"
	"github.com/e-fish/api/pkg/domain/chat/model"
	errorchat "github.com/e-fish/api/pkg/domain/chat/model/error-chat"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func newQuery(db *gorm.DB) Query {
	return &query{db: db}
}

// lock implements Query.
func (q *query) lock() Query {
	db := q.db.Clauses(clause.Locking{Strength: "UPDATE"})
	return &query{db: db}
}

type query struct {
	db *gorm.DB
}

// ReadChatItemsByID implements Query.
func (q *query) ReadChatItemsByID(ctx context.Context, id uuid.UUID) (*model.ChatItemOutput, error) {
	var (
		data = model.ChatItemOutput{}
	)

	err := q.db.Where("deleted_at IS NULL and id = ?", id).Take(&data).Error
	if err != nil {
		return nil, errorchat.ErrReadChatData.AttacthDetail(map[string]any{"err": err})
	}

	return &data, nil
}

// ReadChatByID implements Query.
func (q *query) ReadChatByID(ctx context.Context, id uuid.UUID) (*model.Chat, error) {
	var (
		data = model.Chat{}
	)
	err := q.db.Where("deleted_at IS NULL and id = ?", id).Take(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorchat.ErrFoundChat.AttacthDetail(map[string]any{"err": err})
		}
		return nil, errorchat.ErrReadChatData.AttacthDetail(map[string]any{"err": err})
	}
	return &data, nil
}

// ReadChatByCsAndPondID implements Query.
func (q *query) ReadChatByCsAndPondID(ctx context.Context, csID uuid.UUID, pondID uuid.UUID) (*model.Chat, error) {
	var (
		data = model.Chat{}
	)
	err := q.db.Where("deleted_at IS NULL and user_id = ? AND pond_id = ?", csID, pondID).Take(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorchat.ErrFoundChat.AttacthDetail(map[string]any{"err": err})
		}
		return nil, errorchat.ErrReadChatData.AttacthDetail(map[string]any{"err": err})
	}
	return &data, nil
}

// ReadChatItemsByChatID implements Query.
func (q *query) ReadChatItemsByChatID(ctx context.Context, id uuid.UUID) ([]*model.ChatItemOutput, error) {
	var (
		data      = []*model.ChatItemOutput{}
		userID, _ = ctxutil.GetUserID(ctx)
		pondID, _ = ctxutil.GetPondID(ctx)
	)

	err := q.db.Where("deleted_at IS NULL and chat_id = ?", id).Order("created_at desc").Find(&data).Error
	if err != nil {
		return nil, errorchat.ErrReadChatData.AttacthDetail(map[string]any{"err": err})
	}

	for idx := range data {
		if userID == data[idx].SenderID {
			data[idx].IsMe = true
		}
		if pondID == data[idx].SenderID {
			data[idx].IsMe = true
		}
	}

	return data, nil
}

// ReadListChat implements Query.
func (q *query) ReadListChat(ctx context.Context) ([]*model.ChatOutput, error) {
	var (
		data       = []*model.ChatOutput{}
		userID, _  = ctxutil.GetUserID(ctx)
		pondID, _  = ctxutil.GetPondID(ctx)
		appType, _ = ctxutil.GetUserAppType(ctx)
		db         = q.db
	)

	switch appType {
	case userModel.BUYER:
		db = db.Where("user_id = ?", userID)
	case userModel.SELLER:
		db = db.Where("pond_id = ?", pondID)
	}

	err := db.Where("deleted_at IS NULL").Order("created_at desc").Find(&data).Error
	if err != nil {
		return nil, errorchat.ErrReadChatData.AttacthDetail(map[string]any{"err": err})
	}

	return data, nil
}
