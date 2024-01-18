package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SecretKey     string
}

var AppConfig Config

func InitConfig() {
	godotenv.Load()

	AppConfig = Config{
		SecretKey:     os.Getenv("SECRET_KEY"),
	}
}

