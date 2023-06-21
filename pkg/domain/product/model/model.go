package model

import (
	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/e-fish/api/pkg/domain/budidaya/model"
	"github.com/google/uuid"
)

type Product struct {
	ID         uuid.UUID `gorm:"primaryKey,size:256"`
	PondID     uuid.UUID `gorm:"size:256"`
	BudidayaID uuid.UUID `gorm:"size:256"`
	Budidaya   model.Budidaya
	EstPrice   int
	orm.OrmModel
}
