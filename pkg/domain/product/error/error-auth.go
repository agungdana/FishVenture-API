package errorauth

import "github.com/e-fish/api/pkg/common/helper/werror"

var (
	ErrCommit = werror.Error{
		Code:    "FailedCommitTransaction",
		Message: "can't commit transaction product",
	}
	ErrRollback = werror.Error{
		Code:    "FailedRollbackTransaction",
		Message: "can't rollback transaction product",
	}
)
