package authhandler

import (
	authconfig "github.com/e-fish/api/auth_http/auth_config"
	authservice "github.com/e-fish/api/auth_http/auth_service"
	"github.com/e-fish/api/pkg/common/helper/restsvr"
	"github.com/e-fish/api/pkg/domain/auth/model"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Conf    authconfig.AuthConfig
	Service authservice.Service
}

func (h *Handler) CreateUser(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req model.CreateUserInput
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}

	result, err := h.Service.CreateUser(ctx, req)
	res.Add(result, err)
}

func (h *Handler) Login(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req model.UserLoginInput
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}

	result, err := h.Service.Login(ctx, req)
	res.Add(result, err)
}

func (h *Handler) LoginByGoogle(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req model.UserLoginByGooleInput
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}

	result, err := h.Service.LoginByGoogle(ctx, req)
	res.Add(result, err)
}

func (h *Handler) Profile(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	result, err := h.Service.Profile(ctx)
	res.Add(result, err)
}

func (h *Handler) SaveImage(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		res = new(restsvr.HttpResponse)
	)
	defer restsvr.ResponsJson(c, res)
	file, err := c.FormFile("image")
	if err != nil {
		res.Add(nil, err)
		return
	}
	result, err := h.Service.SaveImages(ctx, file)
	res.Add(result, err)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req model.UpdateUserInput
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}

	result, err := h.Service.UpdateUser(ctx, req)
	res.Add(result, err)
}
