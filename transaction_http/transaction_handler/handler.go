package transactionhandler

import (
	"strconv"

	"github.com/e-fish/api/pkg/common/helper/restsvr"
	"github.com/e-fish/api/pkg/common/infra/orm"
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

func (h *Handler) GetOrder(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		res = new(restsvr.HttpResponse)
	)
	defer restsvr.ResponsJson(c, res)

	findBy := c.Query("findBy")
	keyword := c.Query("keyword")
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	sort := c.Query("sort")
	direction := c.Query("direction")
	objectTable := c.Query("objectTable")

	result, err := h.Service.ReadOrder(ctx, model.ReadInput{
		Paginantion: orm.Paginantion{
			FindBy:      findBy,
			Keyword:     keyword,
			Limit:       limit,
			Page:        page,
			Sort:        sort,
			Direction:   direction,
			ObjectTable: objectTable,
		},
	})
	res.Add(result, err)
}

func (h *Handler) GetOrderSuccess(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		res = new(restsvr.HttpResponse)
	)
	defer restsvr.ResponsJson(c, res)

	findBy := c.Query("findBy")
	keyword := c.Query("keyword")
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	sort := c.Query("sort")
	direction := c.Query("direction")
	objectTable := c.Query("objectTable")

	result, err := h.Service.ReadOrderSuccess(ctx, model.ReadInput{
		Paginantion: orm.Paginantion{
			FindBy:      findBy,
			Keyword:     keyword,
			Limit:       limit,
			Page:        page,
			Sort:        sort,
			Direction:   direction,
			ObjectTable: objectTable,
		},
	})
	res.Add(result, err)
}

func (h *Handler) GetOrderCancel(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		res = new(restsvr.HttpResponse)
	)
	defer restsvr.ResponsJson(c, res)

	findBy := c.Query("findBy")
	keyword := c.Query("keyword")
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	sort := c.Query("sort")
	direction := c.Query("direction")
	objectTable := c.Query("objectTable")

	result, err := h.Service.ReadOrderCancel(ctx, model.ReadInput{
		Paginantion: orm.Paginantion{
			FindBy:      findBy,
			Keyword:     keyword,
			Limit:       limit,
			Page:        page,
			Sort:        sort,
			Direction:   direction,
			ObjectTable: objectTable,
		},
	})
	res.Add(result, err)
}

func (h *Handler) UpdateOrderCancel(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req = transactionservice.TransactionReqUpdate{}
		res = new(restsvr.HttpResponse)
	)
	defer restsvr.ResponsJson(c, res)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}

	result, err := h.Service.UpdateOrderCancel(ctx, req.ID)
	res.Add(result, err)
}

func (h *Handler) UpdateSuccessOrder(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req = transactionservice.TransactionReqUpdate{}
		res = new(restsvr.HttpResponse)
	)
	defer restsvr.ResponsJson(c, res)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}

	result, err := h.Service.UpdateOrderSuccess(ctx, req.ID)
	res.Add(result, err)
}
