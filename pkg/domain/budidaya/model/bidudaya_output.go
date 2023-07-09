package model

import (
	"time"

	"github.com/e-fish/api/pkg/domain/pond/model"
	"github.com/google/uuid"
)

type BudidayaOutput struct {
	ID              uuid.UUID          `gorm:"primaryKey,size:256" json:"id"`
	PoolID          uuid.UUID          `gorm:"size:256" json:"poolID"`
	Pool            model.Pool         `json:"pool"`
	DateOfSeed      time.Time          `json:"dateOfSeed"`
	FishSpeciesID   uuid.UUID          `json:"fishSpeciesID"`
	FishSpecies     FishSpecies        `gorm:"foreignKey:FishSpeciesID;references:ID" json:"fishSpecies"`
	FishSpeciesName string             `json:"fishSpeciesName"`
	EstTonase       float64            `json:"estTonase"`
	EstPanenDate    time.Time          `json:"estPanenDate"`
	EstPrice        int                `json:"estPrice"`
	Status          string             `json:"status"`
	PriceList       []*PriceListOutput `gorm:"foreignKey:BudidayaID;references:ID" json:"priceList"`
}

type PriceListOutput struct {
	ID         uuid.UUID      `gorm:"primaryKey,size:256" json:"id,omitempty"`
	BudidayaID uuid.UUID      `json:"budidayaID,omitempty"`
	Budidaya   BudidayaOutput `gorm:"foreignKey:BudidayaID;references:ID" json:"budidaya,omitempty"`
	Limit      int            `json:"limit,omitempty"`
	Price      int            `json:"price,omitempty"`
}

type FishSpeciesOutput struct {
	ID       uuid.UUID   `gorm:"primaryKey,size:256" json:"id"`
	Name     string      `json:"name"`
	Asal     string      `json:"asal"`
	Budidaya []*Budidaya `gorm:"foreignKey:FishSpeciesID;references:ID" json:"budidaya,omitempty"`
}
