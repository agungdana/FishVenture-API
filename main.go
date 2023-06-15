package main

import (
	authhttp "github.com/e-fish/api/auth_http"
	mainconfig "github.com/e-fish/api/main_config"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/common/helper/ptime"
	"github.com/e-fish/api/pkg/common/helper/restsvr"
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

	// migrations.Migrations()

	restsvr.Run()

}
