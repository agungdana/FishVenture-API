package wshttp

import (
	"github.com/e-fish/api/pkg/common/helper/restsvr"
	wsconfig "github.com/e-fish/api/ws-http/ws_config"
)

func NewWsHttp() {
	var (
		ginEngine = restsvr.GetGinRoute()
		conf      = wsconfig.GetConfig()
	)

	newRoute(route{
		conf: *conf,
		gin:  ginEngine,
	})
}
