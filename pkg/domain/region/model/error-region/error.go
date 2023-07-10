package errorregion

import "github.com/e-fish/api/pkg/common/helper/werror"

var (
	ErrValidateInput = werror.Error{
		Code:    "ValidateError",
		Message: "field can't be empty",
	}

	ErrFailedCreateAddress = werror.Error{
		Code:    "FailedCreateAddress",
		Message: "failed create address",
	}

	ErrFailedUpdateAddress = werror.Error{
		Code:    "FailedUpdateAddress",
		Message: "failed update address",
	}

	ErrFailedDeletedAddress = werror.Error{
		Code:    "FaileDeleteAddress",
		Message: "failed delete address",
	}

	ErrFoundCountry = werror.Error{
		Code:    "FailedFoundCountry",
		Message: "country not found",
	}

	ErrFoundProvince = werror.Error{
		Code:    "FailedFoundProvince",
		Message: "province not found",
	}

	ErrFoundCity = werror.Error{
		Code:    "FailedFoundCity",
		Message: "city not found",
	}

	ErrFoundDistrict = werror.Error{
		Code:    "FailedFoundDistrict",
		Message: "country not found",
	}

	ErrReadCountry = werror.Error{
		Code:    "FailedReadRecordQueryCountry",
		Message: "internal server error",
	}

	ErrReadProvince = werror.Error{
		Code:    "FailedReadRecordQueryProvince",
		Message: "internal server error",
	}

	ErrReadCity = werror.Error{
		Code:    "FailedReadRecordQueryCity",
		Message: "internal server error",
	}

	ErrReadDistrict = werror.Error{
		Code:    "FailedReadRecordQueryDistrict",
		Message: "internal server error",
	}

	ErrAddressNotFound = werror.Error{
		Code:    "FailedFoundAddress",
		Message: "address not found",
	}

	ErrReadAddress = werror.Error{
		Code:    "FailedReadRecordQueryAddress",
		Message: "internal server error",
	}

	ErrCommit = werror.Error{
		Code:    "FailedCommitTransaction",
		Message: "can't commit region transaction",
	}
	ErrRollback = werror.Error{
		Code:    "FailedRollbackTransaction",
		Message: "can't rollback region transaction",
	}

	ErrFoundNearestTenant = werror.Error{
		Code:    "FailedFoundNearestTenant",
		Message: "cannot find the nearest tenant you are looking for",
	}

	ErrReadNearestTenant = werror.Error{
		Code:    "FailedReadNearestTenant",
		Message: "failed to read nearest tenant data",
	}

	ErrFailedSetTenantInAddress = werror.Error{
		Code:    "FailedSetTenantInAddress",
		Message: "failed to update tenant id in address",
	}
)
