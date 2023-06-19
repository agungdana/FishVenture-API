package model

import (
	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/google/uuid"
)

type TeamOutput struct {
	ID         uuid.UUID `gorm:"size:256"`
	Name       string
	CountryID  uuid.UUID `gorm:"size:256"`
	ProvinceID uuid.UUID `gorm:"size:256"`
	CityID     uuid.UUID `gorm:"size:256"`
	DistrictID uuid.UUID `gorm:"size:256"`
	Detail     string
	Note       string
	ListPond   []PondOutput `gorm:"foreignKey:TeamID;references:ID"`
}

func (t *TeamOutput) TableName() string {
	return "teams"
}

type PondOutput struct {
	ID            uuid.UUID `gorm:"size:256" json:"id"`
	UserID        uuid.UUID `gorm:"size:256" json:"user_id"`
	User          UserPond
	Name          string         `json:"name"`
	CountryID     uuid.UUID      `gorm:"size:256" json:"country_id"`
	ProvinceID    uuid.UUID      `gorm:"size:256" json:"province_id"`
	CityID        uuid.UUID      `gorm:"size:256" json:"city_id"`
	DistrictID    uuid.UUID      `gorm:"size:256" json:"district_id"`
	DetailAddress string         `json:"detailAddress"`
	NoteAddress   string         `json:"noteAddress"`
	Type          string         `json:"type"`
	Latitude      float64        `json:"latitude"`
	Longitude     float64        `json:"longitude"`
	TeamID        uuid.UUID      `gorm:"size:256" json:"team_id"`
	Team          *TeamOutput    `json:"team,omitempty" gorm:"foreignKey:TeamID;references:ID"`
	Status        string         `json:"status"`
	ListPool      []PoolOutput   `json:"list_pool,omitempty" gorm:"foreignKey:PondID;references:ID"`
	ListBerkas    []BerkasOutput `json:"berkas,omitempty" gorm:"foreignKey:PondID;references:ID"`
	orm.OrmModel
}

func (t *PondOutput) TableName() string {
	return "ponds"
}

type BerkasOutput struct {
	ID     uuid.UUID   `gorm:"primaryKey,size:256" json:"id"`
	PondID uuid.UUID   `gorm:"size:256" json:"pond_id"`
	Pond   *PondOutput `gorm:"foreignKey:PondID;references:ID" json:"pond,omitempty"`
	Name   string      `json:"name"`
	File   string      `json:"file"`
	orm.OrmModel
}

func (t *BerkasOutput) TableName() string {
	return "berkas"
}

type PoolOutput struct {
	ID     uuid.UUID   `gorm:"size:256" json:"id"`
	PondID uuid.UUID   `gorm:"size:256" json:"pond_id"`
	Pond   *PondOutput `json:"pond,omitempty" gorm:"foreignKey:PondID;references:ID"`
	Name   string      `json:"name"`
	Long   float64     `json:"long"`
	Wide   float64     `json:"wide"`
	Image  string      `json:"image"`
	orm.OrmModel
}

func (t *PoolOutput) TableName() string {
	return "pools"
}
