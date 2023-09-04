package bannerhttp

import (
	bannerconfig "github.com/e-fish/api/banner_http/banner_config"
	bannerHandler "github.com/e-fish/api/banner_http/banner_handler"
	bannerService "github.com/e-fish/api/banner_http/banner_service"
	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/gin-contrib/static"
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
	ginEngine.POST("/update-banner", ctxutil.Authorization(), handler.UpdateBanner)
	ginEngine.POST("/create-banner", ctxutil.Authorization(), handler.CreateBanner)
	ginEngine.POST("/create-banner", ctxutil.Authorization(), handler.CreateBanner)

	ginEngine.POST("/upload-file", handler.SaveImageBanner)
	ginEngine.Use(static.Serve("/assets/image/banner", static.LocalFile(ro.conf.BannerImageConfig.Path, false)))
}
