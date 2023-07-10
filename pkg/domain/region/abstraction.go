package region

import (
	"context"

	"github.com/e-fish/api/pkg/domain/region/model"
	"github.com/google/uuid"
)

type Repo interface {
	NewCommand(ctx context.Context) Command
	NewQuery() Query
}

type Command interface {
	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
}

type Query interface {
	ReadAllCountry(ctx context.Context) ([]model.CountryOutput, error)
	ReadAllProvinceByCountryID(ctx context.Context, input uuid.UUID) ([]model.ProvinceOutput, error)
	ReadAllCityByProvinceID(ctx context.Context, input uuid.UUID) ([]model.CityOutput, error)
	ReadAllDistrictByCityID(ctx context.Context, input uuid.UUID) ([]model.DistrictOutput, error)

	ReadCountryByID(ctx context.Context, input uuid.UUID) (*model.CountryOutput, error)
	ReadProvinceByID(ctx context.Context, input uuid.UUID) (*model.ProvinceOutput, error)
	ReadCityByID(ctx context.Context, input uuid.UUID) (*model.CityOutput, error)
	ReadDistrictByID(ctx context.Context, input uuid.UUID, withPreload bool) (*model.DistrictOutput, error)
	getDataDistrict(ctx context.Context, input uuid.UUID, withPreload bool) (*model.DistrictOutput, bool)

	lock() Query
}
