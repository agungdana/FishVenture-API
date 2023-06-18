package model

import (
	"time"

	"github.com/e-fish/api/pkg/common/helper/werror"
	"github.com/e-fish/api/pkg/common/infra/orm"
	errorverification "github.com/e-fish/api/pkg/domain/verification/error_verification"
	"github.com/google/uuid"
)

type CreateCodeOTPInput struct {
	UserID   uuid.UUID `json:"user_id"`
	Activity string    `json:"activity"`
}

func (c *CreateCodeOTPInput) Validate() error {
	errs := werror.NewError("failed validate create code otp input")

	if c.UserID == uuid.Nil {
		errs.Add(errorverification.ErrFailedCreateCodeOTP.AttacthDetail(map[string]any{"UserID": "empty"}))
	}
	if c.Activity == "" {
		errs.Add(errorverification.ErrFailedCreateCodeOTP.AttacthDetail(map[string]any{"Activity": "empty"}))
	}
	return errs.Return()
}

func (c *CreateCodeOTPInput) ToOTP(userID uuid.UUID) OTP {
	return OTP{
		ID:       uuid.New(),
		UserID:   c.UserID,
		Code:     GenereatedCodeOTP(),
		ExpCode:  time.Now().Add(time.Minute * 2),
		Activity: c.Activity,
		OrmModel: orm.OrmModel{
			CreatedAt: time.Now(),
			CreatedBy: userID,
		},
	}
}

type FindOTPInput struct {
	UserID   uuid.UUID `json:"user_id"`
	Activity string    `json:"activity"`
}
