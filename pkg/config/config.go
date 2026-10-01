package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Mode        string
	HOST        string
	PORT        int
	TimeZone    string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	DB_USERNAME string
	DB_PASSWORD string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	port, err := strconv.Atoi(getEnv("PORT", "9000"))
	if err != nil {
		fmt.Println("Failed convert port")
	}

	return &Config{
		Mode:        getEnv("MODE", "development"),
		HOST:        getEnv("HOST", "localhost"),
		PORT:        port,
		TimeZone:    getEnv("TIME_ZONE", "Asia/Jakarta"),
		DB_HOST:     getEnv("DB_HOST", "localhost"),
		DB_PORT:     getEnv("DB_PORT", "5432"),
		DB_NAME:     getEnv("DB_NAME", "db_go_test"),
		DB_USERNAME: getEnv("DB_USERNAME", "postgres"),
		DB_PASSWORD: getEnv("DB_PASSWORD", ""),
	}, nil
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}
