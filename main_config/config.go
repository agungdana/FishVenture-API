package mainconfig

import (
	"sync"

	"github.com/e-fish/api/pkg/common/helper/config"
)

var (
	conf *Config
	once sync.Once
)

type Config struct {
	config.AppConfig
}

// single tone
// to avoid reading env multiple times
func getConfig() *Config {
	if conf == nil {
		once.Do(func() {
			conf = &Config{
				AppConfig: config.AppConfig{},
			}
		})
	}
	return conf
}

func GetConfig() *Config {
	return getConfig()
}
