package wsconfig

import (
	"sync"

	mainconfig "github.com/e-fish/api/main_config"
	"github.com/e-fish/api/pkg/common/helper/config"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/common/helper/werror"
	"github.com/joho/godotenv"
)

var (
	conf *WsConfig
	once sync.Once
)

type WsConfig struct {
	WsAppConfig config.AppConfig
}

func getConfig() *WsConfig {
	if conf == nil {
		once.Do(func() {
			err := godotenv.Load()
			if err != nil {
				logger.Fatal("error load env err: %v", config.ErrLoadEnv.AttacthDetail(map[string]any{"location": "region-config", "err": err}))
				return
			}

			appConf := mainconfig.GetConfig()

			conf = &WsConfig{
				WsAppConfig: config.AppConfig{
					Name:        appConf.Name,
					Host:        appConf.Host,
					Port:        appConf.Port,
					Debug:       appConf.Debug,
					RedisConfig: appConf.RedisConfig,
				},
			}
		})
	}
	return conf
}

func GetConfig() *WsConfig {
	conf := getConfig()

	redisConfig := conf.WsAppConfig.RedisConfig

	errs := werror.NewError("failed get redis config from env")

	if redisConfig.Host == "" {
		errs.Add(werror.Error{
			Code:    "FailedFoundEnv",
			Message: "failed found env",
			Details: map[string]any{"Host": "empty"},
		})
	}
	if redisConfig.Port == "" {
		errs.Add(werror.Error{
			Code:    "FailedFoundEnv",
			Message: "failed found env",
			Details: map[string]any{"Port": "empty"},
		})
	}

	if err := errs.Return(); err != nil {
		logger.Fatal("failed found redis config err: %v", errs.Return())
	}

	return conf
}
