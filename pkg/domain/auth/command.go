package auth

import (
	"context"
	"time"

	"github.com/e-fish/api/pkg/common/helper/bcrypt"
	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/common/infra/firebase"
	"github.com/e-fish/api/pkg/common/infra/orm"
	"github.com/e-fish/api/pkg/common/infra/token"
	errorauth "github.com/e-fish/api/pkg/domain/auth/error"
	"github.com/e-fish/api/pkg/domain/auth/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func newCommand(ctx context.Context, db *gorm.DB, maker token.Token, gauth firebase.GoogleAuth) Command {
	var (
		dbTxn = orm.BeginTxn(ctx, db)
	)

	return &command{
		dbTxn:      dbTxn,
		tokenMaker: maker,
		gauth:      gauth,
		query:      newQuery(dbTxn),
	}
}

type command struct {
	dbTxn      *gorm.DB
	tokenMaker token.Token
	gauth      firebase.GoogleAuth
	query      Query
}

// CreateUserRoleByRoleName implements Command.
func (c *command) CreateUserRoleByRoleName(ctx context.Context, input model.AddUserRoleInput) (*uuid.UUID, error) {

	role, err := c.query.GetRoleByName(ctx, input.RoleName)
	if err != nil {
		return nil, err
	}

	newUserRole := input.ToUserRole(role.ID)

	err = c.dbTxn.WithContext(ctx).Create(&newUserRole).Error
	if err != nil {
		return nil, errorauth.ErrCreateUserRole.AttacthDetail(map[string]any{"errors": err})
	}

	return &newUserRole.ID, nil
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
	logger.DebugWithContext(ctx, "#### proces create user")

	exist, err := c.query.GetUserByEmail(ctx, input.Email, false)
	if err != nil {
		if !errorauth.ErrUserNotFound.Is(err) {
			return nil, err
		}
	}

	if exist != nil {
		return nil, errorauth.ErrUserAlreadyExist.AttacthDetail(map[string]any{"email": input.Email})
	}

	err = input.Validate()
	if err != nil {
		return nil, err
	}

	newUser := input.ToUser()

	if input.ApplicationType == model.SELLER {
		pondID := uuid.New()
		newUser.PondID = &pondID
	}

	err = c.dbTxn.WithContext(ctx).Create(&newUser).Error
	if err != nil {
		return nil, errorauth.ErrFailedCreateUser.AttacthDetail(map[string]any{"err": err})
	}

	return &newUser.ID, nil
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
	userID, _ := ctxutil.GetUserID(ctx)

	err := input.Validate()
	if err != nil {
		return nil, err
	}

	updateUser := input.ToUser(userID)

	err = c.dbTxn.WithContext(ctx).Updates(&updateUser).Error
	if err != nil {
		return nil, err
	}

	return &userID, nil
}

// Login implements Command.
func (c *command) Login(ctx context.Context, input model.UserLoginInput) (*model.UserLoginOutput, error) {
	user, err := c.query.GetUserByEmail(ctx, input.Email, true)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.ComparePassword(input.Password, user.Password); err != nil {
		return nil, errorauth.ErrUserPasswordNotMatch.AttacthDetail(map[string]any{"input-pw": input.Password, "email": input.Email, "err": err})
	}

	role := []uuid.UUID{}

	for _, v := range user.UserRole {
		role = append(role, v.RoleID)
		if input.ApplicationType != v.Role.Name {
			return nil, errorauth.ErrUserAccess.AttacthDetail(map[string]any{"app-type": input.ApplicationType})
		}
	}

	if len(role) < 1 {
		return nil, errorauth.ErrUserAccess.AttacthDetail(map[string]any{"app-type": input.ApplicationType})
	}

	payload := token.Payload{
		UserID:    user.ID,
		UserRole:  role,
		AppType:   input.ApplicationType,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().AddDate(1, 0, 0),
	}

	if user.PondID != nil {
		payload.PondID = *user.PondID
	}

	token, err := c.tokenMaker.CreateToken(&payload)

	if err != nil {
		return nil, errorauth.ErrTokenError.AttacthDetail(map[string]any{"error": err})
	}

	return &model.UserLoginOutput{
		ApplicationType: input.ApplicationType,
		Token:           token,
	}, nil
}

// LoginByGoogle implements Command.
func (c *command) LoginByGoogle(ctx context.Context, input model.UserLoginByGooleInput) (*model.UserLoginOutput, error) {

	signin, err := c.gauth.Signin(ctx, input.Token)
	if err != nil {
		return nil, errorauth.ErrSigninFirbaseAuth.AttacthDetail(map[string]any{"err": err})
	}

	user, err := c.query.GetUserByEmail(ctx, signin.Email, true)
	if err != nil {
		if !errorauth.ErrUserNotFound.Is(err) {
			return nil, err
		}
	}

	role := []uuid.UUID{}

	for _, v := range user.UserRole {
		role = append(role, v.RoleID)
	}

	token, err := c.tokenMaker.CreateToken(&token.Payload{
		UserID:    user.ID,
		UserRole:  role,
		AppType:   input.ApplicationType,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().AddDate(1, 0, 0),
	})
	if err != nil {
		return nil, errorauth.ErrTokenError.AttacthDetail(map[string]any{"error": err})
	}

	return &model.UserLoginOutput{
		ApplicationType: input.ApplicationType,
		Token:           token,
	}, nil
}

func (c *command) UpdateUserStatusAndPondID(ctx context.Context, input uuid.UUID) (*uuid.UUID, error) {
	var (
		userID, _ = ctxutil.GetUserID(ctx)
		today     = time.Now()
	)

	user := model.User{
		ID:     userID,
		PondID: &input,
		OrmModel: orm.OrmModel{
			UpdatedAt: &today,
			UpdatedBy: &userID,
		},
	}

	err := c.dbTxn.Where("deleted_at is NULL and id = ?", userID).Updates(&user).Error
	if err != nil {
		return nil, errorauth.ErrUpdateUser.AttacthDetail(map[string]any{"error": err})
	}

	return &userID, nil
}

// Commit implements Command.
func (c *command) Commit(ctx context.Context) error {
	if err := orm.CommitTxn(ctx); err != nil {
		return errorauth.ErrCommit.AttacthDetail(map[string]any{"errors": err})
	}
	return nil
}

// Rollback implements Command.
func (c *command) Rollback(ctx context.Context) error {
	if err := orm.RollbackTxn(ctx); err != nil {
		return errorauth.ErrRollback.AttacthDetail(map[string]any{"errors": err})
	}
	return nil
}
