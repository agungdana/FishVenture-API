package consumer

import consumerconfig "github.com/e-fish/api/consumer/consumer_config"

func NewConsumer() {
	conf := consumerconfig.GetConfig()
	newRoute(*conf)
}
