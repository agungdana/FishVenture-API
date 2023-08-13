package model

import (
	"time"

	"github.com/e-fish/api/pkg/common/helper/werror"
	"github.com/e-fish/api/pkg/common/infra/orm"
	errorbudidaya "github.com/e-fish/api/pkg/domain/budidaya/error-budidaya"
	"github.com/google/uuid"
)

type CreateBudidayaInput struct {
	Code          string    `json:"code"`
	PoolID        uuid.UUID `json:"poolID"`
	DateOfSeed    time.Time `json:"dateOfSeed"`
	FishSpeciesID uuid.UUID `json:"fishSpeciesID"`
}

func (c *CreateBudidayaInput) Validate() error {
	errs := werror.NewError("failed validate create input budidaya")

	if c.PoolID == uuid.Nil {
		errs.Add(errorbudidaya.ErrValidateInputBudidaya.AttacthDetail(map[string]any{"poolID": "empty"}))
	}
	if c.DateOfSeed.IsZero() {
		errs.Add(errorbudidaya.ErrValidateInputBudidaya.AttacthDetail(map[string]any{"dateOfSeed": "empty"}))
	}
	if c.FishSpeciesID == uuid.Nil {
		errs.Add(errorbudidaya.ErrValidateInputBudidaya.AttacthDetail(map[string]any{"fishSpeciesID": "empty"}))
	}

	return errs.Return()
}

func (c *CreateBudidayaInput) ToBudidaya(userID, pondID uuid.UUID) Budidaya {
	return Budidaya{
		ID:            uuid.New(),
		Code:          c.Code,
		PondID:        pondID,
		PoolID:        c.PoolID,
		DateOfSeed:    c.DateOfSeed,
		FishSpeciesID: c.FishSpeciesID,
		Status:        BUDIDAYA,
		OrmModel: orm.OrmModel{
			CreatedAt: time.Now(),
			CreatedBy: userID,
		},
	}
}

type UpdateBudidayaStatusInput struct {
	ID        uuid.UUID `json:"id"`
	EstTonase int       `json:"estTonase"`
	EstDate   time.Time `json:"estDate"`
	Status    string    `json:"status"`
}

func (c *UpdateBudidayaStatusInput) ToBudidaya(userID uuid.UUID) Budidaya {
	today := time.Now()

	return Budidaya{
		ID:           c.ID,
		Status:       c.Status,
		EstTonase:    float64(c.EstTonase),
		EstPanenDate: &c.EstDate,
		OrmModel: orm.OrmModel{
			UpdatedBy: &userID,
			UpdatedAt: &today,
		},
	}
}

type CreateMultiplePriceListInput struct {
	BudidayaID uuid.UUID              `json:"budidayaID"`
	EstTonase  int                    `json:"estTonase"`
	EstDate    time.Time              `json:"estDate"`
	Input      []CreatePriceListInput `json:"input"`
}

func (c *CreateMultiplePriceListInput) Validate() error {
	errs := werror.NewError("failed validate create input budidaya")

	if c.BudidayaID == uuid.Nil {
		errs.Add(errorbudidaya.ErrValidateInputPriceList.AttacthDetail(map[string]any{"budidayaID": "empty"}))
	}
	if c.EstTonase < 1 {
		errs.Add(errorbudidaya.ErrValidateInputPriceList.AttacthDetail(map[string]any{"estTonase": "empty"}))
	}
	if c.EstDate.IsZero() {
		errs.Add(errorbudidaya.ErrValidateInputPriceList.AttacthDetail(map[string]any{"estDate": "empty"}))
	}

	for _, v := range c.Input {
		v.BudidayaID = c.BudidayaID
		v.EstTonase = c.EstTonase
		v.EstDate = c.EstDate
		err := v.Validate()
		if err != nil {
			errs.Add(errorbudidaya.ErrValidateMultipleInputPriceList.AttacthDetail(map[string]any{"limit": v.Limit, "error": err}))
		}
	}

	return errs.Return()
}

func (c *CreateMultiplePriceListInput) ToMultiplePriceList(userID uuid.UUID) (newPricelist []PriceList) {
	for _, v := range c.Input {
		newPricelist = append(newPricelist, PriceList{
			ID:         uuid.New(),
			BudidayaID: c.BudidayaID,
			Limit:      v.Limit,
			Price:      v.Price,
			OrmModel: orm.OrmModel{
				CreatedAt: time.Now(),
				CreatedBy: userID,
			},
		})
	}

	return
}

type CreatePriceListInput struct {
	BudidayaID uuid.UUID `json:"budidayaID"`
	EstTonase  int       `json:"estTonase"`
	EstDate    time.Time `json:"estDate"`
	Limit      int       `json:"limit"`
	Price      int       `json:"price"`
}

func (c *CreatePriceListInput) Validate() error {
	errs := werror.NewError("failed validate create input budidaya")

	if c.BudidayaID == uuid.Nil {
		errs.Add(errorbudidaya.ErrValidateInputPriceList.AttacthDetail(map[string]any{"budidayaID": "empty"}))
	}
	if c.EstTonase < 1 {
		errs.Add(errorbudidaya.ErrValidateInputPriceList.AttacthDetail(map[string]any{"estTonase": "empty"}))
	}
	if c.EstDate.IsZero() {
		errs.Add(errorbudidaya.ErrValidateInputPriceList.AttacthDetail(map[string]any{"estDate": "empty"}))
	}
	if c.Limit < 1 {
		errs.Add(errorbudidaya.ErrValidateInputPriceList.AttacthDetail(map[string]any{"limit": "empty"}))
	}
	if c.Price < 1 {
		errs.Add(errorbudidaya.ErrValidateInputPriceList.AttacthDetail(map[string]any{"price": "empty"}))
	}

	return errs.Return()
}

func (c *CreatePriceListInput) ToPriceList(userID uuid.UUID) PriceList {
	return PriceList{
		ID:         uuid.New(),
		BudidayaID: c.BudidayaID,
		Limit:      c.Limit,
		Price:      c.Price,
		OrmModel: orm.OrmModel{
			CreatedAt: time.Now(),
			CreatedBy: userID,
		},
	}
}

type CreateFishSpeciesInput struct {
	Name string `json:"name,omitempty"`
	Asal string `json:"asal,omitempty"`
}

func (c *CreateFishSpeciesInput) Validate() error {
	errs := werror.NewError("failed validate create input budidaya")

	if c.Name == "" {
		errs.Add(errorbudidaya.ErrValidateInputFishSpecies.AttacthDetail(map[string]any{"name": "empty"}))
	}
	if c.Asal == "" {
		errs.Add(errorbudidaya.ErrValidateInputFishSpecies.AttacthDetail(map[string]any{"name": "empty"}))
	}

	return errs.Return()
}

func (c *CreateFishSpeciesInput) ToFishSpecies(userID uuid.UUID) FishSpecies {
	return FishSpecies{
		ID:       uuid.New(),
		Name:     c.Name,
		Asal:     c.Asal,
		OrmModel: orm.OrmModel{CreatedAt: time.Now(), CreatedBy: userID},
	}
}

type GetBudidayaInput struct {
	PondID uuid.UUID
}

type ReadPricelistBudidayaInput struct {
	BudidayaID uuid.UUID
	Qty        int
}
