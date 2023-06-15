package auth

import (
	"context"

	"github.com/e-fish/api/pkg/domain/auth/model"
	"github.com/google/uuid"
)

type Repo interface {
	NewCommand(ctx context.Context) Command
	NewQuery() Query
}

type Command interface {
	Login(ctx context.Context, input model.UserLoginInput) (*model.UserLoginOutput, error)
	LoginByGoogle(ctx context.Context, input model.UserLoginByGooleInput) (*model.UserLoginOutput, error)

	CreateUser(ctx context.Context, input model.CreateUserInput) (*uuid.UUID, error)
	UpdateUser(ctx context.Context, input model.UpdateUserInput) (*uuid.UUID, error)
	AddVerificationCode(ctx context.Context, input model.AddVerificationCodeInput) (*uuid.UUID, error)

	CreateUserRoleByRoleName(ctx context.Context, input model.AddUserRoleInput) (*uuid.UUID, error)
	CreateUserPermission(ctx context.Context, input model.AddUserPermissionInput) (*uuid.UUID, error)
	CreateRolePermission(ctx context.Context, input model.AddRolePermissionInput) (*uuid.UUID, error)
	DeleteRolePermission(ctx context.Context, input uuid.UUID) (*uuid.UUID, error)
	DeleteUserPermission(ctx context.Context, input uuid.UUID) (*uuid.UUID, error)

	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
}

type Query interface {
	GetProfile(ctx context.Context) (*model.Profile, error)
	GetUserByEmail(ctx context.Context, input string, withPermissionPreload bool) (*model.User, error)

	GetRoleByName(ctx context.Context, input string) (*model.Role, error)

	GetAllUserPermission(ctx context.Context) ([]*model.UserPermissionOutput, error)
	GetAllRolePermission(ctx context.Context) ([]*model.RolePermission, error)

	GetUserPermissionByCreated(ctx context.Context) ([]*model.UserPermissionOutput, error)
	GetUserRolePermissionIsNotCustomer(ctx context.Context) ([]*model.UserRoleOutput, error)

	lock() Query
}
