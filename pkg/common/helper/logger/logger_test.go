package logger_test

import (
	"testing"

	"github.com/e-fish/api/pkg/common/helper/config"
	"github.com/e-fish/api/pkg/common/helper/logger"
)

func TestLooger(t *testing.T) {
	conf := config.AppConfig{
		Name:    "Paled-ID",
		Address: "localhost",
		Port:    "8081",
		Debug:   "true",
	}

	logger.SetupLogger(conf.Debug)

	logger.Info("example")
}

func BenchmarkLooger(b *testing.B) {
	conf := config.AppConfig{
		Name:    "Paled-ID",
		Address: "localhost",
		Port:    "8081",
		Debug:   "true",
	}

	logger.SetupLogger(conf.Debug)

	logger.Info("example")
}
