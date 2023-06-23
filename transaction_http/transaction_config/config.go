package transactionconfig

import (
	"os"
	"sync"

	"github.com/e-fish/api/pkg/common/helper/config"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/common/helper/werror"
	"github.com/joho/godotenv"
)

var (
	conf *TransactionConfig
	once sync.Once
)

type TransactionConfig struct {
	DbConfig config.DbConfig
}

// single tone
// to avoid reading env multiple times
func getConfig() *TransactionConfig {
	if conf == nil {
		once.Do(func() {
			err := godotenv.Load()
			if err != nil {
				logger.Fatal("error load env err: %v", config.ErrLoadEnv.AttacthDetail(map[string]any{"location": "budidaya-config", "err": err}))
				return
			}

			driver := os.Getenv("DB_DRIVER")
			host := os.Getenv("DB_HOST")
			database := os.Getenv("DB_NAME")
			username := os.Getenv("DB_USERNAME")
			password := os.Getenv("DB_PASSWORD")
			port := os.Getenv("DB_PORT")

			conf = &TransactionConfig{
				DbConfig: config.DbConfig{
					Driver:   driver,
					Host:     host,
					User:     username,
					Password: password,
					Database: database,
					Port:     port,
				},
			}
		})
	}
	return conf
}

func GetConfig() *TransactionConfig {
	conf := getConfig()

	errs := werror.NewError("incomplete configuration budidaya")

	dbConf := conf.DbConfig

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
