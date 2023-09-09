package chatconfig

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
	conf *ChatConfig
	once sync.Once
)

type ChatConfig struct {
	AppConfig       config.AppConfig
	ChatImageConfig config.ImageConfig
	ChatDBConfig    config.DbConfig
}

func getConfig() *ChatConfig {
	if conf == nil {
		once.Do(func() {
			err := godotenv.Load()
			if err != nil {
				logger.Fatal("error load env err: %v", config.ErrLoadEnv.AttacthDetail(map[string]any{"location": "chat-config", "err": err}))
				return
			}

			driver := os.Getenv("DB_DRIVER")
			host := os.Getenv("DB_HOST")
			database := os.Getenv("DB_NAME")
			username := os.Getenv("DB_USERNAME")
			password := os.Getenv("DB_PASSWORD")
			port := os.Getenv("DB_PORT")

			bannerImagePath := os.Getenv("PATH_IMAGE_CHAT")
			bannerImageUrl := os.Getenv("URL_IMAGE_CHAT")

			mainConf := mainconfig.GetConfig()

			conf = &ChatConfig{
				AppConfig: mainConf.AppConfig,
				ChatDBConfig: config.DbConfig{
					Driver:   driver,
					Host:     host,
					User:     username,
					Password: password,
					Database: database,
					Port:     port,
				},
				ChatImageConfig: config.ImageConfig{
					Url:  bannerImageUrl,
					Path: bannerImagePath,
				},
			}
		})
	}
	return conf
}

func GetConfig() *ChatConfig {
	conf := getConfig()

	errs := werror.NewError("incomplete configuration")

	dbConf := conf.ChatDBConfig

	if dbConf.Driver == "" {
		errs.Add(config.ErrEmptyConfig.AttacthDetail(map[string]any{"field empty Driver": "empty"}))
	}
	if dbConf.Host == "" {
		errs.Add(config.ErrEmptyConfig.AttacthDetail(map[string]any{"field empty Host": "empty"}))
	}
	if dbConf.User == "" {
		errs.Add(config.ErrEmptyConfig.AttacthDetail(map[string]any{"field empty User": "empty"}))
	}
	if dbConf.Password == "" {
		errs.Add(config.ErrEmptyConfig.AttacthDetail(map[string]any{"field empty Password": "empty"}))
	}
	if dbConf.Database == "" {
		errs.Add(config.ErrEmptyConfig.AttacthDetail(map[string]any{"field empty Database": "empty"}))
	}
	if dbConf.Port == "" {
		errs.Add(config.ErrEmptyConfig.AttacthDetail(map[string]any{"field empty Port": "empty"}))
	}

	if err := errs.Return(); err != nil {
		logger.Fatal("auth-config err: %v", err)
		return nil
	}

	return conf
}
