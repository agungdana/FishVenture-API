package chatservice

import (
	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/google/uuid"
)

type ReadChatRequest struct {
	ChatID     uuid.UUID `json:"chatID"`
	ChatItemID uuid.UUID `json:"chatItemID"`
	Valid      `json:"valid"`
}

func (r *ReadChatRequest) ChatRequestFromMap(data map[string]any) *ReadChatRequest {
	res := ReadChatRequest{}

	if val, ok := data["chatID"]; ok {
		r.ChatID = ctxutil.ToUUID(val)
	}

	if val, ok := data["chatItemID"]; ok {
		r.ChatItemID = ctxutil.ToUUID(val)
	}

	r.Valid.FromMap(data)

	return &res
}

type Valid struct {
	Counter int   `json:"counter"`
	Err     error `json:"err"`
}

func (v *Valid) FromMap(data map[string]any) {
	if val, ok := data["counter"]; ok {
		v.Counter = val.(int)
	}

	if val, ok := data["counter"]; ok {
		v.Err = val.(error)
	}
}
