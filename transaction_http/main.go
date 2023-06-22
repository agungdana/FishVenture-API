package transactionhttp

import (
	"github.com/e-fish/api/pkg/common/helper/restsvr"
	productconfig "github.com/e-fish/api/product_http/product_config"
)

func NewTransactionHttp() {
	var (
		ginEngine = restsvr.GetGinRoute()
		conf      = productconfig.GetConfig()
	)

	newRoute(route{
		conf: *conf,
		gin:  ginEngine,
	})
}
