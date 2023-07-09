package model

import (
	"time"

	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/e-fish/api/pkg/domain/pond/model"
	"github.com/google/uuid"
)

type Budidaya struct {
	ID              uuid.UUID `gorm:"primaryKey,size:256"`
	PoolID          uuid.UUID `gorm:"size:256"`
	Pool            model.Pool
	DateOfSeed      time.Time
	FishSpeciesID   uuid.UUID
	FishSpecies     FishSpecies
	FishSpeciesName string
	EstTonase       float64
	EstPanenDate    time.Time
	EstPrice        int
	Status          string
	PriceList       []*PriceList
	orm.OrmModel
}

type PriceList struct {
	ID         uuid.UUID `gorm:"primaryKey,size:256"`
	BudidayaID uuid.UUID
	Budidaya   Budidaya
	Limit      int
	Price      int
	orm.OrmModel
}

type FishSpecies struct {
	ID       uuid.UUID `gorm:"primaryKey,size:256"`
	Name     string
	Asal     string
	Budidaya []*Budidaya
	orm.OrmModel
}
