package budidayahttp

import (
	"fmt"

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

	fmt.Printf("handler: %v\n", handler)

	ginEngine.POST("/create-budidaya", ctxutil.Authorization())

}
