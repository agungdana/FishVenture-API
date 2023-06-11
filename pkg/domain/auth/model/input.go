package model

import (
	"strings"
	"time"

	"github.com/e-fish/api/pkg/common/helper/rand"
	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/google/uuid"
)

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Photo    string `json:"photo"`
}

func (c *CreateUserInput) ToUser() User {
	var (
		userID = uuid.New()
	)

	return User{
		ID:       userID,
		Name:     c.Name,
		Email:    c.Email,
		Phone:    c.Phone,
		Password: c.Password,
		OrmModel: orm.OrmModel{
			CreatedAt: time.Now(),
			CreatedBy: userID,
		},
	}
}

type UpdateUserInput struct {
	Name  string `json:"name"`
	Photo string `json:"photo"`
}

func (c *UpdateUserInput) ToUser(userID uuid.UUID) User {
	var (
		today = time.Now()
	)

	return User{
		ID:    userID,
		Name:  c.Name,
		Photo: c.Photo,
		OrmModel: orm.OrmModel{
			UpdatedAt: &today,
			UpdatedBy: &userID,
		},
	}
}

type UserLoginInput struct {
	Email           string
	Password        string
	ApplicationType string
}

type UserLoginByGooleInput struct {
	Token           string
	FullName        string
	ApplicationType string
}

type AddVerificationCodeInput struct {
	UserID           uuid.UUID
	VerificationCode string
}

type AddRolePermissionInput struct {
	RoleID       uuid.UUID
	PermissionID uuid.UUID
}

func (c *AddRolePermissionInput) ToRolePermission(userID uuid.UUID, permissionPath, permissionName string) RolePermission {
	return RolePermission{
		ID:             uuid.New(),
		RoleID:         c.RoleID,
		PermissionID:   c.PermissionID,
		PermissionName: permissionName,
		PermissionPath: permissionPath,
		OrmModel: orm.OrmModel{
			CreatedAt: time.Now(),
			CreatedBy: userID,
		},
	}
}

type AddUserPermissionInput struct {
	UserID       uuid.UUID
	PermissionID uuid.UUID
}

func (c *AddUserPermissionInput) ToUserPermission(userID uuid.UUID, permissionPath, permissionName string) UserPermission {
	return UserPermission{
		ID:             uuid.New(),
		UserID:         c.UserID,
		PermissionID:   c.PermissionID,
		PermissionName: permissionName,
		PermissionPath: permissionPath,
		OrmModel: orm.OrmModel{
			CreatedAt: time.Now(),
			CreatedBy: userID,
		},
	}
}

type AddUserRoleInput struct {
	UserID uuid.UUID
	RoleID uuid.UUID
}

func (c *AddUserRoleInput) ToUserPermission(userID uuid.UUID, permissionPath, permissionName string) UserRole {
	return UserRole{
		ID:       uuid.New(),
		UserID:   c.UserID,
		RoleID:   c.RoleID,
		OrmModel: orm.OrmModel{CreatedAt: time.Now(), CreatedBy: userID},
	}
}

type CreateRoleInput struct {
	Name  string `json:"name"`
	Scope string `json:"scope"`
}

func (c *CreateRoleInput) ToRole(userID uuid.UUID) Role {
	return Role{
		ID:    uuid.New(),
		Code:  GenereatedRandCode("RC"),
		Name:  c.Name,
		Scope: c.Scope,
		OrmModel: orm.OrmModel{
			CreatedAt: time.Now(),
			CreatedBy: userID,
		},
	}
}

type CreatePermissionInput struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func (c *CreatePermissionInput) ToPermission(userID uuid.UUID) Permission {
	return Permission{
		ID:       uuid.New(),
		Code:     GenereatedRandCode("PC"),
		Name:     c.Name,
		Path:     c.Path,
		OrmModel: orm.OrmModel{CreatedAt: time.Now(), CreatedBy: userID},
	}
}

func GenereatedRandCode(prefix string) string {
	prefix = strings.ToUpper(prefix)
	return prefix + rand.RandCode(6)
}
