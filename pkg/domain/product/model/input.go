package model

import (
	"time"

	"github.com/e-fish/api/pkg/common/helper/werror"
	"github.com/e-fish/api/pkg/common/infra/orm"
	errorproduct "github.com/e-fish/api/pkg/domain/product/error-product"
	"github.com/google/uuid"
)

type CreateProductInput struct {
	PondID       uuid.UUID `gorm:"size:256" json:"pondID"`
	BudidayaID   uuid.UUID `gorm:"size:256" json:"budidayaID"`
	EstTonase    float64   `json:"estTonase"`
	EstPanenDate time.Time `json:"estPanenDate"`
	EstPrice     int       `json:"estPrice"`
}

func (c *CreateProductInput) Validate() error {
	errs := werror.NewError("failed validate create product input")

	if c.BudidayaID == uuid.Nil {
		errs.Add(errorproduct.ErrValidateCreateInput.AttacthDetail(map[string]any{"budidayaID": "empty"}))
	}
	if c.EstTonase == 0 {
		errs.Add(errorproduct.ErrValidateCreateInput.AttacthDetail(map[string]any{"estTonase": "empty"}))
	}
	if c.EstPanenDate.IsZero() {
		errs.Add(errorproduct.ErrValidateCreateInput.AttacthDetail(map[string]any{"estPanenDate": "empty"}))
	}
	if c.EstPrice == 0 {
		errs.Add(errorproduct.ErrValidateCreateInput.AttacthDetail(map[string]any{"estPrice": "empty"}))
	}

	return errs.Return()
}

func (c *CreateProductInput) ToProduct(userID, pondID uuid.UUID) Product {
	return Product{
		ID:         uuid.New(),
		PondID:     pondID,
		BudidayaID: c.BudidayaID,
		EstPrice:   c.EstPrice,
		OrmModel: orm.OrmModel{
			CreatedAt: time.Now(),
			CreatedBy: userID,
		},
	}
}
