package mainconfig

import (
	"os"
	"sync"

	"github.com/e-fish/api/pkg/common/helper/config"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/joho/godotenv"
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
			err := godotenv.Load()
			if err != nil {
				logger.Fatal("error load env err: %v", config.ErrLoadEnv.AttacthDetail(map[string]any{"location": "auth-config", "err": err}))
				return
			}

			name := os.Getenv("APP_NAME")
			host := os.Getenv("APP_HOST")
			port := os.Getenv("APP_PORT")
			debug := os.Getenv("APP_DEBUG")
			firebaseConf := os.Getenv("FIREBASE_CONF")

			conf = &Config{
				AppConfig: config.AppConfig{
					Name:  name,
					Host:  host,
					Port:  port,
					Debug: debug,
					FirebaseConfig: config.FirebaseConfig{
						FireBase: firebaseConf,
					},
				},
			}
		})
	}
	return conf
}

func GetConfig() *Config {
	return getConfig()
}
