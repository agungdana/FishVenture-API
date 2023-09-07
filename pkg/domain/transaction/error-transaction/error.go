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

	ErrFoundOrder = werror.Error{
		Code:    "FailedFoundOrder",
		Message: "unable to find what you want",
	}

	ErrReadOrderData = werror.Error{
		Code:    "FailedReadOrderData",
		Message: "unable to read data",
	}

	ErrUpdateOrderStatus = werror.Error{
		Code:    "FailedUpdateOrderStatus",
		Message: "failed update order status",
	}
)
