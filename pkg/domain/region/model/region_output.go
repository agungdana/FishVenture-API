package model

import (
	"github.com/google/uuid"
)

type CountryOutput struct {
	ID           uuid.UUID        `gorm:"primarykey" json:"id"`
	Name         string           `json:"name"`
	ListProvince []ProvinceOutput `json:"listProvince,omitempty" gorm:"foreignKey:CountryID;references:ID"`
}

func (c *CountryOutput) TableName() string {
	return "countries"
}

type ProvinceOutput struct {
	ID        uuid.UUID      `gorm:"primarykey" json:"id"`
	CountryID uuid.UUID      `json:"countryID"`
	Country   *CountryOutput `json:"country,omitempty" gorm:"foreignKey:ID;references:CountryID"`
	Name      string         `json:"name"`
	ListCity  []CityOutput   `json:"listCity,omitempty" gorm:"foreignKey:ProvinceID;references:ID"`
}

func (c *ProvinceOutput) TableName() string {
	return "provinces"
}

type CityOutput struct {
	ID           uuid.UUID         `gorm:"primarykey" json:"id"`
	ProvinceID   uuid.UUID         `json:"provinceID"`
	Province     *ProvinceOutput   `json:"province,omitempty" gorm:"foreignKey:ID;references:ProvinceID"`
	Name         string            `json:"name"`
	ListDistrict []*DistrictOutput `json:"listDistrict,omitempty" gorm:"foreignKey:CityID;references:ID"`
}

func (c *CityOutput) TableName() string {
	return "cities"
}

type DistrictOutput struct {
	ID     uuid.UUID   `gorm:"primarykey" json:"id"`
	CityID uuid.UUID   `json:"cityID"`
	City   *CityOutput `json:"city,omitempty" gorm:"foreignKey:ID;references:CityID"`
	Name   string      `json:"name"`
}

func (c *DistrictOutput) TableName() string {
	return "districts"
}
