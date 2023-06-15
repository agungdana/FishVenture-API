package errorauth

import "github.com/e-fish/api/pkg/common/helper/werror"

var (
	ErrUserNotFound = werror.Error{
		Code:    "UserNotFound",
		Message: "user not found",
	}

	ErrUser = werror.Error{
		Code:    "UserQueryErr",
		Message: "internal server error",
	}

	ErrUserPasswordNotMatch = werror.Error{
		Code:    "UserQueryErr",
		Message: "internal server error",
	}

	ErrTokenError = werror.Error{
		Code:    "CreateTokenFailed",
		Message: "internal server error",
	}

	ErrSigninFirbaseAuth = werror.Error{
		Code:    "GauthFailed",
		Message: "internal server error",
	}
	ErrValidateCreateUserInput = werror.Error{
		Code:    "ValidateError",
		Message: "field can't be empty",
	}
	ErrHashedPassword = werror.Error{
		Code:    "FailedHashedPassword",
		Message: "internal server error, can't create user",
	}
	ErrUserAlreadyExist = werror.Error{
		Code:    "FailedCreateUser",
		Message: "user is already exist",
	}
	ErrFailedCreateUser = werror.Error{
		Code:    "FailedCreateUser",
		Message: "internal server error, can't create user",
	}
	ErrPhoneValidate = werror.Error{
		Code:    "FailedUpdateUser",
		Message: "phone number is invalid",
	}

	ErrRoleNotFound = werror.Error{
		Code:    "RoleNotFound",
		Message: "can't find role",
	}
	ErrRole = werror.Error{
		Code:    "RoleQueryErr",
		Message: "internal server error",
	}

	ErrCreateUserRole = werror.Error{
		Code:    "FailedCreateUserRole",
		Message: "internal server error",
	}

	ErrCommit = werror.Error{
		Code:    "FailedCommitTransaction",
		Message: "can't commit transaction",
	}
	ErrRollback = werror.Error{
		Code:    "FailedRollbackTransaction",
		Message: "can't rollback transaction",
	}
	ErrRolePermisionEmpty = werror.Error{
		Code:    "RolePermisionEmpty",
		Message: "can't find role permission",
	}
	ErrGetUserPermissionEmpty = werror.Error{
		Code:    "UserPermisionEmpty",
		Message: "can't find user permission",
	}

	ErrGetUserPermission = werror.Error{
		Code:    "UserPermissionErr",
		Message: "internal server error",
	}
)
