package authhandler

import (
	authconfig "github.com/e-fish/api/auth_http/auth_config"
	authservice "github.com/e-fish/api/auth_http/auth_service"
	"github.com/e-fish/api/pkg/common/helper/restsvr"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Conf    authconfig.AuthConfig
	Service authservice.Service
}

func (h *Handler) CreateUser(c *gin.Context) {
	var (
		// ctx = c.Request.Context()
		// req any
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	// err := c.ShouldBindJSON(&req)

	// result, err := h.Service.CreateUser(ctx, req)
	// res.Add(result, err)
}
