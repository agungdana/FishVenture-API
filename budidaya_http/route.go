package budidayahttp

import (
	budidayaconfig "github.com/e-fish/api/budidaya_http/budidaya_config"
	budidayahandler "github.com/e-fish/api/budidaya_http/budidaya_handler"
	budidayaservice "github.com/e-fish/api/budidaya_http/budidaya_service"
	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type route struct {
	conf budidayaconfig.BudidayaConfig
	gin  *gin.Engine
}

func newRoute(ro route) {
	ginEngine := ro.gin

	service := budidayaservice.NewService(ro.conf)
	handler := budidayahandler.Handler{
		Conf:    ro.conf,
		Service: service,
	}

	ginEngine.POST("/create-pond", ctxutil.Authorization(), handler.CreatePond)
	ginEngine.POST("/update-pond", ctxutil.Authorization(), handler.UpdatePond)
	ginEngine.GET("/pond", ctxutil.Authorization(), handler.GetPondByUserAdmin)
	ginEngine.GET("/list-pond", handler.GetAllPondForUser)

	ginEngine.GET("/all-pond-submission", ctxutil.Authorization(), handler.GetAllPondSubmission)
	ginEngine.POST("/update-pond-status", ctxutil.Authorization(), handler.UpdatePondStatus)

	ginEngine.GET("/list-budidaya", ctxutil.Authorization())
	ginEngine.GET("/list-budidaya-active")

	ginEngine.POST("/upload-photo-product", handler.SaveImage)
	ginEngine.Use(static.Serve("/assets/image/product", static.LocalFile(ro.conf.ImageConfig.Path, false)))

}
