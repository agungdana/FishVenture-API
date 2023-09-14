package bannerhttp

import (
	chatconfig "github.com/e-fish/api/chat_http/chat_config"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/common/helper/restsvr"
	"github.com/e-fish/api/pkg/common/infra/event/redis"
)

func NewRegionHttp() {
	var (
		ginEngine = restsvr.GetGinRoute()
		conf      = chatconfig.GetConfig()
	)

	event, err := redis.NewRedisClient(conf.AppConfig.RedisConfig)
	if err != nil {
		logger.Fatal("failed create new redis client")
	}

	newRoute(route{
		conf:  *conf,
		gin:   ginEngine,
		event: event,
	})
}
