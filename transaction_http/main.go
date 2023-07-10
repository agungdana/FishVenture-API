package transactionhttp

import (
	"github.com/e-fish/api/pkg/common/helper/restsvr"
	transactionconfig "github.com/e-fish/api/transaction_http/transaction_config"
)

func NewTransactionHttp() {
	var (
		ginEngine = restsvr.GetGinRoute()
		conf      = transactionconfig.GetConfig()
	)

	newRoute(route{
		conf: *conf,
		gin:  ginEngine,
	})
}
