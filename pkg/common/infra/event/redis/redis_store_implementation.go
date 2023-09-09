package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/e-fish/api/pkg/common/helper/config"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/common/infra/event"
	"github.com/redis/go-redis/v9"
)

func newRedisEventStore(conf config.RedisConfig, client *redis.Client) event.EventStore {
	return &redisEventStore{
		conf:   conf,
		client: client,
	}
}

type redisEventStore struct {
	conf   config.RedisConfig
	client *redis.Client
}

// Get implements Redis.
func (r *redisEventStore) Get(ctx context.Context, key string, data interface{}) error {
	logger.InfoWithContext(ctx, "####Get data to redis wiht key: %s", key)

	result, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return ErrFailedGetData.AttacthDetail(map[string]any{"error-redis": err})
	}

	err = json.Unmarshal([]byte(result), data)
	if err != nil {
		return ErrFailedUnmarshalData.AttacthDetail(map[string]any{"error-marshal": err})
	}

	return nil
}

// Set implements Redis.
func (r *redisEventStore) Set(ctx context.Context, param event.Payload) (err error) {
	logger.InfoWithContext(ctx, "####Set data to redis wiht key: %s", param.Key)
	if param.ExpiredTime == 0 {
		param.ExpiredTime = time.Minute * time.Duration(r.conf.ExpiredTime)
	}

	data, ok := param.Data.([]byte)
	if !ok {
		data, err = json.Marshal(&param.Data)
		if err != nil {
			err = ErrFailedSetData.AttacthDetail(map[string]any{"error": err})
			return
		}
	}
	fmt.Printf("data: %v\n", data)

	err = r.client.Set(ctx, param.Key, data, param.ExpiredTime).Err()
	if err != nil {
		err = ErrFailedSetData.AttacthDetail(map[string]any{"error": err})
	}

	return
}
