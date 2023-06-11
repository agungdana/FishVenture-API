package auth

import (
	"context"

	"github.com/e-fish/api/pkg/common/infra/firebase"
	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/e-fish/api/pkg/common/infra/token"
	"github.com/e-fish/api/pkg/domain/auth/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func newCommand(ctx context.Context, db *gorm.DB, maker token.Token, firebase firebase.GoogleAuth) Command {
	var (
		dbTxn = orm.BeginTxn(ctx, db)
	)

	return &command{
		dbTxn: dbTxn,
		query: newQuery(dbTxn, maker, firebase),
	}
}

type command struct {
	dbTxn *gorm.DB
	query Query
}

// AddVerificationCode implements Command.
func (c *command) AddVerificationCode(ctx context.Context, input model.AddVerificationCodeInput) (*uuid.UUID, error) {
	panic("unimplemented")
}

// CreateRolePermission implements Command.
func (c *command) CreateRolePermission(ctx context.Context, input model.AddRolePermissionInput) (*uuid.UUID, error) {
	panic("unimplemented")
}

// CreateUser implements Command.
func (c *command) CreateUser(ctx context.Context, input model.CreateUserInput) (*uuid.UUID, error) {
	panic("unimplemented")
}

// CreateUserPermission implements Command.
func (c *command) CreateUserPermission(ctx context.Context, input model.AddUserPermissionInput) (*uuid.UUID, error) {
	panic("unimplemented")
}

// DeleteRolePermission implements Command.
func (c *command) DeleteRolePermission(ctx context.Context, input uuid.UUID) (*uuid.UUID, error) {
	panic("unimplemented")
}

// DeleteUserPermission implements Command.
func (c *command) DeleteUserPermission(ctx context.Context, input uuid.UUID) (*uuid.UUID, error) {
	panic("unimplemented")
}

// UpdateUser implements Command.
func (c *command) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*uuid.UUID, error) {
	panic("unimplemented")
}

// Commit implements Command.
func (c *command) Commit(ctx context.Context) error {
	if err := orm.CommitTxn(ctx); err != nil {
		return err
	}
	return nil
}

// Rollback implements Command.
func (c *command) Rollback(ctx context.Context) error {
	if err := orm.RollbackTxn(ctx); err != nil {
		return err
	}
	return nil
}
