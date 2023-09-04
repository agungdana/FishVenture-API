package bannerHandler

import (
	bannerconfig "github.com/e-fish/api/banner_http/banner_config"
	bannerService "github.com/e-fish/api/banner_http/banner_service"
	"github.com/e-fish/api/pkg/common/helper/restsvr"
	"github.com/e-fish/api/pkg/common/helper/werror"
	"github.com/e-fish/api/pkg/domain/banner/model"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Conf    bannerconfig.BannerConfig
	Service bannerService.Service
}

func (h *Handler) GetListBanner(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	result, err := h.Service.ListBanner(ctx)
	res.Add(result, err)
}

func (h *Handler) CreateBanner(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req = model.BannerInputCreate{}
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	if err := c.ShouldBindJSON(&req); err != nil {
		res.Add(nil, werror.Error{
			Code:    "FailedBindJson",
			Message: "failed ShouldBindJSON",
			Details: map[string]any{
				"err": err,
			},
		})
		return
	}

	result, err := h.Service.CreateBanner(ctx, req)
	res.Add(result, err)
}

func (h *Handler) UpdateBanner(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req = model.BannerInputUpdate{}
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	if err := c.ShouldBindJSON(&req); err != nil {
		res.Add(nil, werror.Error{
			Code:    "FailedBindJson",
			Message: "failed ShouldBindJSON",
			Details: map[string]any{
				"err": err,
			},
		})
		return
	}

	result, err := h.Service.UpdateBanner(ctx, req)
	res.Add(result, err)
}
