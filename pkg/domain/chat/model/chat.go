package model

import (
	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/google/uuid"
)

type Chat struct {
	ID     uuid.UUID `gorm:"primaryKey,size:256"`
	UserID uuid.UUID
	User   User
	PondID uuid.UUID
	Pond   Pond
	orm.OrmModel
}

type ChatItem struct {
	ID         uuid.UUID `gorm:"primaryKey,size:256"`
	ChatID     uuid.UUID
	SenderID   uuid.UUID
	ReceiverID uuid.UUID
	Image      string
	Text       string
	Payload    string
	Type       string
	orm.OrmModel
}
