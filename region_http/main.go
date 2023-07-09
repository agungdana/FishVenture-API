package regionhttp

import (
	"github.com/e-fish/api/pkg/common/helper/restsvr"
	regionconfig "github.com/e-fish/api/region_http/region_config"
)

func NewRegionHttp() {
	var (
		ginEngine = restsvr.GetGinRoute()
		conf      = regionconfig.GetConfig()
	)

	newRoute(route{
		conf: *conf,
		gin:  ginEngine,
	})
}
