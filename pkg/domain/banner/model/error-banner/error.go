package errorBanner

import "github.com/e-fish/api/pkg/common/helper/werror"

var (
	ErrFindBanner = werror.Error{
		Code:    "FailedFindBanner",
		Message: "can't find banners",
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
