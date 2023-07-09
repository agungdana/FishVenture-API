package budidayahandler

import (
	budidayaconfig "github.com/e-fish/api/budidaya_http/budidaya_config"
	budidayaservice "github.com/e-fish/api/budidaya_http/budidaya_service"
	"github.com/e-fish/api/pkg/common/helper/restsvr"
	"github.com/e-fish/api/pkg/common/helper/werror"
	"github.com/e-fish/api/pkg/domain/budidaya/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	Conf    budidayaconfig.BudidayaConfig
	Service budidayaservice.Service
}

func (h *Handler) CreateBudidaya(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req model.CreateBudidayaInput
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}

	result, err := h.Service.CreateBudidaya(ctx, req)
	res.Add(result, err)
}

func (h *Handler) CreateFishSpecies(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req model.CreateFishSpeciesInput
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}

	result, err := h.Service.CreateFishSpecies(ctx, req)
	res.Add(result, err)
}

func (h *Handler) CreateMultiplePricelist(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req model.CreateMultiplePriceListInput
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}

	result, err := h.Service.CreateMultiplePricelist(ctx, req)
	res.Add(result, err)
}

func (h *Handler) GetBudidayaAdminAndCustomer(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	pondID, err := uuid.Parse(c.Query("id"))
	if err != nil {
		res.Add(nil, werror.Error{
			Code:    "NeedPondID",
			Message: "need param pond id",
			Details: map[string]any{
				"error": err,
			},
		})
		return
	}

	result, err := h.Service.GetBudidayaByUserLoginAdminOrCustomer(ctx, model.GetBudidayaInput{
		PondID: pondID,
	})
	res.Add(result, err)
}

func (h *Handler) GetBudidayaForSaller(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	result, err := h.Service.GetBudidayaByUserSaller(ctx)
	res.Add(result, err)
}
