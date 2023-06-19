package budidayahttp

import (
	budidayaconfig "github.com/e-fish/api/budidaya_http/budidaya_config"
	"github.com/e-fish/api/pkg/common/helper/restsvr"
)

func NewAuthHttp() {
	var (
		ginEngine = restsvr.GetGinRoute()
		conf      = budidayaconfig.GetConfig()
	)

	newRoute(route{
		conf: *conf,
		gin:  ginEngine,
	})
}
