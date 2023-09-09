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

	ginEngine.POST("/create-budidaya", ctxutil.Authorization(), handler.CreateBudidaya)
	ginEngine.POST("/create-fish-species", ctxutil.Authorization(), handler.CreateFishSpecies)
	ginEngine.POST("/create-multiple-pricelist", ctxutil.Authorization(), handler.CreateMultiplePricelist)

	ginEngine.GET("/list-fish-species", handler.GetAllFishSpecies)
	ginEngine.GET("/list-budidaya-seller", ctxutil.Authorization(), handler.GetBudidayaForSeller)
	ginEngine.GET("/list-budidaya", ctxutil.Authorization(), handler.GetBudidayaAdminAndCustomer)
	ginEngine.GET("/nearest-budidaya", handler.ReadBudidayaNeaerest)
}
