package producthandler

import (
	"github.com/e-fish/api/pkg/common/helper/restsvr"
	"github.com/e-fish/api/pkg/domain/product/model"
	productconfig "github.com/e-fish/api/product_http/product_config"
	productservice "github.com/e-fish/api/product_http/product_service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	Conf    productconfig.ProductConfig
	Service productservice.Service
}

func (h *Handler) CreateProduct(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req model.CreateProductInput
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}

	result, err := h.Service.CreateProduct(ctx, req)
	res.Add(result, err)
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	uid, err := uuid.Parse(c.Query("id"))
	if err != nil {
		res.Add(nil, err)
		return
	}

	result, err := h.Service.DeleteProduct(ctx, uid)
	res.Add(result, err)
}
