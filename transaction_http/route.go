package transactionhttp

import (
	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	transactionconfig "github.com/e-fish/api/transaction_http/transaction_config"
	transactionhandler "github.com/e-fish/api/transaction_http/transaction_handler"
	transactionservice "github.com/e-fish/api/transaction_http/transaction_service"

	"github.com/gin-gonic/gin"
)

type route struct {
	conf transactionconfig.TransactionConfig
	gin  *gin.Engine
}

func newRoute(ro route) {
	ginEngine := ro.gin

	service := transactionservice.NewService(ro.conf)
	handler := transactionhandler.Handler{
		Conf:    ro.conf,
		Service: service,
	}

	ginEngine.POST("/create-order", ctxutil.Authorization(), handler.CreateOrder)
	ginEngine.POST("/update-order-cancel", ctxutil.Authorization(), handler.UpdateOrderCancel)
	ginEngine.POST("/update-order-success", ctxutil.Authorization(), handler.UpdateSuccessOrder)
	ginEngine.GET("/order", ctxutil.Authorization(), handler.GetOrder)
	ginEngine.GET("/order-cancel", ctxutil.Authorization(), handler.GetOrderCancel)
	ginEngine.GET("/order-success", ctxutil.Authorization(), handler.GetOrderSuccess)
}
