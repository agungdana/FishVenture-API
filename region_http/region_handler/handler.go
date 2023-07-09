package regionhandler

import (
	"github.com/e-fish/api/pkg/common/helper/restsvr"
	"github.com/e-fish/api/pkg/common/helper/werror"
	regionconfig "github.com/e-fish/api/region_http/region_config"
	regionservice "github.com/e-fish/api/region_http/region_service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	Conf    regionconfig.RegionConfig
	Service regionservice.Service
}

func (h *Handler) ListCountry(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	result, err := h.Service.ListCountry(ctx)
	res.Add(result, err)
}

func (h *Handler) ListProvinceByCountryID(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	uid, err := uuid.Parse(c.Query("id"))

	if err != nil {
		res.Add(nil, werror.Error{
			Code:    "FailedMarshalID",
			Message: "need id",
			Details: map[string]any{
				"error": err,
			},
		})
		return
	}

	result, err := h.Service.ListProvinceByCountryID(ctx, uid)
	res.Add(result, err)
}

func (h *Handler) ListCityByProvinceID(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	uid, err := uuid.Parse(c.Query("id"))

	if err != nil {
		res.Add(nil, werror.Error{
			Code:    "FailedMarshalID",
			Message: "need id",
			Details: map[string]any{
				"error": err,
			},
		})
		return
	}

	result, err := h.Service.ListCityByProvinceID(ctx, uid)
	res.Add(result, err)
}

func (h *Handler) ListDistrictByCityID(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	uid, err := uuid.Parse(c.Query("id"))

	if err != nil {
		res.Add(nil, werror.Error{
			Code:    "FailedMarshalID",
			Message: "need id",
			Details: map[string]any{
				"error": err,
			},
		})
		return
	}

	result, err := h.Service.ListDistrictByCityID(ctx, uid)
	res.Add(result, err)
}
