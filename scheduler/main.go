package scheduler

import (
	schedulerconfig "github.com/e-fish/api/scheduler/scheduler_config"
	schedulerservice "github.com/e-fish/api/scheduler/scheduler_service"
)

func Scheduler() {
	var (
		conf = schedulerconfig.GetConfig()
	)

	schedulerservice.NewService(*conf)
}
