package errorbudidaya

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

	ErrValidateInputPond = werror.Error{
		Code:    "ValidatedFailed",
		Message: "field can't be empty",
	}

	ErrValidateInputbBerkas = werror.Error{
		Code:    "ValidatedFailed",
		Message: "field can't be empty",
	}
)
