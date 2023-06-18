package errorverification

import "github.com/e-fish/api/pkg/common/helper/werror"

var (
	ErrFailedCreateCodeOTP = werror.Error{
		Code:    "FailedCreateCodeOTP",
		Message: "failed create code otp",
	}

	ErrFailedFindCodeOTP = werror.Error{
		Code:    "FailedFincCodeOTP",
		Message: "failed find code otp",
	}

	ErrExpiredCodeOTP = werror.Error{
		Code:    "CodeOTPExpired",
		Message: "otp code has expired",
	}

	ErrNotfoundCodeOTP = werror.Error{
		Code:    "CodeOTPNotFound",
		Message: "code otp not found",
	}

	ErrFailedDeleteCodeOTP = werror.Error{
		Code:    "FailedDeleteCodeOTP",
		Message: "failed delete code OTP",
	}

	ErrCommit = werror.Error{
		Code:    "FailedCommitTransaction",
		Message: "can't commit transaction product",
	}
	ErrRollback = werror.Error{
		Code:    "FailedRollbackTransaction",
		Message: "can't rollback transaction product",
	}
)
