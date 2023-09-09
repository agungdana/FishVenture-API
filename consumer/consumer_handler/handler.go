package consumerhandler

import (
	"context"
	"fmt"

	consumerconfig "github.com/e-fish/api/consumer/consumer_config"
	chatservice "github.com/e-fish/api/consumer/consumer_service/chat_service"
	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/common/infra/event"
)

func NewHandler(conf consumerconfig.ConsumerConfig, event event.EventPubsub) *Handler {
	return &Handler{
		conf:  conf,
		event: event,
	}
}

type Handler struct {
	conf  consumerconfig.ConsumerConfig
	event event.EventPubsub
}

func (h *Handler) ChatHandler(ctx context.Context) {
	var (
		service = chatservice.NewService(h.event)
		topic   = event.GetTopic(ctx, h.conf.OrderConsumerConfig.Topic, h.conf.AppConfig.Debug)
	)

	_, err := h.event.Subscribe(ctx, topic, func(m *event.ClientMessages) {
		var (
			ctx = ctxutil.ToContextUsingMap(m.CtxMap)
		)

		defer func() {
			if err := recover(); err != nil {
				logger.ErrorWithContext(ctx, fmt.Sprintf("panic error %s recover: %v", topic, err))
			}
		}()

		switch m.Action {
		case PUBLISH_TO_ALL:
			req := chatservice.ReadChatRequest{}
			req.OrderRequestFromMap(m.Data.(map[string]any))
			service.PublishChatToAll(ctx, req)
		}

	})

	if err != nil {
		logger.Fatal("failed subs topic [%s]", topic)
	}

}
