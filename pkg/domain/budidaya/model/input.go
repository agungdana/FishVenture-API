package model

import (
	"time"

	"github.com/e-fish/api/pkg/common/helper/werror"
	"github.com/e-fish/api/pkg/common/infra/orm"
	errorbudidaya "github.com/e-fish/api/pkg/domain/budidaya/error-budidaya"
	"github.com/google/uuid"
)

type CreatePondInput struct {
	Name          string              `json:"name"`
	CountryID     uuid.UUID           `gorm:"size:256" json:"country_id"`
	ProvinceID    uuid.UUID           `gorm:"size:256" json:"province_id"`
	CityID        uuid.UUID           `gorm:"size:256" json:"city_id"`
	DistrictID    uuid.UUID           `gorm:"size:256" json:"district_id"`
	DetailAddress string              `json:"detailAddress"`
	NoteAddress   string              `json:"noteAddress"`
	Type          string              `json:"type"`
	Latitude      float64             `json:"latitude"`
	Longitude     float64             `json:"longitude"`
	TeamID        uuid.UUID           `gorm:"size:256" json:"team_id"`
	ListPool      []CreatePoolInput   `json:"list_pool"`
	ListBerkas    []CreateBerkasInput `json:"berkas"`
}

func (c *CreatePondInput) Validate() error {
	errs := werror.NewError("error validate input")

	if c.Name == "" {
		errs.Add(errorbudidaya.ErrValidateInputPond.AttacthDetail(map[string]any{"Name": "empty"}))
	}
	if c.CountryID == uuid.Nil {
		errs.Add(errorbudidaya.ErrValidateInputPond.AttacthDetail(map[string]any{"CountryID": "empty"}))
	}
	if c.ProvinceID == uuid.Nil {
		errs.Add(errorbudidaya.ErrValidateInputPond.AttacthDetail(map[string]any{"ProvinceID": "empty"}))
	}
	if c.CityID == uuid.Nil {
		errs.Add(errorbudidaya.ErrValidateInputPond.AttacthDetail(map[string]any{"CityID": "empty"}))
	}
	if c.DistrictID == uuid.Nil {
		errs.Add(errorbudidaya.ErrValidateInputPond.AttacthDetail(map[string]any{"DistrictID": "empty"}))
	}
	if c.Type == "" {
		errs.Add(errorbudidaya.ErrValidateInputPond.AttacthDetail(map[string]any{"Type": "empty"}))
	}

	if c.Type == TEAM {
		if c.TeamID == uuid.Nil {
			errs.Add(errorbudidaya.ErrValidateInputPond.AttacthDetail(map[string]any{"Team": "empty"}))
		}
		errs.Add(ValidateCreateberkasInput(c.ListBerkas))
	}

	if len(c.ListPool) < 1 {
		errs.Add(errorbudidaya.ErrValidateInputPond.AttacthDetail(map[string]any{"Pool": "empty"}))
	}

	errs.Add(ValidateCreatePoolInput(c.ListPool))

	return errs.Return()
}

func (c *CreatePondInput) ToPond(userID uuid.UUID) Pond {
	var (
		pondID = uuid.New()
	)

	return Pond{
		ID:            pondID,
		UserID:        userID,
		Name:          c.Name,
		CountryID:     c.CountryID,
		ProvinceID:    c.ProvinceID,
		CityID:        c.CityID,
		DistrictID:    c.DistrictID,
		DetailAddress: c.DetailAddress,
		NoteAddress:   c.NoteAddress,
		Type:          c.Type,
		Latitude:      c.Latitude,
		Longitude:     c.Longitude,
		TeamID:        c.TeamID,
		Status:        SUBMISION,
		ListPool:      ListPoolInputToListPool(userID, pondID, c.ListPool),
		ListBerkas:    ListBerkasInputToListBerkas(userID, pondID, c.ListBerkas),
		OrmModel: orm.OrmModel{
			CreatedAt: time.Time{},
			CreatedBy: userID,
		},
	}
}

type CreateBerkasInput struct {
	Name string `json:"name"`
	File string `json:"file"`
}

func (c *CreateBerkasInput) Validate() error {
	errs := werror.NewError("error validate input")

	if c.Name == "" {
		errs.Add(errorbudidaya.ErrValidateInputbBerkas.AttacthDetail(map[string]any{"Name": "empty"}))
	}
	if c.File == "" {
		errs.Add(errorbudidaya.ErrValidateInputbBerkas.AttacthDetail(map[string]any{"File": "empty"}))
	}

	return errs.Return()
}

func ValidateCreateberkasInput(input []CreateBerkasInput) error {
	errs := werror.NewError("error validate input")

	for _, v := range input {
		errs.Add(v.Validate())
	}

	return errs.Return()
}

