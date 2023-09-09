package wshttp

import (
	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	wsconfig "github.com/e-fish/api/ws-http/ws_config"
	wshandler "github.com/e-fish/api/ws-http/ws_handler"
	wsservice "github.com/e-fish/api/ws-http/ws_service"
	"github.com/gin-gonic/gin"
)

type route struct {
	conf wsconfig.WsConfig
	gin  *gin.Engine
}

func newRoute(ro route) {
	ginEngine := ro.gin

	service := wsservice.NewService(ro.conf)
	handler := wshandler.Handler{
		Serv: service,
	}

	ginEngine.GET("/ws", ctxutil.Authorization(), handler.WsConect)
}
