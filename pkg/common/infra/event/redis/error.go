package redis

import "github.com/e-fish/api/pkg/common/helper/werror"

var (
	ErrFailedCreateRedisConnection = werror.Error{
		Code:    "FailedCreateRedisConnection",
		Message: "failed to establish a new connection to the Redis server",
	}

	ErrFailedSetData = werror.Error{
		Code:    "FailedSetDataRedis",
		Message: "failed to set data to the Redis server.",
	}

	ErrFailedGetData = werror.Error{
		Code:    "FailedGetDataRedis",
		Message: "failed to get data to the Redis server.",
	}

	ErrFailedUnmarshalData = werror.Error{
		Code:    "FailedUnmarsha;DataRedis",
		Message: "failed to unmarshal data from the Redis server.",
	}
)
