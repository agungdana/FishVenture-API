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

	ErrValidateInputBudidaya = werror.Error{
		Code:    "ValidatedFailedInputBudidaya",
		Message: "field can't be empty",
	}

	ErrValidateMultipleInputPriceList = werror.Error{
		Code:    "ValidatedFailedMultipleInputPriceList",
		Message: "field can't be empty",
	}

	ErrValidateInputPriceList = werror.Error{
		Code:    "ValidatedFailedInputPriceList",
		Message: "field can't be empty",
	}

	ErrValidateInputFishSpecies = werror.Error{
		Code:    "ValidatedFailedInputFishSpecies",
		Message: "field can't be empty",
	}

	ErrFailedCreateBudidayaExist = werror.Error{
		Code:    "FailedCreateBudidayaExist",
		Message: "budidaya exist",
	}

	ErrFailedCreateBudidaya = werror.Error{
		Code:    "FailedCreateBudidaya",
		Message: "failed create Budidaya",
	}

	ErrFailedUpdateBudidaya = werror.Error{
		Code:    "FailedUpdateBudidaya",
		Message: "failed update Budidaya",
	}

	ErrValidateInputbUpdateStatus = werror.Error{
		Code:    "ValidatedFailedInputUpdateStatus",
		Message: "field can't be empty",
	}

	ErrFoundBudidaya = werror.Error{
		Code:    "ValidatedFailedFoundBudidaya",
		Message: "Budidaya not found",
	}

	ErrFailedReadBudidaya = werror.Error{
		Code:    "ValidatedReadBudidaya",
		Message: "failed read budidaya data",
	}

	ErrCannotUpdateStatusBudidaya = werror.Error{
		Code:    "CannotUpdateStatusBudidaya",
		Message: "failed to update status",
	}

	ErrUnsuportedFindBudidaya = werror.Error{
		Code:    "FailedFindBudidaya",
		Message: "failed find Budidaya unsuported type",
	}

	ErrFoundPricelist = werror.Error{
		Code:    "FailedFindPricelistBudidaya",
		Message: "failed find pricelist budidaya",
	}

	ErrReadPricelistData = werror.Error{
		Code:    "FailedReadPricelistData",
		Message: "failed read pricelist data",
	}
)