func (c *CreateBerkasInput) ToBerkas(userID uuid.UUID, pondID uuid.UUID) Berkas {
	return Berkas{
		ID:     uuid.New(),
		PondID: pondID,
		Name:   c.Name,
		File:   c.File,
		OrmModel: orm.OrmModel{
			CreatedAt: time.Now(),
			CreatedBy: userID,
		},
	}
}

func ListBerkasInputToListBerkas(userID, pondID uuid.UUID, input []CreateBerkasInput) (newBerkas []Berkas) {
	for _, v := range input {
		newBerkas = append(newBerkas, v.ToBerkas(userID, pondID))
	}
	return
}

type CreatePoolInput struct {
	Name  string  `json:"name"`
	Long  float64 `json:"long"`
	Wide  float64 `json:"wide"`
	Image string  `json:"image"`
}

func (c *CreatePoolInput) Validate() error {
	errs := werror.NewError("error validate input")

	if c.Name == "" {
		errs.Add(errorbudidaya.ErrValidateInputbBerkas.AttacthDetail(map[string]any{"Name": "empty"}))
	}
	if c.Long == 0 {
		errs.Add(errorbudidaya.ErrValidateInputbBerkas.AttacthDetail(map[string]any{"Long": "empty"}))
	}
	if c.Wide == 0 {
		errs.Add(errorbudidaya.ErrValidateInputbBerkas.AttacthDetail(map[string]any{"Wide": "empty"}))
	}
	if c.Image == "" {
		errs.Add(errorbudidaya.ErrValidateInputbBerkas.AttacthDetail(map[string]any{"Image": "empty"}))
	}

	return errs.Return()
}

func ValidateCreatePoolInput(input []CreatePoolInput) error {
	errs := werror.NewError("error validate input")

	for _, v := range input {
		errs.Add(v.Validate())
	}

	return errs.Return()
}

func (c *CreatePoolInput) ToPool(userID uuid.UUID, pondID uuid.UUID) Pool {
	return Pool{
		ID:       uuid.New(),
		PondID:   pondID,
		Name:     c.Name,
		Long:     c.Long,
		Wide:     c.Wide,
		Image:    c.Image,
		OrmModel: orm.OrmModel{CreatedAt: time.Now(), CreatedBy: userID},
	}
}

func ListPoolInputToListPool(userID, pondID uuid.UUID, input []CreatePoolInput) (newPool []Pool) {
	for _, v := range input {
		newPool = append(newPool, v.ToPool(userID, pondID))
	}
	return
}

type UpdatePondInput struct {
	Name          string    `json:"name"`
	CountryID     uuid.UUID `gorm:"size:256" json:"country_id"`
	ProvinceID    uuid.UUID `gorm:"size:256" json:"province_id"`
	CityID        uuid.UUID `gorm:"size:256" json:"city_id"`
	DistrictID    uuid.UUID `gorm:"size:256" json:"district_id"`
	DetailAddress string    `json:"detailAddress"`
	NoteAddress   string    `json:"noteAddress"`
	Type          string    `json:"type"`
	Latitude      float64   `json:"latitude"`
	Longitude     float64   `json:"longitude"`
}

func (u *UpdatePondInput) ToPond(userID, pondID uuid.UUID) Pond {
	var (
		today = time.Now()
	)

	return Pond{
		ID:            pondID,
		Name:          u.Name,
		CountryID:     u.CountryID,
		ProvinceID:    u.ProvinceID,
		CityID:        u.CityID,
		DistrictID:    u.DistrictID,
		DetailAddress: u.DetailAddress,
		NoteAddress:   u.NoteAddress,
		Type:          u.Type,
		Latitude:      u.Latitude,
		Longitude:     u.Longitude,
		OrmModel: orm.OrmModel{
			UpdatedAt: &today,
			UpdatedBy: &userID,
		},
	}
}

type UpdatePondStatus struct {
	PondID uuid.UUID
	Status string
}

func (u *UpdatePondStatus) Validate() error {
	errs := werror.NewError("failed validate input update pond status")

	if u.PondID == uuid.Nil {
		errs.Add(errorbudidaya.ErrValidateInputbUpdateStatus.AttacthDetail(map[string]any{"PondID": "empty"}))
	}
	if u.Status == "" {
		errs.Add(errorbudidaya.ErrValidateInputbUpdateStatus.AttacthDetail(map[string]any{"Status": "empty"}))
	}

	return errs.Return()
}

func (u *UpdatePondStatus) ToPond(userID uuid.UUID) Pond {
	var (
		today = time.Now()
	)

	return Pond{
		ID:     u.PondID,
		Status: u.Status,
		OrmModel: orm.OrmModel{
			UpdatedAt: &today,
			UpdatedBy: &userID,
		},
	}
}
