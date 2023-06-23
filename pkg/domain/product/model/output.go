package model

import (
	"github.com/e-fish/api/pkg/domain/budidaya/model"
	"github.com/google/uuid"
)

type ProductOutput struct {
	ID         uuid.UUID            `gorm:"primaryKey,size:256" json:"id"`
	BudidayaID uuid.UUID            `gorm:"size:256" json:"budidayaID"`
	Budidaya   model.BudidayaOutput `gorm:"size:256" json:"budidaya"`
	EstPrice   int                  `json:"estPrice"`
}

func (p *ProductOutput) TableName() string {
	return "products"
}
