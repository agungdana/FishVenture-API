package model

import (
	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/google/uuid"
)

type User struct {
	ID                  uuid.UUID `gorm:"size:256"`
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
	ID             uuid.UUID `gorm:"size:256"`
	Code           string
	Name           string
	Scope          string
	RolePermission []*RolePermission
	orm.OrmModel
}

type Permission struct {
	ID             uuid.UUID `gorm:"size:256"`
	Code           string
	Name           string
	Path           string
	RolePermission []*RolePermission
	orm.OrmModel
}

type RolePermission struct {
	ID             uuid.UUID `gorm:"size:256"`
	RoleID         uuid.UUID `gorm:"size:256"`
	Role           Role
	PermissionID   uuid.UUID `gorm:"size:256"`
	PermissionName string
	PermissionPath string
	Permission     Permission
	orm.OrmModel
}

type UserRole struct {
	ID     uuid.UUID `gorm:"size:256"`
	UserID uuid.UUID `gorm:"size:256"`
	User   User
	RoleID uuid.UUID `gorm:"size:256"`
	Role   Role
	orm.OrmModel
}

type UserPermission struct {
	ID             uuid.UUID `gorm:"size:256"`
	UserID         uuid.UUID `gorm:"size:256"`
	User           User
	PermissionID   uuid.UUID `gorm:"size:256"`
	PermissionPath string
	PermissionName string
	Permission     Permission
	orm.OrmModel
}
