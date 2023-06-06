package restsvr

import (
	"github.com/e-fish/api/pkg/common/helper/werror"
	"github.com/google/uuid"
)

const (
	StatusSuccess = "success"
	StatusWarning = "warning"
	StatusFailed  = "failed"
)

type ErrorResponse struct {
	Code    string         `json:"code,omitempty"`
	Message string         `json:"message,omitempty"`
	Details map[string]any `json:"details,omitempty"`
}

type HttpResponse struct {
	Status  string          `json:"status,omitempty"`
	Id      *uuid.UUID      `json:"id,omitempty"`
	Ids     []*uuid.UUID    `json:"ids,omitempty"`
	Data    any             `json:"data,omitempty"`
	Message string          `json:"message,omitempty"`
	Error   []ErrorResponse `json:"error,omitempty"`
}

func errorFromWerror(err werror.Error) ErrorResponse {
	return ErrorResponse{
		Code:    err.Code,
		Message: err.Message,
		Details: err.Details,
	}
}

func (h *HttpResponse) generateError(err error) {
	if err == nil {
		return
	}
	errs := []ErrorResponse{}
	switch e := err.(type) {
	case werror.Error:
		h.Message = e.Message
		errs = append(errs, errorFromWerror(e))
	case *werror.Error:
		h.Message = e.Message
		errs = append(errs, errorFromWerror(*e))
	case werror.Errors:
		h.Message = e.Message
		for _, v := range e.Errors {
			errs = append(errs, errorFromWerror(v))
		}
	case *werror.Errors:
		h.Message = e.Message
		for _, v := range e.Errors {
			errs = append(errs, errorFromWerror(v))
		}
	case error:
		h.Message = e.Error()
		errs = append(errs, ErrorResponse{
			Code:    "InternalServerError",
			Message: e.Error(),
			Details: nil,
		})
	}

	h.Error = errs
}

func (h *HttpResponse) Add(data any, err error) {

	h.Status = StatusSuccess
	h.Message = StatusSuccess

	switch d := data.(type) {
	case uuid.UUID:
		h.Id = &d
	case *uuid.UUID:
		h.Id = d
	case []*uuid.UUID:
		h.Ids = d
	default:
		h.Data = d
	}

	if err != nil {
		h.Status = StatusFailed
		h.generateError(err)
	}
}
