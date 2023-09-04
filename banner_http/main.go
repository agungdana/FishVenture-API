package bannerhttp

import (
	bannerconfig "github.com/e-fish/api/banner_http/banner_config"
	"github.com/e-fish/api/pkg/common/helper/restsvr"
)

func NewRegionHttp() {
	var (
		ginEngine = restsvr.GetGinRoute()
		conf      = bannerconfig.GetConfig()
	)

	newRoute(route{
		conf: *conf,
		gin:  ginEngine,
	})
}
