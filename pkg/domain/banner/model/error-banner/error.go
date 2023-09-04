package errorBanner

import "github.com/e-fish/api/pkg/common/helper/werror"

var (
	ErrFindBanner = werror.Error{
		Code:    "FailedFindBanner",
		Message: "can't find banners",
	}
	ErrValidateInputBanner = werror.Error{
		Code:    "FailedValidateInputBanner",
		Message: "failed input banner",
	}

	ErrCreateBanner = werror.Error{
		Code:    "FailedCreateBanner",
		Message: "failed create banner",
	}

	ErrUpdateBanner = werror.Error{
		Code:    "FailedUpdateBanner",
		Message: "failed update banner",
	}

	ErrCommit = werror.Error{
		Code:    "FailedCommitTransaction",
		Message: "can't commit region transaction",
	}
	ErrRollback = werror.Error{
		Code:    "FailedRollbackTransaction",
		Message: "can't rollback region transaction",
	}
)
