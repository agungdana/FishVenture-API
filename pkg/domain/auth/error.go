package auth

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
)
