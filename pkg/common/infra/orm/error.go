package orm

import "github.com/e-fish/api/pkg/common/helper/werror"

var (
	ErrDriverNotSupported = werror.Error{
		Code:    "DriverNotSupported",
		Message: "driver not suported",
	}

	ErrCreateConnection = werror.Error{
		Code:    "ConnectionFailed",
		Message: "failed to establish a new connection to the database",
	}

	WithoutTransaction = werror.Error{
		Code:    "TransactionNotFound",
		Message: "can't found transaction in ctx",
	}
)
