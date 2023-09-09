package consumerservice

import (
	consumerconfig "github.com/e-fish/api/consumer/consumer_config"
	"github.com/e-fish/api/pkg/common/infra/event"
)

type Service struct {
	conf  consumerconfig.ConsumerConfig
	event event.EventPubsub
}

func NewService(conf consumerconfig.ConsumerConfig, event event.EventPubsub) Service {
	return Service{
		conf:  conf,
		event: event,
	}
}
