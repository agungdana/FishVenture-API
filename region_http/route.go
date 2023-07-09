package regionhttp

import (
	regionconfig "github.com/e-fish/api/region_http/region_config"
	regionhandler "github.com/e-fish/api/region_http/region_handler"
	regionservice "github.com/e-fish/api/region_http/region_service"
	"github.com/gin-gonic/gin"
)

type route struct {
	conf regionconfig.RegionConfig
	gin  *gin.Engine
}

func newRoute(ro route) {
	ginEngine := ro.gin

	service := regionservice.NewService(ro.conf)
	handler := regionhandler.Handler{
		Conf:    ro.conf,
		Service: service,
	}

	ginEngine.GET("/list-country", handler.ListCountry)
	ginEngine.GET("/list-province-country", handler.ListProvinceByCountryID)
	ginEngine.GET("/list-city-province", handler.ListCityByProvinceID)
	ginEngine.GET("/list-district-city", handler.ListDistrictByCityID)

}
