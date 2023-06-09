package authhttp

import (
	authconfig "github.com/e-fish/api/auth_http/auth_config"
	"github.com/e-fish/api/pkg/common/helper/restsvr"
)

func NewAuthHttp() {
	var (
		ginEngine = restsvr.GetGinRoute()
		conf      = authconfig.GetConfig()
	)

	newRoute(route{
		conf: *conf,
		gin:  ginEngine,
	})
}
