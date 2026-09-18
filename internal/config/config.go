package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	DBHost string
	DBPort string
	DBUser string
	DBPassword string
	DBName string
}

func LoadConfig() Config {
	godotenv.Load()

	return Config{
		Port: os.Getenv("PORT"),
		DBHost: os.Getenv("DBHOST"),
		DBPort: os.Getenv("DBPORT"),
		DBUser: os.Getenv("DBUSER"),
		DBPassword: os.Getenv("DBPASSWORD"),
		DBName: os.Getenv("DBNAME"),
	}
}