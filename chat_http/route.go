package bannerhttp

import (
	chatconfig "github.com/e-fish/api/chat_http/chat_config"
	chathandler "github.com/e-fish/api/chat_http/chat_handler"
	"github.com/e-fish/api/chat_http/chatservice"
	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type route struct {
	conf chatconfig.ChatConfig
	gin  *gin.Engine
}

func newRoute(ro route) {
	ginEngine := ro.gin

	service := chatservice.NewService(ro.conf)
	handler := chathandler.Handler{
		Conf:    ro.conf,
		Service: service,
	}

	ginEngine.POST("/create-chat", ctxutil.Authorization(), handler.CreateChat)
	ginEngine.POST("/create-chat-item", ctxutil.Authorization(), handler.CreateChatItem)
	ginEngine.GET("/list-chat", ctxutil.Authorization(), handler.ReadListChat)
	ginEngine.GET("/list-chat-item", ctxutil.Authorization(), handler.ReadListChatItems)

	ginEngine.POST("/upload-chat", handler.SaveImageChat)
	ginEngine.Use(static.Serve("/assets/image/chat", static.LocalFile(ro.conf.ChatImageConfig.Path, false)))
}
