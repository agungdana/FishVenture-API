package model

import (
	"time"

	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/e-fish/api/pkg/domain/budidaya/model"
	"github.com/google/uuid"
)

type User struct {
	ID   uuid.UUID `gorm:"primaryKey,size:256"`
	Name string
}

func (u *User) TableName() string {
	return "users"
}

type Order struct {
	ID          uuid.UUID `gorm:"primaryKey,size:256"`
	Code        string
	BudidayaID  uuid.UUID
	Budidaya    model.Budidaya
	UserID      uuid.UUID
	User        User
	Qty         int
	BookingDate time.Time
	Status      string
	orm.OrmModel
}
