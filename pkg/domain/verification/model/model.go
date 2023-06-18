package model

import (
	"time"

	"github.com/e-fish/api/pkg/common/helper/rand"
	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/google/uuid"
)

type OTP struct {
	ID       uuid.UUID `gorm:"size:256"`
	UserID   uuid.UUID
	Code     string
	ExpCode  time.Time
	Activity string
	orm.OrmModel
}

func GenereatedCodeOTP() string {
	return rand.GenereatedCodeOTP(6)
}
