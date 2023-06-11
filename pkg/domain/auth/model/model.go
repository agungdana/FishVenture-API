package model

import (
	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/google/uuid"
)

type User struct {
	ID                  uuid.UUID
	Name                string
	Email               string
	Password            string
	Phone               string
	Photo               string
	Status              bool
	VarificationCode    string
	ExpVerificationCode string
	UserRole            []*UserRole
	UserPermission      []*UserPermission
	orm.OrmModel
}

type Role struct {
	ID             uuid.UUID
	Code           string
	Name           string
	Scope          string
	RolePermission []*RolePermission
	orm.OrmModel
}

type Permission struct {
	ID             uuid.UUID
	Code           string
	Name           string
	Path           string
	RolePermission []*RolePermission
	orm.OrmModel
}

type RolePermission struct {
	ID             uuid.UUID
	RoleID         uuid.UUID
	Role           Role
	PermissionID   uuid.UUID
	PermissionName string
	PermissionPath string
	Permission     Permission
	orm.OrmModel
}

type UserRole struct {
	ID     uuid.UUID
	UserID uuid.UUID
	User   User
	RoleID uuid.UUID
	Role   Role
	orm.OrmModel
}

type UserPermission struct {
	ID             uuid.UUID
	UserID         uuid.UUID
	User           User
	PermissionID   uuid.UUID
	PermissionPath string
	PermissionName string
	Permission     Permission
	orm.OrmModel
}
