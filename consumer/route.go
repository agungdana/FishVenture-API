package consumer

import (
	"context"

	consumerconfig "github.com/e-fish/api/consumer/consumer_config"
	consumerhandler "github.com/e-fish/api/consumer/consumer_handler"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/common/infra/event/redis"
)

func newRoute(conf consumerconfig.ConsumerConfig) {
	var (
		ctx = context.Background()
	)

	event, err := redis.NewRedisClient(conf.AppConfig.RedisConfig)
	if err != nil {
		logger.Fatal("failed create new redis client")
	}

	handler := consumerhandler.NewHandler(conf, event.NewEventPubsub())

	handler.ChatHandler(ctx)
}
