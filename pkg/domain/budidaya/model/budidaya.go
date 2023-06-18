package model

import (
	"time"

	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/google/uuid"
)

type Budidaya struct {
	ID              uuid.UUID `gorm:"size:256"`
	PoolID          uuid.UUID `gorm:"size:256"`
	Pool            Pool
	DateOfSeed      time.Time
	FishSpeciesID   uuid.UUID
	FishSpecies     FishSpecies
	FishSpeciesName string
	EstTonase       float64
	Status          bool
	orm.OrmModel
}

type FishSpecies struct {
	ID       uuid.UUID
	Name     string
	Budidaya []*Budidaya
	orm.OrmModel
}
