package consumerconfig

import (
	"os"
	"sync"

	mainconfig "github.com/e-fish/api/main_config"
	"github.com/e-fish/api/pkg/common/helper/config"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/common/helper/werror"
	"github.com/joho/godotenv"
)

var (
	conf *ConsumerConfig
	once sync.Once
)

type ConsumerConfig struct {
	AppConfig             config.AppConfig
	OrderConsumerConfig   config.ConsumerConfig
	ProductConsumerConfig config.ConsumerConfig
}

func getConfig() *ConsumerConfig {
	if conf == nil {
		once.Do(func() {
			err := godotenv.Load()
			if err != nil {
				logger.Fatal("error load env err: %v", config.ErrLoadEnv.AttacthDetail(map[string]any{"location": "region-config", "err": err}))
				return
			}

			appConfg := mainconfig.GetConfig()
			conf = &ConsumerConfig{
				AppConfig: appConfg.AppConfig,
				OrderConsumerConfig: config.ConsumerConfig{
					Topic: os.Getenv("CHAT_CONSUMER_TOPIC"),
				},
			}

		})
	}
	return conf
}

func GetConfig() *ConsumerConfig {
	conf := getConfig()

	errs := werror.NewError("incomplete consumer configuration")

	if conf.OrderConsumerConfig.Topic == "" {
		errs.Add(werror.Error{
			Code:    "FailedFoundConfig",
			Message: "order consumer config is empty",
		})
	}

	if err := errs.Return(); err != nil {
		logger.Fatal("consumer-config err: %v", err)
		return nil
	}

	return conf
}
