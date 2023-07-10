package region

import (
	"context"
	"errors"

	"github.com/e-fish/api/pkg/domain/region/model"
	errorregion "github.com/e-fish/api/pkg/domain/region/model/error-region"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func newQuery(db *gorm.DB) Query {
	return &query{db: db}
}

// lock implements Query.
func (q *query) lock() Query {
	db := q.db.Clauses(clause.Locking{Strength: "UPDATE"})
	return &query{db: db}
}

type query struct {
	db *gorm.DB
}

// ReadCityByID implements Query.
func (q *query) ReadCityByID(ctx context.Context, input uuid.UUID) (*model.CityOutput, error) {
	data := model.CityOutput{}

	if val, ok := mapCity[input]; ok {
		return &val, nil
	}

	err := q.db.Where("deleted_at IS NULL and id = ?", input).Take(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorregion.ErrFoundCity.AttacthDetail(map[string]any{"id": input})
		}
		return nil, errorregion.ErrReadCity.AttacthDetail(map[string]any{"error": err})
	}

	mapCity[input] = data
	return &data, nil
}

// ReadCountryByID implements Query.
func (q *query) ReadCountryByID(ctx context.Context, input uuid.UUID) (*model.CountryOutput, error) {
	data := model.CountryOutput{}

	if val, ok := mapCountry[input]; ok {
		return &val, nil
	}

	err := q.db.Where("deleted_at IS NULL and id = ?", input).Take(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorregion.ErrFoundCountry.AttacthDetail(map[string]any{"id": input})
		}
		return nil, errorregion.ErrReadCountry.AttacthDetail(map[string]any{"error": err})
	}

	mapCountry[input] = data
	return &data, nil
}

// ReadDistrictByID implements Query.
func (q *query) ReadDistrictByID(ctx context.Context, input uuid.UUID, withPreload bool) (*model.DistrictOutput, error) {
	var (
		data = model.DistrictOutput{}
		db   = q.db
	)

	if val, ok := q.getDataDistrict(ctx, input, withPreload); ok {
		return val, nil
	}

	if withPreload {
		db = db.Preload("City.Province.Country")
	}

	//TODO: enhance
	err := db.Where("deleted_at IS NULL and id = ?", input).Take(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorregion.ErrFoundDistrict.AttacthDetail(map[string]any{"id": input})
		}
		return nil, errorregion.ErrReadDistrict.AttacthDetail(map[string]any{"error": err})
	}

	mapDistrict[input] = data

	return &data, nil
}

func (q *query) getDataDistrict(ctx context.Context, input uuid.UUID, withPreload bool) (*model.DistrictOutput, bool) {
	val, ok := mapDistrict[input]
	if !ok {
		return nil, false
	}

	if withPreload {
		city, err := q.ReadCityByID(ctx, val.CityID)
		if err != nil {
			return nil, false
		}
		province, err := q.ReadProvinceByID(ctx, city.ProvinceID)
		if err != nil {
			return nil, false
		}
		country, err := q.ReadCountryByID(ctx, province.CountryID)
		if err != nil {
			return nil, false
		}
		val.City = city
		val.City.Province = province
		val.City.Province.Country = country
	}

	return &val, true

}

// ReadProvinceByID implements Query.
func (q *query) ReadProvinceByID(ctx context.Context, input uuid.UUID) (*model.ProvinceOutput, error) {
	data := model.ProvinceOutput{}

	if val, ok := mapProvince[input]; ok {
		return &val, nil
	}

	err := q.db.Where("deleted_at IS NULL and id = ?", input).Take(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorregion.ErrFoundProvince.AttacthDetail(map[string]any{"id": input})
		}
		return nil, errorregion.ErrReadProvince.AttacthDetail(map[string]any{"error": err})
	}

	mapProvince[input] = data
	return &data, nil
}

// ReadAllCityByProvinceID implements Query.
func (q *query) ReadAllCityByProvinceID(ctx context.Context, input uuid.UUID) ([]model.CityOutput, error) {
	data := []model.CityOutput{}

	err := q.db.Where("deleted_at IS NULL and province_id = ? and is_coverage = ?", input, true).Find(&data).Error
	if err != nil {
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errorregion.ErrFoundCity.AttacthDetail(map[string]any{"province_id": input})
			}
			return nil, errorregion.ErrReadCity.AttacthDetail(map[string]any{"error": err})
		}
	}

	if len(data) < 1 {
		return data, errorregion.ErrFoundCity.AttacthDetail(map[string]any{"province_id": input})
	}

	return data, nil
}

// ReadAllCountry implements Query.
func (q *query) ReadAllCountry(ctx context.Context) ([]model.CountryOutput, error) {
	data := []model.CountryOutput{}

	err := q.db.Where("deleted_at IS NULL and is_coverage = ?", true).Find(&data).Error
	if err != nil {
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errorregion.ErrFoundCountry
			}
			return nil, errorregion.ErrReadCountry.AttacthDetail(map[string]any{"error": err})
		}
	}

	if len(data) < 1 {
		return data, errorregion.ErrFoundCountry
	}

	return data, nil
}

// ReadAllDistrictByCityID implements Query.
func (q *query) ReadAllDistrictByCityID(ctx context.Context, input uuid.UUID) ([]model.DistrictOutput, error) {
	data := []model.DistrictOutput{}

	err := q.db.Where("deleted_at IS NULL and city_id = ? and is_coverage = ?", input, true).Find(&data).Error
	if err != nil {
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errorregion.ErrFoundDistrict.AttacthDetail(map[string]any{"city_id": input})
			}
			return nil, errorregion.ErrReadDistrict.AttacthDetail(map[string]any{"error": err})
		}
	}

	if len(data) < 1 {
		return data, errorregion.ErrFoundDistrict.AttacthDetail(map[string]any{"city_id": input})
	}

	return data, nil
}

// ReadAllProvinceByCountryID implements Query.
func (q *query) ReadAllProvinceByCountryID(ctx context.Context, input uuid.UUID) ([]model.ProvinceOutput, error) {
	data := []model.ProvinceOutput{}

	err := q.db.Where("deleted_at IS NULL and country_id = ? and is_coverage = ?", input, true).Find(&data).Error
	if err != nil {
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errorregion.ErrFoundProvince.AttacthDetail(map[string]any{"country_id": input})
			}
			return nil, errorregion.ErrReadProvince.AttacthDetail(map[string]any{"error": err})
		}
	}

	if len(data) < 1 {
		return data, errorregion.ErrFoundProvince.AttacthDetail(map[string]any{"country_id": input})
	}

	return data, nil
}
