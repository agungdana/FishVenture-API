package bannerhttp

import (
	chatconfig "github.com/e-fish/api/chat_http/chat_config"
	"github.com/e-fish/api/pkg/common/helper/restsvr"
)

func NewRegionHttp() {
	var (
		ginEngine = restsvr.GetGinRoute()
		conf      = chatconfig.GetConfig()
	)

	newRoute(route{
		conf: *conf,
		gin:  ginEngine,
	})
}
