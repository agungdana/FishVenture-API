package regionconfig

import (
	"os"
	"sync"

	"github.com/e-fish/api/pkg/common/helper/config"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/common/helper/werror"
	"github.com/joho/godotenv"
)

var (
	conf *RegionConfig
	once sync.Once
)

type RegionConfig struct {
	RegionDBConfig config.DbConfig
}

func getConfig() *RegionConfig {
	if conf == nil {
		once.Do(func() {
			err := godotenv.Load()
			if err != nil {
				logger.Fatal("error load env err: %v", config.ErrLoadEnv.AttacthDetail(map[string]any{"location": "region-config", "err": err}))
				return
			}

			driverRegion := os.Getenv("DB_DRIVER_REGION")
			hostRegion := os.Getenv("DB_HOST_REGION")
			databaseRegion := os.Getenv("DB_NAME_REGION")
			usernameRegion := os.Getenv("DB_USERNAME_REGION")
			passwordRegion := os.Getenv("DB_PASSWORD_REGION")
			portRegion := os.Getenv("DB_PORT_REGION")

			conf = &RegionConfig{
				RegionDBConfig: config.DbConfig{
					Driver:   driverRegion,
					Host:     hostRegion,
					User:     usernameRegion,
					Password: passwordRegion,
					Database: databaseRegion,
					Port:     portRegion,
				},
			}
		})
	}
	return conf
}

func GetConfig() *RegionConfig {
	conf := getConfig()

	errs := werror.NewError("incomplete configuration")

	dbConf := conf.RegionDBConfig

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
