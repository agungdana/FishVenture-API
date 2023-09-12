package schedulerconfig

import (
	"os"
	"strconv"
	"sync"

	budidayaconfig "github.com/e-fish/api/budidaya_http/budidaya_config"
	"github.com/e-fish/api/pkg/common/helper/config"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/common/helper/werror"
	transactionconfig "github.com/e-fish/api/transaction_http/transaction_config"
	"github.com/joho/godotenv"
)

var (
	conf *SchedulerConfig
	once sync.Once
)

type SchedulerConfig struct {
	BudidayaConfig        budidayaconfig.BudidayaConfig
	TransactionConfig     transactionconfig.TransactionConfig
	TimerUpdate           int
	CreateProductAfterRun bool
}

func getConfig() *SchedulerConfig {
	if conf == nil {
		once.Do(func() {
			err := godotenv.Load()
			if err != nil {
				logger.Fatal("error load env err: %v", config.ErrLoadEnv.AttacthDetail(map[string]any{"location": "region-config", "err": err}))
				return
			}

			//timer for what time the data will be updated
			timerUpdate, _ := strconv.Atoi(os.Getenv("TIMER_UPDATE_DATA"))

			isRun, _ := strconv.ParseBool(os.Getenv("UPDATE_DATA_AFTER_RUN"))

			conf = &SchedulerConfig{
				BudidayaConfig:        *budidayaconfig.GetConfig(),
				TransactionConfig:     *transactionconfig.GetConfig(),
				TimerUpdate:           timerUpdate,
				CreateProductAfterRun: isRun,
			}
		})
	}
	return conf
}

func GetConfig() *SchedulerConfig {
	conf := getConfig()

	errs := werror.NewError("incomplete tenant configuration")

	dbConf := conf.BudidayaConfig.DbConfig

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
