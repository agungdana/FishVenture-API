package config

import "github.com/e-fish/api/pkg/common/helper/werror"

var (
	ErrLoadEnv = werror.Error{
		Code:    "ErrLoadENV",
		Message: "unable to load env",
	}
	ErrEmptyConfig = werror.Error{
		Code:    "ErrConfigEmpty",
		Message: "config is empty",
	}
)
