package budidayahttp

import (
	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	productconfig "github.com/e-fish/api/product_http/product_config"
	productservice "github.com/e-fish/api/product_http/product_service"
	"github.com/e-fish/api/product_http/producthandler"
	"github.com/gin-gonic/gin"
)

type route struct {
	conf productconfig.ProductConfig
	gin  *gin.Engine
}

func newRoute(ro route) {
	ginEngine := ro.gin

	service := productservice.NewService(ro.conf)
	handler := producthandler.Handler{
		Conf:    ro.conf,
		Service: service,
	}

	ginEngine.POST("/create-product", ctxutil.Authorization(), handler.CreateProduct)
	ginEngine.POST("/delete-product", ctxutil.Authorization(), handler.DeleteProduct)
}
