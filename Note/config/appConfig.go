package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
	Dsn string
	AppSecret string
}

func SetupEnv() (cfg AppConfig, err error) {

	if os.Getenv("APP_ENV") == "dev"{
		godotenv.Load()
	}
	httpPort := os.Getenv("HTTP_PORT")
	if len(httpPort) < 1{
		return AppConfig{}, errors.New("port env variables not found")
	}

	//DB connection string
	Dsn:= os.Getenv("DSN");
	if len(Dsn) < 1{
		return AppConfig{}, errors.New("dsn env variables not found")
	}

	appSecret := os.Getenv("APP_SECRET")
	if len(appSecret) < 1{
		return AppConfig{}, errors.New("app secret not found")
	}
	return AppConfig{ServerPort: httpPort, Dsn: Dsn, AppSecret: appSecret,}, nil
}