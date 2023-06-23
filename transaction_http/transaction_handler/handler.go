package producthandler

import (
	"github.com/e-fish/api/pkg/common/helper/restsvr"
	"github.com/e-fish/api/pkg/domain/transaction/model"
	transactionconfig "github.com/e-fish/api/transaction_http/transaction_config"
	transactionservice "github.com/e-fish/api/transaction_http/transaction_service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Conf    transactionconfig.TransactionConfig
	Service transactionservice.Service
}

func (h *Handler) CreateOrder(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req = model.CreateOrderInput{}
		res = new(restsvr.HttpResponse)
	)
	defer restsvr.ResponsJson(c, res)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}

	result, err := h.Service.CreateOrderInput(ctx, req)
	res.Add(result, err)
}
