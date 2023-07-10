package region

import (
	"github.com/e-fish/api/pkg/domain/region/model"
	"github.com/google/uuid"
)

// map[countryID]Country
var mapCountry = map[uuid.UUID]model.CountryOutput{}

// map[provinceID]Province
var mapProvince = map[uuid.UUID]model.ProvinceOutput{}

// map[cityID]City
var mapCity = map[uuid.UUID]model.CityOutput{}

// map[districtID]District
var mapDistrict = map[uuid.UUID]model.DistrictOutput{}
