package model

import (
	"time"

	"github.com/e-fish/api/pkg/domain/pond/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BudidayaOutput struct {
	ID              uuid.UUID          `gorm:"pricmaryKey,size:256" json:"id"`
	PondID          uuid.UUID          `gorm:"size:256" json:"pondID"`
	Pond            *model.PondOutput  `gorm:"foreignKey:PondID;references:ID" json:"pond,omitempty"`
	PoolID          uuid.UUID          `gorm:"size:256" json:"poolID"`
	Pool            *model.PoolOutput  `gorm:"foreignKey:PoolID;references:ID" json:"pool,omitempty"`
	DateOfSeed      time.Time          `json:"dateOfSeed"`
	FishSpeciesID   uuid.UUID          `json:"fishSpeciesID"`
	FishSpecies     *FishSpeciesOutput `gorm:"foreignKey:FishSpeciesID;references:ID" json:"fishSpecies,omitempty"`
	FishSpeciesName string             `json:"fishSpeciesName,omitempty"`
	EstTonase       float64            `json:"estTonase"`
	EstPanenDate    *time.Time         `json:"estPanenDate,omitempty"`
	EstPrice        int                `json:"estPrice"`
	Status          string             `json:"status"`
	PriceList       []*PriceListOutput `gorm:"foreignKey:BudidayaID;references:ID" json:"priceList,omitempty"`

	Sold  int `json:"sold"`
	Stock int `json:"stock"`
}

func (p *BudidayaOutput) TableName() string {
	return "budidayas"
}

func (p *BudidayaOutput) AfterFind(db *gorm.DB) (err error) {
	p.Stock = int(p.EstTonase) - p.Sold
	return
}

type PriceListOutput struct {
	ID         uuid.UUID       `gorm:"primaryKey,size:256" json:"id,omitempty"`
	BudidayaID *uuid.UUID      `json:"budidayaID,omitempty"`
	Budidaya   *BudidayaOutput `gorm:"foreignKey:BudidayaID;references:ID" json:"budidaya,omitempty"`
	Limit      int             `json:"limit,omitempty"`
	Price      int             `json:"price,omitempty"`
}

func (p *PriceListOutput) TableName() string {
	return "price_lists"
}

type FishSpeciesOutput struct {
	ID       uuid.UUID         `gorm:"primaryKey,size:256" json:"id"`
	Name     string            `json:"name"`
	Asal     string            `json:"asal"`
	Budidaya []*BudidayaOutput `gorm:"foreignKey:FishSpeciesID;references:ID" json:"budidaya,omitempty"`
}

func (p *FishSpeciesOutput) TableName() string {
	return "fish_species"
}
