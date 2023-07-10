package main

import (
	authhttp "github.com/e-fish/api/auth_http"
	budidayahttp "github.com/e-fish/api/budidaya_http"
	mainconfig "github.com/e-fish/api/main_config"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/common/helper/ptime"
	"github.com/e-fish/api/pkg/common/helper/restsvr"
	pondhttp "github.com/e-fish/api/pond_http"
	regionhttp "github.com/e-fish/api/region_http"
	transactionhttp "github.com/e-fish/api/transaction_http"
)

func main() {
	ptime.SetDefaultTimeToUTC()

	//get main config
	conf := mainconfig.GetConfig()
	logger.SetupLogger(conf.Debug)

	//create new route
	restsvr.NewRoute(conf.AppConfig)

	//register auth http in main
	authhttp.NewAuthHttp()
	//register budidaya http in main
	budidayahttp.NewBudidayaHttp()
	//register pond http in main
	pondhttp.NewPondHttp()
	//register transaction http in main
	transactionhttp.NewTransactionHttp()
	//register region http in main
	regionhttp.NewRegionHttp()

	restsvr.Run()

}
