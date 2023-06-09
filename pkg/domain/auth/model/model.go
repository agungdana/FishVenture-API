package model

import (
	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID
	Name        string
	Email       string
	PhoneNumber string
	orm.OrmModel
}

type Role struct {
	ID             uuid.UUID
	Number         string
	Name           string
	Scope          string
	RolePermission []RolePermission
	orm.OrmModel
}

type Permission struct {
	ID             uuid.UUID
	Number         string
	Name           string
	RolePermission []RolePermission
	orm.OrmModel
}

type RolePermission struct {
	ID               uuid.UUID
	RoleID           uuid.UUID
	Role             Role
	PermissionID     uuid.UUID
	PermissionNumber string
	PermissionName   string
	Permission       Permission
	orm.OrmModel
}

type UserPermission struct {
	ID               uuid.UUID
	UserID           uuid.UUID
	User             User
	PermissionID     uuid.UUID
	PermissionNumber string
	PermissionName   string
	Permission       Permission
	orm.OrmModel
}
