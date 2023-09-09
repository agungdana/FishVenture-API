package errorchat

import "github.com/e-fish/api/pkg/common/helper/werror"

var (
	ErrCommit = werror.Error{
		Code:    "FailedCommitTransaction",
		Message: "can't commit chat transaction",
	}
	ErrRollback = werror.Error{
		Code:    "FailedRollbackTransaction",
		Message: "can't rollback chat transaction",
	}

	ErrFailedCreateChat = werror.Error{
		Code:    "FailedCreateChat",
		Message: "failed create chat",
	}

	ErrFailedCreateChatItem = werror.Error{
		Code:    "FailedCreateChatItem",
		Message: "failed create chat item",
	}

	ErrFoundChat = werror.Error{
		Code:    "FailedFoundChat",
		Message: "failed found chat",
	}

	ErrReadChatData = werror.Error{
		Code:    "FailedReadData",
		Message: "failed read chat data",
	}
)
