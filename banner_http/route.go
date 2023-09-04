package bannerhttp

import (
	bannerconfig "github.com/e-fish/api/banner_http/banner_config"
	bannerHandler "github.com/e-fish/api/banner_http/banner_handler"
	bannerService "github.com/e-fish/api/banner_http/banner_service"
	"github.com/gin-gonic/gin"
)

type route struct {
	conf bannerconfig.BannerConfig
	gin  *gin.Engine
}

func newRoute(ro route) {
	ginEngine := ro.gin

	service := bannerService.NewService(ro.conf)
	handler := bannerHandler.Handler{
		Conf:    ro.conf,
		Service: service,
	}

	ginEngine.GET("/banner", handler.GetListBanner)
}
