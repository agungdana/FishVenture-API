package errortransaction

import "github.com/e-fish/api/pkg/common/helper/werror"

var (
	ErrCommit = werror.Error{
		Code:    "FailedCommitTransaction",
		Message: "can't commit transaction",
	}
	ErrRollback = werror.Error{
		Code:    "FailedRollbackTransaction",
		Message: "can't rollback transaction",
	}

	ErrValidateCreateInput = werror.Error{
		Code:    "FailedValidateCreateOrderInput",
		Message: "field can't by empty",
	}

	ErrCreateOrder = werror.Error{
		Code:    "FailedCreateOrder",
		Message: "failed create order",
	}
)
