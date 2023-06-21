package model

import (
	"github.com/e-fish/api/pkg/domain/budidaya/model"
	"github.com/google/uuid"
)

type ProductOutput struct {
	ID         uuid.UUID `gorm:"primaryKey,size:256"`
	BudidayaID uuid.UUID `gorm:"size:256"`
	Budidaya   model.BudidayaOutput
	EstPrice   int
}
