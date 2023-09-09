package model

import (
	"github.com/google/uuid"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Photo string `json:"photo"`
}

type Pond struct {
	ID    uuid.UUID `gorm:"size:256" json:"id"`
	Name  string    `json:"name"`
	Image string    `json:"image"`
}

type ChatOutput struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"-"`
	User   User      `json:"user"`
	PondID uuid.UUID `json:"-"`
	Pond   Pond      `json:"pond"`
}

func (c *ChatOutput) TableName() string {
	return "chats"
}

type ChatItemOutput struct {
	ID         uuid.UUID `gorm:"primaryKey,size:256" json:"id"`
	ChatID     uuid.UUID `json:"chatID"`
	SenderID   uuid.UUID `json:"senderID"`
	ReceiverID uuid.UUID `json:"receiverID"`
	Image      string    `json:"image"`
	Text       string    `json:"text"`
	Payload    string    `json:"payload"`
	Type       string    `json:"type"`
}
