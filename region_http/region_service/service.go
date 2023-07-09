package regionservice

import (
	"context"

	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/domain/region"
	"github.com/e-fish/api/pkg/domain/region/model"
	regionconfig "github.com/e-fish/api/region_http/region_config"
	"github.com/google/uuid"
)

type Service struct {
	conf regionconfig.RegionConfig
	repo region.Repo
}

func NewService(conf regionconfig.RegionConfig) Service {
	var (
		service = Service{
			conf: conf,
		}
	)

	regionRepo, err := region.NewRepo(conf.RegionDBConfig)
	if err != nil {
		logger.Fatal("failed to create a new repo, can't create region service err causes failed create region repo: %v", err)
	}

	service.repo = regionRepo

	return service
}

func (s *Service) ListCountry(ctx context.Context) ([]model.CountryOutput, error) {
	query := s.repo.NewQuery()
	result, err := query.ReadAllCountry(ctx)
	if err != nil {
		logger.ErrorWithContext(ctx, "failed get list country: %v", err)
	}
	return result, err
}

func (s *Service) ListProvinceByCountryID(ctx context.Context, input uuid.UUID) ([]model.ProvinceOutput, error) {
	query := s.repo.NewQuery()
	result, err := query.ReadAllProvinceByCountryID(ctx, input)
	if err != nil {
		logger.ErrorWithContext(ctx, "failed get list province by country id: %v", err)
	}
	return result, err
}

func (s *Service) ListCityByProvinceID(ctx context.Context, input uuid.UUID) ([]model.CityOutput, error) {
	query := s.repo.NewQuery()
	result, err := query.ReadAllCityByProvinceID(ctx, input)
	if err != nil {
		logger.ErrorWithContext(ctx, "failed get list city by province id: %v", err)
	}
	return result, err
}

func (s *Service) ListDistrictByCityID(ctx context.Context, input uuid.UUID) ([]model.DistrictOutput, error) {
	query := s.repo.NewQuery()
	result, err := query.ReadAllDistrictByCityID(ctx, input)
	if err != nil {
		logger.ErrorWithContext(ctx, "failed get list district by city id: %v", err)
	}
	return result, err
}
