package errorproduct

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

	ErrValidateCreateInput = werror.Error{
		Code:    "FailedValidateInputProduct",
		Message: "field can't be emtpy",
	}

	ErrFoundProduct = werror.Error{
		Code:    "FailedFoundProduct",
		Message: "The product you are looking for cannot be found",
	}

	ErrReadProduct = werror.Error{
		Code:    "FailedFoundProduct",
		Message: "Failed to read data product due to internal server error",
	}

	ErrCreateProductExist = werror.Error{
		Code:    "FailedCreateProductExist",
		Message: "Cannot create product data. Product already exists",
	}

	ErrCreateProduct = werror.Error{
		Code:    "FailedCreateProduct",
		Message: "Cannot create product data. internal server error",
	}

	ErrDeleteProduct = werror.Error{
		Code:    "FailedDeleteProduct",
		Message: "Cannot delete product data. internal server error",
	}
)
