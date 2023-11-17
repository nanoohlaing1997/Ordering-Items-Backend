package env

import (
	"log"
	"os"
	"path"
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/nanoohlaing1997/online-ordering-items/env/config"
)

const settings = "ORDERING_SETTINGS"

type Envs struct {
	*config.AppConfig
	*config.DbConfig
}

var (
	envOnce     sync.Once
	envInstance *Envs
)

func isProduction(env string, appConfig *config.AppConfig) {
	envs := map[string]bool{
		"local":   true,
		"test":    true,
		"testing": true,
	}

	appConfig.IsProduction = !envs[env]
}

func GetEnviroment() *Envs {
	envOnce.Do(func() {
		settingPath, ok := os.LookupEnv(settings)

		if !ok {
			settingPath = ""
		}
		file := path.Join(settingPath, ".env")
		if err := godotenv.Load(file); err != nil {
			log.Println(err)
		}

		var appConfig config.AppConfig
		var dbConfig config.DbConfig

		if err := envconfig.Process("", &appConfig); err != nil {
			log.Println(err)
		}

		if err := envconfig.Process("", &dbConfig); err != nil {
			log.Println(err)
		}

		isProduction(appConfig.AppEnv, &appConfig)

		envInstance = &Envs{
			AppConfig: &appConfig,
			DbConfig:  &dbConfig,
		}
	})

	return envInstance
}
