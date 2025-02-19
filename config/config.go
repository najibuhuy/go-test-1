package config

import (
	"fmt"
	"go-test-1/infrastructure/database"
	"go-test-1/infrastructure/shared/constant"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvironmentConfig struct {
	Env      string
	App      App
	Database database.DatabaseConfig
}

type App struct {
	Name    string
	Version string
	Port    int
}

func LoadENVConfig() (config EnvironmentConfig, err error) {
	err = godotenv.Load()
	if err != nil {
		err = fmt.Errorf(constant.ErrLoadENV, err)
		return
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		err = fmt.Errorf(constant.ErrConvertStringToInt, err)
		return
	}

	// rmqPort := 0
	// if os.Getenv("RABBITMQ_PORT") != "" {
	// 	rmqPort, err = strconv.Atoi(os.Getenv("RABBITMQ_PORT"))
	// 	if err != nil {
	// 		err = fmt.Errorf(constant.ErrConvertStringToInt, err)
	// 		return
	// 	}
	// }

	config = EnvironmentConfig{
		Env: os.Getenv("ENV"),
		App: App{
			Name:    os.Getenv("APP_NAME"),
			Version: os.Getenv("APP_VERSION"),
			Port:    port,
		},
	}

	return
}
