package wsservice

import (
	"net/http"

	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/common/helper/restsvr"
	"github.com/e-fish/api/pkg/common/infra/event"
	"github.com/e-fish/api/pkg/common/infra/event/redis"
	wsconfig "github.com/e-fish/api/ws-http/ws_config"
)

type Service struct {
	conf      wsconfig.WsConfig
	client    event.Event
	wsManager *restsvr.WsManager
}

func NewService(wsConf wsconfig.WsConfig) *Service {
	client, err := redis.NewRedisClient(wsConf.WsAppConfig.RedisConfig)
	if err != nil {
		logger.Fatal("###Failed Create new redis client err: %v", err)
		return nil
	}

	return &Service{
		conf:      wsConf,
		client:    client,
		wsManager: restsvr.NewWsManager(wsConf.WsAppConfig, client),
	}
}

func (s *Service) WsConnection(writer http.ResponseWriter, req *http.Request) {
	var (
		ctx = req.Context()
	)

	wsConn, err := restsvr.Upgrade(writer, req)
	if err != nil {
		logger.ErrorWithContext(ctx, "failed to connect ws %v", err)
		return
	}

	s.wsManager.WsRun(ctx, wsConn)
}
