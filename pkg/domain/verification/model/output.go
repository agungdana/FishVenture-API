package model

import (
	"time"

	"github.com/google/uuid"
)

type OutpuOTP struct {
	ID        uuid.UUID
	Code      string
	ExpCode   time.Time
	Activity  string
	CreatedAt time.Time
}

func (o *OutpuOTP) TableName() string {
	return "otp"
}
