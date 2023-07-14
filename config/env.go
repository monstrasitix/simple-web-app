package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadDev() {
	godotenv.Load(".env.development")
}

func LoadProd() {
	godotenv.Load(".env.production")
}

func GetPort() string {
	return os.Getenv("PORT")
}
