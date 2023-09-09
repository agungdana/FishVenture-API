package wshandler

import (
	wsservice "github.com/e-fish/api/ws-http/ws_service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Serv *wsservice.Service
}

func (ws *Handler) WsConect(c *gin.Context) {
	ws.Serv.WsConnection(c.Writer, c.Request)
}
