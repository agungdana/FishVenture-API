package model

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/e-fish/api/pkg/common/helper/werror"
	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/e-fish/api/pkg/domain/pond/model"
	"github.com/google/uuid"
)

type Budidaya struct {
	ID              uuid.UUID `gorm:"primaryKey,size:256"`
	Code            string
	PondID          uuid.UUID `gorm:"size:256"`
	Pond            model.Pond
	PoolID          uuid.UUID `gorm:"size:256"`
	Pool            model.Pool
	DateOfSeed      time.Time
	FishSpeciesID   uuid.UUID
	FishSpecies     FishSpecies
	FishSpeciesName string
	EstTonase       float64
	EstPanenDate    *time.Time
	EstPrice        int
	Status          string
	Sold            int
	PriceList       []*PriceList
	orm.OrmModel
}

// PondName/Years/int
func GeneratedCodeBudidaya(pondName, last string) (string, error) {
	var (
		today = time.Now()
	)

	if last == "" {
		newCode := fmt.Sprintf("%v/%v/%v", pondName, today.Year(), 1)
		return newCode, nil
	}

	listCode := strings.Split(last, "/")
	if len(listCode) != 3 {
		return "", werror.Error{
			Code:    "code invalid",
			Message: "invalid exist code",
		}
	}

	lastString := listCode[2]
	lastInt, err := strconv.Atoi(lastString)
	if err != nil {
		return "", werror.Error{
			Code:    "code invalid",
			Message: "invalid exist code",
		}
	}

	newCode := fmt.Sprintf("%v/%v/%v", pondName, today.Year(), lastInt+1)

	return newCode, nil
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
