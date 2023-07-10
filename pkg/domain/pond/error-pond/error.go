package errorpond

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
		Code:    "ValidatedFailedInputPond",
		Message: "field can't be empty",
	}

	ErrFailedUpdatePond = werror.Error{
		Code:    "FailedUpdatePond",
		Message: "failed update pond",
	}

	ErrValidateInputbBerkas = werror.Error{
		Code:    "ValidatedFailedInputBerkas",
		Message: "field can't be empty",
	}

	ErrValidateInputbUpdateStatus = werror.Error{
		Code:    "ValidatedFailedInputUpdateStatus",
		Message: "field can't be empty",
	}

	ErrFoundPond = werror.Error{
		Code:    "ValidatedFailedFoundPond",
		Message: "pond not found",
	}
	ErrFailedFindPond = werror.Error{
		Code:    "ValidatedFailedFoundPond",
		Message: "pond not found",
	}
	ErrCannotUpdateStatusPond = werror.Error{
		Code:    "CannotUpdateStatusPond",
		Message: "failed to update status",
	}
)
