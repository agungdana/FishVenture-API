package pondhandler

import (
	"github.com/e-fish/api/pkg/common/helper/restsvr"
	"github.com/e-fish/api/pkg/common/helper/werror"
	"github.com/e-fish/api/pkg/domain/pond/model"
	pondconfig "github.com/e-fish/api/pond_http/pond_config"
	pondservice "github.com/e-fish/api/pond_http/pond_service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	Conf    pondconfig.PondConfig
	Service pondservice.Service
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

func (h *Handler) ResubmissionPond(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req model.Resubmission
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}

	result, err := h.Service.ResubmissionPond(ctx, req)
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

func (h *Handler) GetAllPond(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	result, err := h.Service.GetListPond(ctx)
	res.Add(result, err)
}

func (h *Handler) GetListPool(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		res = new(restsvr.HttpResponse)
	)

	defer restsvr.ResponsJson(c, res)

	uid, err := uuid.Parse(c.Query("id"))
	if err != nil {
		res.Add(nil, werror.Error{
			Code:    "NeedUUID",
			Message: "need uuid",
			Details: map[string]any{"error": err},
		})
		return
	}

	result, err := h.Service.GetListPool(ctx, uid)
	res.Add(result, err)
}

func (h *Handler) SaveImagePond(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		res = new(restsvr.HttpResponse)
	)
	defer restsvr.ResponsJson(c, res)
	file, err := c.FormFile("image")
	if err != nil {
		res.Add(nil, err)
		return
	}
	result, err := h.Service.SaveImagesPond(ctx, file)
	res.Add(result, err)
}

func (h *Handler) SaveFilePond(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		res = new(restsvr.HttpResponse)
	)
	defer restsvr.ResponsJson(c, res)
	file, err := c.FormFile("file")
	if err != nil {
		res.Add(nil, err)
		return
	}
	result, err := h.Service.SaveFilePond(ctx, file)
	res.Add(result, err)
}

func (h *Handler) SaveImagePool(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		res = new(restsvr.HttpResponse)
	)
	defer restsvr.ResponsJson(c, res)
	file, err := c.FormFile("image")
	if err != nil {
		res.Add(nil, err)
		return
	}
	result, err := h.Service.SaveImagesPool(ctx, file)
	res.Add(result, err)
}
