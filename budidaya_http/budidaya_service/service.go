package budidayaservice

import (
	budidayaconfig "github.com/e-fish/api/budidaya_http/budidaya_config"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/domain/budidaya"
	"github.com/e-fish/api/pkg/domain/verification"
)

func NewService(conf budidayaconfig.BudidayaConfig) Service {

	verificationRepo, err := verification.NewRepo(conf.DbConfig)
	if err != nil {
		logger.Fatal("###failed create budidaya service err: %v", err)
	}

	buidayaRepo, err := budidaya.NewRepo(conf.DbConfig, verificationRepo)
	if err != nil {
		logger.Fatal("###failed create budidaya service err: %v", err)
	}

	service := Service{
		conf: conf,
		repo: buidayaRepo,
	}

	return service
}

type Service struct {
	conf budidayaconfig.BudidayaConfig
	repo budidaya.Repo
}
