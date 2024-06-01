package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database_URI string
}

var Configs = InitConfig()

func InitConfig() Config {

	err := godotenv.Load()

	if err!=nil {
		log.Fatal("Error while loading .env files: ",err)
	}

	return Config{
		Database_URI: os.Getenv("DATABASE_URI"),
	}

}