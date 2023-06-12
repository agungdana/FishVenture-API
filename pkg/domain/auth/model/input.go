package model

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/e-fish/api/pkg/common/helper/bcrypt"
	"github.com/e-fish/api/pkg/common/helper/rand"
	"github.com/e-fish/api/pkg/common/helper/werror"
	"github.com/e-fish/api/pkg/common/infra/orm"
	errorauth "github.com/e-fish/api/pkg/domain/auth/error"
	"github.com/google/uuid"
)

type CreateUserInput struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ApplicationType string `json:"application_type"`
	Status          bool   `json:"-"`
}

func (c *CreateUserInput) Validate() error {
	errs := werror.NewError("error validate input user")

	if c.Name == "" {
		errs.Add(errorauth.ErrValidateCreateUserInput.AttacthDetail(map[string]any{"name": "empty"}))
	}
	if c.Email == "" {
		errs.Add(errorauth.ErrValidateCreateUserInput.AttacthDetail(map[string]any{"email": "empty"}))
	}
	if c.Password == "" {
		errs.Add(errorauth.ErrValidateCreateUserInput.AttacthDetail(map[string]any{"password": "empty"}))
	}

	if c.ApplicationType == "" {
		errs.Add(errorauth.ErrValidateCreateUserInput.AttacthDetail(map[string]any{"application_type": "empty"}))
	}

	if err := errs.Return(); err != nil {
		return err
	}

	if c.ApplicationType == BUYER {
		c.Status = true
	}

	newPassword, err := bcrypt.HashPassowrd(c.Password)
	if err != nil {
		return errorauth.ErrHashedPassword.AttacthDetail(map[string]any{"errors": err})
	}

	c.Password = newPassword

	return nil
}

func (c *CreateUserInput) ToUser() User {
	var (
		userID = uuid.New()
	)

	return User{
		ID:       userID,
		Name:     c.Name,
		Email:    c.Email,
		Password: c.Password,
		Status:   c.Status,
		OrmModel: orm.OrmModel{
			CreatedAt: time.Now(),
			CreatedBy: userID,
		},
	}
}

type UpdateUserInput struct {
	Name  string `json:"name"`
	Photo string `json:"photo"`
	Phone string `json:"phone"`
}

func (u *UpdateUserInput) Validate() error {
	newPhone, err := ValidatePhone(u.Phone)
	if err != nil {
		return errorauth.ErrPhoneValidate.AttacthDetail(map[string]any{"phone": u.Phone})
	}

	u.Phone = newPhone
	return nil
}

func ValidatePhone(no string) (string, error) {
	code := "62"
	fc := no[0:1]

	no = strings.ReplaceAll(no, "+", "")
	no = strings.ReplaceAll(no, "-", "")
	no = strings.ReplaceAll(no, ",", "")
	no = strings.ReplaceAll(no, ".", "")
	no = strings.ReplaceAll(no, " ", "")

	if fc == "0" {
		no = strings.Replace(no, "0", code, 1)
	}

	_, err := strconv.Atoi(no)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return "", err
	}

	return no, nil
}

func (c *UpdateUserInput) ToUser(userID uuid.UUID) User {
	var (
		today = time.Now()
	)

	return User{
		ID:    userID,
		Name:  c.Name,
		Photo: c.Photo,
		Phone: c.Phone,
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
	UserID   uuid.UUID
	RoleName string
}

func (c *AddUserRoleInput) ToUserRole(roleID uuid.UUID) UserRole {
	return UserRole{
		ID:       uuid.New(),
		UserID:   c.UserID,
		RoleID:   roleID,
		OrmModel: orm.OrmModel{CreatedAt: time.Now(), CreatedBy: c.UserID},
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
