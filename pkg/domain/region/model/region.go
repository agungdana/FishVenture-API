package model

import (
	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/google/uuid"
)

type Country struct {
	ID           uuid.UUID `gorm:"primarykey"`
	Name         string
	IsCoverage   bool
	ListProvince []Province
	orm.OrmModel
}

type Province struct {
	ID         uuid.UUID `gorm:"primarykey"`
	CountryID  uuid.UUID
	Country    Country
	Name       string
	IsCoverage bool
	ListCity   []City
	orm.OrmModel
}

type City struct {
	ID           uuid.UUID `gorm:"primarykey"`
	ProvinceID   uuid.UUID
	Province     Province
	Name         string
	ListDistrict []District
	IsCoverage   bool
	orm.OrmModel
}

type District struct {
	ID         uuid.UUID `gorm:"primarykey"`
	CityID     uuid.UUID
	City       City
	Name       string
	IsCoverage bool
	orm.OrmModel
}
