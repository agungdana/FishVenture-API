package auth

import (
	"context"
	"errors"
	"time"

	"github.com/e-fish/api/pkg/common/helper/bcrypt"
	"github.com/e-fish/api/pkg/common/infra/firebase"
	"github.com/e-fish/api/pkg/common/infra/token"
	"github.com/e-fish/api/pkg/domain/auth/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func newQuery(db *gorm.DB, maker token.Token, gauth firebase.GoogleAuth) Query {
	return &query{
		tokenMaker: maker,
		gauth:      gauth,
		db:         db,
	}
}

type query struct {
	tokenMaker token.Token
	gauth      firebase.GoogleAuth
	db         *gorm.DB
}

// GetAllUserPermission implements Query.
func (q *query) GetAllUserPermission(ctx context.Context) ([]*model.UserPermissionOutput, error) {
	panic("unimplemented")
}

// GetAllUserRole implements Query.
func (q *query) GetAllUserRole(ctx context.Context) ([]*model.UserRole, error) {
	panic("unimplemented")
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
		db.Preload("UserRole")
	}

	err := db.Where("deleted_at IS NULL and email = ?", input).Take(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound.AttacthDetail(map[string]any{"email": input})
		}
		return nil, ErrUser.AttacthDetail(map[string]any{"error": err})
	}
	return &data, nil
}

// Login implements Query.
func (q *query) Login(ctx context.Context, input model.UserLoginInput) (*model.UserLoginOutput, error) {
	user, err := q.GetUserByEmail(ctx, input.Email, true)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.ComparePassword(input.Password, user.Password); err != nil {
		return nil, ErrUserPasswordNotMatch.AttacthDetail(map[string]any{"input-pw": input.Password, "email": input.Email, "err": err})
	}

	role := []uuid.UUID{}

	for _, v := range user.UserRole {
		role = append(role, v.RoleID)
	}

	token, err := q.tokenMaker.CreateToken(&token.Payload{
		UserID:    user.ID,
		UserRole:  role,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().AddDate(1, 0, 0),
	})
	if err != nil {
		return nil, ErrTokenError.AttacthDetail(map[string]any{"error": err})
	}

	return &model.UserLoginOutput{
		Token: token,
	}, nil
}

// LoginByGoogle implements Query.
func (q *query) LoginByGoogle(ctx context.Context, input model.UserLoginByGooleInput) (*model.UserLoginOutput, error) {

	signin, err := q.gauth.Signin(ctx, input.Token)
	if err != nil {
		return nil, ErrSigninFirbaseAuth.AttacthDetail(map[string]any{"err": err})
	}

	user, err := q.GetUserByEmail(ctx, signin.Email, true)
	if err != nil {
		if !ErrUserNotFound.Is(err) {
			return nil, err
		}

	}

	role := []uuid.UUID{}

	for _, v := range user.UserRole {
		role = append(role, v.RoleID)
	}

	token, err := q.tokenMaker.CreateToken(&token.Payload{
		UserID:    user.ID,
		UserRole:  role,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().AddDate(1, 0, 0),
	})
	if err != nil {
		return nil, ErrTokenError.AttacthDetail(map[string]any{"error": err})
	}

	return &model.UserLoginOutput{
		Token: token,
	}, nil
}

// lock implements Query.
// lock table row to avoid race condition
func (q *query) lock() Query {
	db := q.db.Clauses(clause.Locking{Strength: "UPDATE"})
	return &query{db: db, tokenMaker: q.tokenMaker}
}
