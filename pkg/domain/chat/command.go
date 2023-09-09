package chat

import (
	"context"

	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/e-fish/api/pkg/domain/chat/model"
	errorchat "github.com/e-fish/api/pkg/domain/chat/model/error-chat"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func newCommand(ctx context.Context, db *gorm.DB) Command {
	var (
		dbTxn = orm.BeginTxn(ctx, db.WithContext(ctx))
	)

	return &command{
		dbTxn: dbTxn,
		query: newQuery(dbTxn),
	}
}

type command struct {
	dbTxn *gorm.DB
	query Query
}

// CreateChatChatItem implements Command.
func (c *command) CreateChatChatItem(ctx context.Context, input model.CreateChatItemInput) (*uuid.UUID, error) {
	var (
		userID, _  = ctxutil.GetUserID(ctx)
		senderID   = uuid.UUID{}
		receiverID = uuid.UUID{}
	)

	exist, err := c.query.ReadChatByID(ctx, input.ChatID)
	if err != nil {
		return nil, err
	}

	if exist == nil {
		return nil, errorchat.ErrFailedCreateChat
	}

	if exist.UserID == userID {
		senderID = exist.UserID
		receiverID = exist.PondID
	} else {
		senderID = exist.PondID
		receiverID = exist.UserID
	}

	newItem := input.NewChatItem(userID, senderID, receiverID)

	err = c.dbTxn.Create(&newItem).Error
	if err != nil {
		return nil, errorchat.ErrFailedCreateChatItem.AttacthDetail(map[string]any{"error": err})
	}

	return &newItem.ID, nil
}

// CreateChat implements Command.
func (c *command) CreateChat(ctx context.Context, input model.CreateChatInput) (*uuid.UUID, error) {
	var (
		userID, _ = ctxutil.GetUserID(ctx)
	)

	if input.ReceiverID == uuid.Nil {
		return nil, errorchat.ErrFailedCreateChat.AttacthDetail(map[string]any{"receiverID": "empty"})
	}

	exist, err := c.query.ReadChatByCsAndPondID(ctx, userID, input.ReceiverID)
	if err != nil {
		if !errorchat.ErrFoundChat.Is(err) {
			return nil, err
		}
		return &exist.ID, nil
	}

	newChat := input.NewChat(userID)

	err = c.dbTxn.Create(&newChat).Error
	if err != nil {
		return nil, err
	}

	return &newChat.ID, nil
}

// Commit implements Command.
func (c *command) Commit(ctx context.Context) error {
	if err := orm.CommitTxn(ctx); err != nil {
		return errorchat.ErrCommit.AttacthDetail(map[string]any{"error": err})
	}
	return nil
}

// Rollback implements Command.
func (c *command) Rollback(ctx context.Context) error {
	if err := orm.RollbackTxn(ctx); err != nil {
		return errorchat.ErrRollback.AttacthDetail(map[string]any{"error": err})
	}
	return nil
}
