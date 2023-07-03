package budidayahandler

import (
	budidayaconfig "github.com/e-fish/api/budidaya_http/budidaya_config"
	budidayaservice "github.com/e-fish/api/budidaya_http/budidaya_service"
	"github.com/e-fish/api/pkg/common/helper/restsvr"
	"github.com/e-fish/api/pkg/domain/budidaya/model"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Conf    budidayaconfig.BudidayaConfig
	Service budidayaservice.Service
}

func (h *Handler) CreatePond(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req model.CreatePondInput
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}

	result, err := h.Service.CreatePond(ctx, req)
	res.Add(result, err)
}

func (h *Handler) UpdatePond(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req model.UpdatePondInput
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}

	result, err := h.Service.UpdatePond(ctx, req)
	res.Add(result, err)
}

func (h *Handler) UpdatePondStatus(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req model.UpdatePondStatus
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}

	result, err := h.Service.UpdatePondStatus(ctx, req)
	res.Add(result, err)
}

func (h *Handler) GetPondByUserAdmin(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	result, err := h.Service.GetPondByUserID(ctx)
	res.Add(result, err)
}

func (h *Handler) GetAllPondSubmission(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	result, err := h.Service.GetAllPondSubmission(ctx)
	res.Add(result, err)
}

func (h *Handler) GetAllPondForUser(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	result, err := h.Service.GetListPondForUser(ctx)
	res.Add(result, err)
}
