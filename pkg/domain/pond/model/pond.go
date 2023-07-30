package model

import (
	"time"

	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/google/uuid"
)

type UserPond struct {
	ID                  uuid.UUID `gorm:"primaryKey,size:256"`
	Name                string
	VarificationCode    string
	ExpVerificationCode time.Time
}

func (u UserPond) TableName() string {
	return "users"
}

type Team struct {
	ID         uuid.UUID `gorm:"primaryKey,size:256"`
	Name       string
	CountryID  uuid.UUID `gorm:"size:256"`
	ProvinceID uuid.UUID `gorm:"size:256"`
	CityID     uuid.UUID `gorm:"size:256"`
	DistrictID uuid.UUID `gorm:"size:256"`
	Detail     string
	Note       string
	ListPond   []*Pond
	orm.OrmModel
}

type Pond struct {
	ID            uuid.UUID `gorm:"primaryKey,size:256" json:"id"`
	UserID        uuid.UUID `gorm:"size:256" json:"user_id"`
	User          UserPond
	Name          string     `json:"name"`
	CountryID     uuid.UUID  `gorm:"size:256" json:"countryID"`
	ProvinceID    uuid.UUID  `gorm:"size:256" json:"provinceID"`
	CityID        uuid.UUID  `gorm:"size:256" json:"cityID"`
	DistrictID    uuid.UUID  `gorm:"size:256" json:"districtID"`
	DetailAddress string     `json:"detailAddress"`
	NoteAddress   string     `json:"noteAddress"`
	Type          string     `json:"type"`
	Latitude      float64    `json:"latitude"`
	Longitude     float64    `json:"longitude"`
	TeamID        *uuid.UUID `gorm:"size:256" json:"teamID"`
	Team          Team       `json:"team"`
	Status        string     `json:"status"`
	Image         string     `json:"image"`
	ListPool      []Pool     `json:"listPool"`
	ListBerkas    []Berkas   `json:"berkas"`
	orm.OrmModel
}

type Berkas struct {
	ID     uuid.UUID `gorm:"primaryKey,size:256" json:"id"`
	PondID uuid.UUID `gorm:"size:256"`
	Pond   Pond
	Name   string `json:"name"`
	File   string `json:"file"`
	orm.OrmModel
}

type Pool struct {
	ID     uuid.UUID `gorm:"primaryKey,size:256" json:"id"`
	PondID uuid.UUID `gorm:"size:256" json:"pond_id"`
	Pond   Pond      `json:"pond"`
	Name   string    `json:"name"`
	Long   float64   `json:"long"`
	Wide   float64   `json:"wide"`
	Image  string    `json:"image"`
	orm.OrmModel
}
