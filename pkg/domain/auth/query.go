package auth

import (
	"context"
	"errors"

	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	errorauth "github.com/e-fish/api/pkg/domain/auth/error"
	"github.com/e-fish/api/pkg/domain/auth/model"
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

// GetProfile implements Query.
func (q *query) GetProfile(ctx context.Context) (*model.Profile, error) {
	var (
		data      = model.Profile{}
		db        = q.db
		userID, _ = ctxutil.GetUserID(ctx)
	)

	err := db.Where("deleted_at IS NULL and id = ?", userID).Take(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorauth.ErrUserNotFound.AttacthDetail(map[string]any{"id": userID})
		}
		return nil, errorauth.ErrUser.AttacthDetail(map[string]any{"error": err})
	}
	return &data, nil
}

// GetRoleByName implements Query.
func (q *query) GetRoleByName(ctx context.Context, input string) (*model.Role, error) {
	role := model.Role{}
	err := q.db.Where("deleted_at IS NULL and name = ?", input).Take(&role).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorauth.ErrRoleNotFound.AttacthDetail(map[string]any{"name": input})
		}
		return nil, errorauth.ErrRole.AttacthDetail(map[string]any{"errors": err})
	}

	return &role, nil
}

// GetAllUserPermission implements Query.
func (q *query) GetAllUserPermission(ctx context.Context) ([]*model.UserPermissionOutput, error) {
	data := []*model.UserPermissionOutput{}
	err := q.db.Find(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorauth.ErrGetUserPermissionEmpty
		}
		return nil, errorauth.ErrGetUserPermission.AttacthDetail(map[string]any{"error": err})
	}
	if len(data) < 1 {
		return nil, errorauth.ErrGetUserPermissionEmpty
	}
	return data, nil
}

// GetAllUserRole implements Query.
func (q *query) GetAllRolePermission(ctx context.Context) ([]*model.RolePermission, error) {
	data := []*model.RolePermission{}
	err := q.db.Find(&data).Error
	if err != nil {
		return nil, err
	}
	if len(data) < 1 {
		return nil, errorauth.ErrRolePermisionEmpty
	}
	return data, nil
}

// GetUserPermissionByCreated implements Query.
func (q *query) GetUserPermissionByCreated(ctx context.Context) ([]*model.UserPermissionOutput, error) {
	panic("unimplemented")
}

// GetUserRolePermissionIsNotCustomer implements Query.
func (q *query) GetUserRolePermissionIsNotCustomer(ctx context.Context) ([]*model.UserRoleOutput, error) {
	panic("unimplemented")
}

// GetUserByEmail implements Query.
func (q *query) GetUserByEmail(ctx context.Context, input string, withPermissionPreload bool) (*model.User, error) {
	var (
		data = model.User{}
		db   = q.db
	)

	if withPermissionPreload {
		//get data role by user exist
		db = db.Preload("UserRole.Role")
	}

	err := db.Where("deleted_at IS NULL and email = ?", input).Take(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorauth.ErrUserNotFound.AttacthDetail(map[string]any{"email": input})
		}
		return nil, errorauth.ErrUser.AttacthDetail(map[string]any{"error": err})
	}
	return &data, nil
}

// lock implements Query.
// lock table row to avoid race condition
func (q *query) lock() Query {
	db := q.db.Clauses(clause.Locking{Strength: "UPDATE"})
	return &query{db: db}
}
