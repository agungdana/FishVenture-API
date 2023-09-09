package schedulerhandler

import (
	schedulerconfig "github.com/e-fish/api/scheduler/scheduler_config"
	schedulerservice "github.com/e-fish/api/scheduler/scheduler_service"
)

type Handler struct {
	Conf    schedulerconfig.SchedulerConfig
	Service schedulerservice.Service
}
