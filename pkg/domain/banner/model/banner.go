package model

import (
	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/google/uuid"
)

type Banner struct {
	ID          uuid.UUID `gorm:"primaryKey,size:256" json:"id"`
	Name        string    `json:"name"`
	Link        string    `json:"link"`
	Description string    `json:"description"`
	orm.OrmModel
}
