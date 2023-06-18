package verification

import (
	"context"

	"github.com/e-fish/api/pkg/domain/verification/model"
	"github.com/google/uuid"
)

type Repo interface {
	NewCommand(ctx context.Context) Command
	NewQuery() Query
}

type Command interface {
	CreateOTP(ctx context.Context, input model.CreateCodeOTPInput) (*uuid.UUID, error)
	DeleteOTP(ctx context.Context, input uuid.UUID) (*uuid.UUID, error)

	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
}

type Query interface {
	GetOTPByActivityAndUserID(ctx context.Context, input model.FindOTPInput) (*model.OutpuOTP, error)
	lock() Query
}
