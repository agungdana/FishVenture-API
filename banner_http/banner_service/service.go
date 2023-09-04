package bannerService

import (
	"context"

	bannerconfig "github.com/e-fish/api/banner_http/banner_config"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/domain/banner"
	"github.com/e-fish/api/pkg/domain/banner/model"
)

type Service struct {
	conf bannerconfig.BannerConfig
	repo banner.Repo
}

func NewService(conf bannerconfig.BannerConfig) Service {
	var (
		service = Service{
			conf: conf,
		}
	)

	bannerRepo, err := banner.NewRepo(conf.BannerDBConfig)
	if err != nil {
		logger.Fatal("failed to create a new repo, can't create region service err causes failed create region repo: %v", err)
	}

	service.repo = bannerRepo

	return service
}

func (s *Service) ListBanner(ctx context.Context) ([]model.BannerOutput, error) {
	query := s.repo.NewQuery()
	result, err := query.ReadAllBanner(ctx)
	if err != nil {
		logger.ErrorWithContext(ctx, "failed get list district by city id: %v", err)
	}
	return result, err
}
