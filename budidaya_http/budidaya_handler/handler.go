package budidayahandler

import (
	budidayaconfig "github.com/e-fish/api/budidaya_http/budidaya_config"
	budidayaservice "github.com/e-fish/api/budidaya_http/budidaya_service"
)

type Handler struct {
	Conf    budidayaconfig.BudidayaConfig
	Service budidayaservice.Service
}
