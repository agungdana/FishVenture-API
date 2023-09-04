package model

import (
	"time"

	"github.com/e-fish/api/pkg/common/helper/werror"
	"github.com/e-fish/api/pkg/common/infra/orm"
	errorBanner "github.com/e-fish/api/pkg/domain/banner/model/error-banner"
	"github.com/google/uuid"
)

type BannerInputCreate struct {
	Name        string `json:"name"`
	Link        string `json:"link"`
	Description string `json:"description"`
}

func (b *BannerInputCreate) Validate() error {
	errs := werror.NewError("failed validate input banner")

	if b.Name == "" {
		errs.Add(errorBanner.ErrValidateInputBanner.AttacthDetail(map[string]any{"name": "empty"}))
	}
	if b.Link == "" {
		errs.Add(errorBanner.ErrValidateInputBanner.AttacthDetail(map[string]any{"link": "empty"}))
	}
	if b.Description == "" {
		errs.Add(errorBanner.ErrValidateInputBanner.AttacthDetail(map[string]any{"description": "empty"}))
	}

	if err := errs.Return(); err != nil {
		return err
	}
	return nil
}

func (b *BannerInputCreate) NewBanner(userID uuid.UUID) Banner {

	return Banner{
		ID:          uuid.New(),
		Name:        b.Name,
		Link:        b.Link,
		Description: b.Description,
		OrmModel: orm.OrmModel{
			CreatedAt: time.Now(),
			CreatedBy: userID,
		},
	}
}

type BannerInputUpdate struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Link        string    `json:"link"`
	Description string    `json:"description"`
}

func (b *BannerInputUpdate) Validate() error {
	errs := werror.NewError("failed validate input banner")

	if b.ID == uuid.Nil {
		errs.Add(errorBanner.ErrValidateInputBanner.AttacthDetail(map[string]any{"id": "empty"}))
	}
	if b.Name == "" {
		errs.Add(errorBanner.ErrValidateInputBanner.AttacthDetail(map[string]any{"name": "empty"}))
	}
	if b.Link == "" {
		errs.Add(errorBanner.ErrValidateInputBanner.AttacthDetail(map[string]any{"link": "empty"}))
	}
	if b.Description == "" {
		errs.Add(errorBanner.ErrValidateInputBanner.AttacthDetail(map[string]any{"description": "empty"}))
	}

	if err := errs.Return(); err != nil {
		return err
	}
	return nil
}

func (b *BannerInputUpdate) NewBanner(userID uuid.UUID) Banner {
	today := time.Now()
	return Banner{
		ID:          b.ID,
		Name:        b.Name,
		Link:        b.Link,
		Description: b.Description,
		OrmModel: orm.OrmModel{
			UpdatedAt: &today,
			UpdatedBy: &userID,
		},
	}
}
