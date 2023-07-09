package pondhttp

import (
	"github.com/e-fish/api/pkg/common/helper/restsvr"
	pondconfig "github.com/e-fish/api/pond_http/pond_config"
)

func NewPondHttp() {
	var (
		ginEngine = restsvr.GetGinRoute()
		conf      = pondconfig.GetConfig()
	)

	newRoute(route{
		conf: *conf,
		gin:  ginEngine,
	})
}
