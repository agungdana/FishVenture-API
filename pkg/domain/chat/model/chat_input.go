package model

import (
	"time"

	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/google/uuid"
)

type CreateChatInput struct {
	ReceiverID uuid.UUID `json:"receiverID"`
}

func (c *CreateChatInput) NewChat(userID uuid.UUID) Chat {
	return Chat{
		ID:       uuid.New(),
		UserID:   userID,
		PondID:   c.ReceiverID,
		OrmModel: orm.OrmModel{CreatedAt: time.Now(), CreatedBy: userID},
	}
}

type CreateChatItemInput struct {
	ID      uuid.UUID `json:"id"`
	ChatID  uuid.UUID `json:"chatID"`
	Image   string    `json:"image"`
	Text    string    `json:"text"`
	Payload string    `json:"payload"`
	Type    string    `json:"type"`
}

func (c *CreateChatItemInput) NewChatItem(userID, senderID, receiverID uuid.UUID) ChatItem {
	return ChatItem{
		ID:         uuid.New(),
		ChatID:     c.ChatID,
		SenderID:   senderID,
		ReceiverID: receiverID,
		Image:      c.Image,
		Text:       c.Text,
		Payload:    c.Payload,
		Type:       c.Type,
		OrmModel: orm.OrmModel{
			CreatedAt: time.Now(),
			CreatedBy: userID,
		},
	}
}
