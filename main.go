package main

import (
	authhttp "github.com/e-fish/api/auth_http"
	bannerhttp "github.com/e-fish/api/banner_http"
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
	//1
	ptime.SetDefaultTimeToUTC()

	//get main config
	//1
	conf := mainconfig.GetConfig()
	//1
	logger.SetupLogger(conf.Debug)

	//create new route
	//1
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
	//register banner http in main
	bannerhttp.NewRegionHttp()

	//1
	restsvr.Run()

}
