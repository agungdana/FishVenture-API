package model

import (
	"time"

	"github.com/google/uuid"
)

type UserLoginOutput struct {
	Token string `json:"token"`
}

type Profile struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Photo string `json:"photo"`
}

func (p *Profile) TableName() string {
	return "users"
}

type UserPermissionOutput struct {
	ID             uuid.UUID  `json:"id"`
	RoleID         uuid.UUID  `json:"role_id"`
	Role           Role       `json:"role"`
	PermissionID   uuid.UUID  `json:"permission_id"`
	PermissionName string     `json:"permission_name"`
	PermissionPath string     `json:"permission_path"`
	Permission     Permission `json:"permission"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

func (u *UserPermissionOutput) TableName() string {
	return "user_permissions"
}

type RolePermissionOutput struct {
	ID             uuid.UUID  `json:"id"`
	UserID         uuid.UUID  `json:"user_id"`
	User           User       `json:"user"`
	PermissionID   uuid.UUID  `json:"permission_id"`
	PermissionPath string     `json:"permission_path"`
	PermissionName string     `json:"permission_name"`
	Permission     Permission `json:"permission"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

func (u *RolePermissionOutput) TableName() string {
	return "role_permissions"
}

type RoleOutput struct {
	ID             uuid.UUID               `json:"id"`
	Code           string                  `json:"code"`
	Name           string                  `json:"name"`
	Scope          string                  `json:"scope"`
	RolePermission []*RolePermissionOutput `json:"role_permission,omitempty"`
}

func (u *RoleOutput) TableName() string {
	return "roles"
}

type UserRoleOutput struct {
	ID     uuid.UUID  `json:"id"`
	UserID uuid.UUID  `json:"user_id"`
	User   User       `json:"user"`
	RoleID uuid.UUID  `json:"role_id"`
	Role   RoleOutput `json:"role,omitempty"`
}

func (u *UserRoleOutput) TableName() string {
	return "user_roles"
}
