package budidayahttp

import (
	budidayaconfig "github.com/e-fish/api/budidaya_http/budidaya_config"
	budidayahandler "github.com/e-fish/api/budidaya_http/budidaya_handler"
	budidayaservice "github.com/e-fish/api/budidaya_http/budidaya_service"
	"github.com/e-fish/api/pkg/common/helper/ctxutil"
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
	ginEngine.GET("/pond", ctxutil.Authorization(), handler.GetPondByUserID)

	ginEngine.GET("/all-pond-submission", ctxutil.Authorization(), handler.GetAllPondSubmission)
	ginEngine.POST("/update-pond-status", ctxutil.Authorization(), handler.UpdatePondStatus)
}
