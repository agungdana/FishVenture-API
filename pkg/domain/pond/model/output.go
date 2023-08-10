package model

import (
	"fmt"

	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/e-fish/api/pkg/domain/region/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TeamOutput struct {
	ID         uuid.UUID    `gorm:"size:256" json:"id"`
	Name       string       `json:"name"`
	CountryID  uuid.UUID    `gorm:"size:256" json:"countryID"`
	ProvinceID uuid.UUID    `gorm:"size:256" json:"provinceID"`
	CityID     uuid.UUID    `gorm:"size:256" json:"cityID"`
	DistrictID uuid.UUID    `gorm:"size:256" json:"districtID"`
	Detail     string       `json:"detail"`
	Note       string       `json:"note"`
	ListPond   []PondOutput `gorm:"foreignKey:TeamID;references:ID" json:"listPond,omitempty"`
}

func (t *TeamOutput) TableName() string {
	return "teams"
}

type PondOutput struct {
	ID            uuid.UUID            `gorm:"size:256" json:"id"`
	UserID        uuid.UUID            `json:"userID"`
	User          *UserPond            `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	Name          string               `json:"name"`
	CountryID     uuid.UUID            `gorm:"size:256" json:"countryID"`
	Country       model.CountryOutput  `gorm:"foreignKey:CountryID;references:ID" json:"country"`
	ProvinceID    uuid.UUID            `gorm:"size:256" json:"provinceID"`
	Province      model.ProvinceOutput `gorm:"foreignKey:ProvinceID;references:ID" json:"province"`
	CityID        uuid.UUID            `gorm:"size:256" json:"cityID"`
	City          model.CityOutput     `gorm:"foreignKey:CityID;references:ID" json:"city"`
	DistrictID    uuid.UUID            `gorm:"size:256" json:"districtID"`
	District      model.DistrictOutput `gorm:"foreignKey:DistrictID;references:ID" json:"district"`
	DetailAddress string               `json:"detailAddress"`
	NoteAddress   string               `json:"noteAddress"`
	Type          string               `json:"type"`
	Latitude      float64              `json:"latitude"`
	Longitude     float64              `json:"longitude"`
	Url           string               `json:"url"`
	TeamID        *uuid.UUID           `gorm:"size:256" json:"teamID,omitempty"`
	Team          *TeamOutput          `json:"team,omitempty" gorm:"foreignKey:TeamID;references:ID"`
	Status        string               `json:"status"`
	Image         string               `json:"image"`
	ListPool      []PoolOutput         `json:"listPool,omitempty" gorm:"foreignKey:PondID;references:ID"`
	ListBerkas    []BerkasOutput       `json:"berkas,omitempty" gorm:"foreignKey:PondID;references:ID"`
}

func (t *PondOutput) TableName() string {
	return "ponds"
}

func (t *PondOutput) AfterFind(db *gorm.DB) (err error) {
	t.Url = fmt.Sprintf("https://www.google.com/maps/search/?api=1&query=%v,%v", t.Latitude, t.Longitude)
	return
}

type BerkasOutput struct {
	ID     uuid.UUID   `gorm:"primaryKey,size:256" json:"id"`
	PondID uuid.UUID   `gorm:"size:256" json:"pondID"`
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
	PondID uuid.UUID   `gorm:"size:256" json:"pondID"`
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
