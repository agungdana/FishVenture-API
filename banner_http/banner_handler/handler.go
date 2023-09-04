package bannerHandler

import (
	bannerconfig "github.com/e-fish/api/banner_http/banner_config"
	bannerService "github.com/e-fish/api/banner_http/banner_service"
	"github.com/e-fish/api/pkg/common/helper/restsvr"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Conf    bannerconfig.BannerConfig
	Service bannerService.Service
}

func (h *Handler) GetListBanner(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	result, err := h.Service.ListBanner(ctx)
	res.Add(result, err)
}
