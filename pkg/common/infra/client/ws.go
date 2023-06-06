package client

import (
	"github.com/e-fish/api/pkg/common/helper/config"
	"github.com/gorilla/websocket"
)

type WsConnection struct {
	conf config.AppConfig
	ws   *websocket.Conn
}
