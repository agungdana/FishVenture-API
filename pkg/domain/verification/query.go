package verification

import (
	"context"
	"errors"
	"time"

	errorverification "github.com/e-fish/api/pkg/domain/verification/error_verification"
	"github.com/e-fish/api/pkg/domain/verification/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func newQuery(db *gorm.DB) Query {
	return &query{
		db: db,
	}
}

type query struct {
	db *gorm.DB
}

// GetRandCodeByActivityAndUserID implements Query.
func (c *query) GetOTPByActivityAndUserID(ctx context.Context, input model.FindOTPInput) (*model.OutpuOTP, error) {
	var (
		otp = model.OutpuOTP{}
	)
	err := c.db.Where("deleted_at IS NULL AND user_id = ? AND activity = ?", input.UserID, input.Activity).Take(&otp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorverification.ErrNotfoundCodeOTP.AttacthDetail(map[string]any{"uid": input.UserID, "activity": input.Activity})
		}
		return nil, errorverification.ErrFailedFindCodeOTP.AttacthDetail(map[string]any{"error": err})
	}

	if time.Now().After(otp.ExpCode) {
		return nil, errorverification.ErrExpiredCodeOTP.AttacthDetail(map[string]any{"otp-id": otp.ID, "exp-date": otp.ExpCode})
	}

	return &otp, nil
}

// lock implements Query.
// lock table row to avoid race condition
func (q *query) lock() Query {
	db := q.db.Clauses(clause.Locking{Strength: "UPDATE"})
	return &query{db: db}
}
