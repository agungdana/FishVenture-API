package chathandler

import (
	chatconfig "github.com/e-fish/api/chat_http/chat_config"
	"github.com/e-fish/api/chat_http/chatservice"
	"github.com/e-fish/api/pkg/common/helper/restsvr"
	"github.com/e-fish/api/pkg/domain/chat/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	Conf    chatconfig.ChatConfig
	Service chatservice.Service
}

func (h *Handler) CreateChat(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req = model.CreateChatInput{}
		res = new(restsvr.HttpResponse)
	)
	defer restsvr.ResponsJson(c, res)

	err := c.ShouldBindJSON(&req)

	if err != nil {
		res.Add(nil, err)
		return
	}
	result, err := h.Service.CreateChat(ctx, req)
	res.Add(result, err)
}

func (h *Handler) CreateChatItem(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req = model.CreateChatItemInput{}
		res = new(restsvr.HttpResponse)
	)
	defer restsvr.ResponsJson(c, res)

	err := c.ShouldBindJSON(&req)

	if err != nil {
		res.Add(nil, err)
		return
	}
	result, err := h.Service.CreateChatItem(ctx, req)
	res.Add(result, err)
}

func (h *Handler) ReadListChat(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		res = new(restsvr.HttpResponse)
	)
	defer restsvr.ResponsJson(c, res)

	result, err := h.Service.ReadListChat(ctx)
	res.Add(result, err)
}

func (h *Handler) ReadListChatItems(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		res = new(restsvr.HttpResponse)
	)
	defer restsvr.ResponsJson(c, res)

	uid, err := uuid.Parse(c.Query("id"))
	if err != nil {
		res.Add(nil, err)
		return
	}
	result, err := h.Service.ReadListChatItemByID(ctx, uid)
	res.Add(result, err)
}

func (h *Handler) SaveImageChat(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		res = new(restsvr.HttpResponse)
	)
	defer restsvr.ResponsJson(c, res)
	file, err := c.FormFile("file")
	if err != nil {
		res.Add(nil, err)
		return
	}
	result, err := h.Service.SaveImageChat(ctx, file)
	res.Add(result, err)
}
